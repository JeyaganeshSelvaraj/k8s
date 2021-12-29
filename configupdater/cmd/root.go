/*
Copyright © 2021 Jeyaganesh Selvaraj <jeyagan@gmail.com>

*/
package cmd

import (
	"configupdater/client"
	"configupdater/types"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var entriesToAdd string
var entriesToUpdate string
var paramsToDelete string
var kubeconfig string
var namespace string
var configmapName string
var k8sclient *client.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "configupdater",
	Short: "A simple application to update k8s configmap of literals",
	Long:  `This is a small application to update k8s configmap of literals.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("accepts only configmap name, received %d arguments", len(args))
		}
		return nil
	},
	Run: run,
}

func updateConfigMap() error {
	cm := types.ConfigMap{
		Name:      configmapName,
		Namespace: namespace,
		DataToAdd: parseEntries(entriesToAdd),
		DataToUpd: parseEntries(entriesToUpdate),
	}
	if paramsToDelete != "" {
		cm.DataToDel = strings.Split(paramsToDelete, ",")
	}
	return k8sclient.UpdateConfigMap(cm)
}
func creatClient() {
	var err error
	k8sclient, err = client.CreateNewK8sClient(kubeconfig)
	if err != nil {
		fmt.Println("Error creating k8s client: ", err)
		os.Exit(1)
	}
}
func run(cmd *cobra.Command, args []string) {
	creatClient()
	configmapName = args[0]
	if err := updateConfigMap(); err != nil {
		fmt.Println("Error updating configmap: ", err)
		os.Exit(1)
	}
}
func parseEntries(entriesStr string) map[string]string {
	entries := make(map[string]string)
	if entriesStr != "" {
		for _, entry := range strings.Split(entriesStr, ",") {
			keyValue := strings.Split(entry, "=")
			if len(keyValue) != 2 {
				fmt.Printf("invalid entry: %s\n", entry)
				continue
			}
			entries[keyValue[0]] = keyValue[1]
		}
	}
	return entries
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.configupdater.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().StringVarP(&entriesToAdd, "add", "a", "", "Add a new configuration parameter with value with format key1=value1,key2=value2")
	rootCmd.Flags().StringVarP(&paramsToDelete, "delete", "d", "", "Delete an existing configuration parameter")
	rootCmd.Flags().StringVarP(&entriesToUpdate, "update", "u", "", "Update an existing configuration parameter with value with format key1=value1,key2=value2")
	rootCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Namespace of the configmap")
}
