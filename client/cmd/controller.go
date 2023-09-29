/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
			res := cli.AddController(controllerName, controllerAddress)
			fmt.Println(res)
		} else if cmd.Parent().Use == "get" {
			res := cli.GetController()
			fmt.Println(res)
		}
	},
}

func init() {

	var getControllerCmd = *controllerCmd
	var addControllerCmd = *controllerCmd
	getCmd.AddCommand(&getControllerCmd)
	addCmd.AddCommand(&addControllerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addControllerCmd.PersistentFlags().StringP("name", "n", "", "Name of the controller")
	addControllerCmd.PersistentFlags().StringP("address", "a", "", "Address of the controller")
	addControllerCmd.MarkPersistentFlagRequired("name")
	addControllerCmd.MarkPersistentFlagRequired("address")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// controllerCmd.Flags().StringP("name", "n", "", "Name of the controller")
	// controllerCmd.Flags().StringP("address", "a", "", "Address of the controller")
}
