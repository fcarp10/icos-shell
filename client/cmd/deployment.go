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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Parent().Use == "create" {
			fileDescriptorString, _ := cmd.Flags().GetString("file")
			fileDescriptor, err := os.ReadFile(fileDescriptorString)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			cli.CreateDeployment(fileDescriptor)
		} else if cmd.Parent().Use == "get" {
			specificId, _ := cmd.Flags().GetInt64("id")
			if specificId != 0 {
				cli.GetDeploymentById(specificId)
			} else {
				cli.GetDeployment()
			}
		}
	},
}

func init() {

	var createDeploymentCmd = *deploymentCmd
	var getDeploymentCmd = *deploymentCmd
	createCmd.AddCommand(&createDeploymentCmd)
	getCmd.AddCommand(&getDeploymentCmd)

	createDeploymentCmd.PersistentFlags().StringP("file", "", "", "App descriptor file")
	getDeploymentCmd.PersistentFlags().Int64P("id", "", 0, "ID of the deployment")

}
