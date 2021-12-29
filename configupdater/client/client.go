package client

import (
	"configupdater/types"
	"context"
	"log"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	k8sclient *kubernetes.Clientset
}

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func CreateNewK8sClient(kubeconfig string) (*Client, error) {
	cfg, err := buildConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	cfg.Timeout = 10 * time.Second
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	return &Client{
		k8sclient: clientset,
	}, nil
}
func updateK8sConfigData(k8scmData map[string]string, cmUpdateData map[string]string) error {
	for k, v := range cmUpdateData {
		if _, ok := k8scmData[k]; !ok {
			log.Printf("Adding key: %s, value: %s", k, v)
		} else {
			log.Printf("Key: %s already exists, updating", k)
		}
		k8scmData[k] = v
	}
	return nil
}
func deleteK8sConfigData(k8scmData map[string]string, dataTodel []string) error {
	for _, k := range dataTodel {
		if _, ok := k8scmData[k]; !ok {
			log.Printf("Key: %s does not exist, skipping", k)
		} else {
			log.Printf("Deleting key: %s", k)
			delete(k8scmData, k)
		}
	}
	return nil
}
func (client *Client) UpdateConfigMap(cm types.ConfigMap) error {
	log.Println("Updating configmap: ", cm.Name)
	k8scm, err := client.k8sclient.CoreV1().ConfigMaps(cm.Namespace).Get(context.TODO(), cm.Name, v1.GetOptions{})
	if err != nil {
		return err
	}
	if k8scm.Data == nil {
		k8scm.Data = make(map[string]string)
	}
	updateK8sConfigData(k8scm.Data, cm.DataToAdd)
	updateK8sConfigData(k8scm.Data, cm.DataToUpd)
	deleteK8sConfigData(k8scm.Data, cm.DataToDel)
	_, err = client.k8sclient.CoreV1().ConfigMaps(cm.Namespace).Update(context.TODO(), k8scm, v1.UpdateOptions{})
	return err
}
