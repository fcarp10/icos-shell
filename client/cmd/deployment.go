/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"shellclient/pkg/cli"

	"github.com/spf13/cobra"
)

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Deployment resource",
	Long:  `ICOS deployment model.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Parent().Use == "create" {
			fileDescriptorString, _ := cmd.Flags().GetString("file")
			fileDescriptor, err := os.ReadFile(fileDescriptorString)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			cli.CreateDeployment(fileDescriptor)
		} else if cmd.Parent().Use == "get" {
			specificId, _ := cmd.Flags().GetString("id")
			if specificId != "" {
				cli.GetDeploymentById(specificId)
			} else {
				cli.GetDeployment()
			}
		} else if cmd.Parent().Use == "update" {
			specificId, _ := cmd.Flags().GetString("id")
			fileDescriptorString, _ := cmd.Flags().GetString("file")
			fileDescriptor, err := os.ReadFile(fileDescriptorString)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			cli.UpdateDeployment(specificId, fileDescriptor)
		} else if cmd.Parent().Use == "delete" {
			specificId, _ := cmd.Flags().GetString("id")
			cli.DeleteDeployment(specificId)
		}
	},
}

func init() {

	var createDeploymentCmd = *deploymentCmd
	var getDeploymentCmd = *deploymentCmd
	var updateDeploymentCmd = *deploymentCmd
	var deleteDeploymentCmd = *deploymentCmd
	createCmd.AddCommand(&createDeploymentCmd)
	deleteCmd.AddCommand(&deleteDeploymentCmd)
	updateCmd.AddCommand(&updateDeploymentCmd)
	getCmd.AddCommand(&getDeploymentCmd)

	createDeploymentCmd.PersistentFlags().StringP("file", "", "", "App descriptor file")
	updateDeploymentCmd.PersistentFlags().StringP("file", "", "", "App descriptor file")
	getDeploymentCmd.PersistentFlags().StringP("id", "", "", "ID of the deployment")
	updateDeploymentCmd.PersistentFlags().StringP("id", "", "", "ID of the deployment")
	deleteDeploymentCmd.PersistentFlags().StringP("id", "", "", "ID of the deployment")

}
