package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8 "k8s.io/apimachinery/pkg/types"
)

func TestUpdateFetchDeployment(t *testing.T) {
	patchData := `
	{
		"spec": {
			"template": {
				"spec": {
					"containers": [
						{
							"name": "main-app",
							"readinessProbe": {
								"failureThreshold": 1,
								"initialDelaySeconds": 1
							},
							"startupProbe": {
								"exec": {
									"command": [
										"idm-service",
										"is-startup",
										"--auto"
									]
								},
								"failureThreshold": 100,
								"initialDelaySeconds": 1,
								"periodSeconds": 10,
								"successThreshold": 1,
								"timeoutSeconds": 10
							}
						}
					]
				}
			}
		}
	}
	`
	initK8sClient()
	res, err := k8sclient.AppsV1().Deployments("lab06").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		t.Fatalf("Error creating k8s client: %s", err)
	}
	for _, dep := range res.Items {
		if strings.HasPrefix(dep.Name, "pdu-app") || strings.HasPrefix(dep.Name, "vendor-app") {
			for _, c := range dep.Spec.Template.Spec.Containers {
				if c.Name == "main-app" {
					if c.ReadinessProbe != nil {
						_, err := k8sclient.AppsV1().Deployments("lab06").Patch(context.TODO(), dep.Name,
							k8.StrategicMergePatchType, []byte(patchData), v1.PatchOptions{})
						if err != nil {
							t.Logf("Error while patching deployment %s. Error %s\n", dep.Name, err)
						} else {
							t.Logf("Deployment %s patched successfully \n", dep.Name)
						}
					} else {
						fmt.Printf("ReadinessProbe Probe is empty for %s\n", dep.Name)
					}
				}
			}
		}
	}
}
