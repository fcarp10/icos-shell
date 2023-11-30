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
	Long:  `ICOS controller model.`,
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
