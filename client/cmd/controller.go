/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"shellclient/pkg/cli"

	"github.com/spf13/cobra"
)

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "Controller resource",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Parent().Use == "add" {
			controllerName, _ := cmd.Flags().GetString("name")
			controllerAddress, _ := cmd.Flags().GetString("address")
			cli.AddController(controllerName, controllerAddress)
		} else if cmd.Parent().Use == "get" {
			cli.GetController()
		}
	},
}

func init() {

	var getControllerCmd = *controllerCmd
	var addControllerCmd = *controllerCmd
	getCmd.AddCommand(&getControllerCmd)
	addCmd.AddCommand(&addControllerCmd)

	addControllerCmd.PersistentFlags().StringP("name", "n", "", "Name of the controller")
	addControllerCmd.PersistentFlags().StringP("address", "a", "", "Address of the controller")
	addControllerCmd.MarkPersistentFlagRequired("name")
	addControllerCmd.MarkPersistentFlagRequired("address")
}
