package cmd

import (
	"configupdater/client"
	"configupdater/types"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestUpdateConfigMap(t *testing.T) {
	cm := types.ConfigMap{
		Name:      "test-cm",
		Namespace: "default",
		DataToAdd: map[string]string{"key1": "value1", "key2": "value2"},
		DataToDel: []string{"key9"},
		DataToUpd: map[string]string{"key3": "value3", "key4": "value4"},
	}
	cm.DataToAdd["key3"] = "value3"
	cm.DataToDel = append(cm.DataToDel, "key2")
	cm.DataToUpd["key2"] = "value2"
	cm.DataToUpd["key3"] = "value3"
	home, err := homedir.Dir()
	if err != nil {
		t.Fatalf("Error while getting home directory: %s", err)
	}
	k8sclient, err := client.CreateNewK8sClient(home + "/.kube/config")
	if err != nil {
		t.Fatalf("Error creating k8s client: %s", err)
	}
	err = k8sclient.UpdateConfigMap(cm)
	if err != nil {
		t.Errorf("Error updating configmap: %v", err)
	}
}
