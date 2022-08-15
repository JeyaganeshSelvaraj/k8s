package main

import (
	"log"
	"time"

	"github.com/mitchellh/go-homedir"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var k8sclient *kubernetes.Clientset

func createNewK8sClient(kubeconfig string) (*kubernetes.Clientset, error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	cfg.Timeout = 10 * time.Second
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
func initK8sClient() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Unable to get home directory. %s", err)
	}
	k8sclient, err = createNewK8sClient(home + "/.kube/config")
	if err != nil {
		log.Fatalf("Unable to create a k8s client. %s", err)
	}
}
