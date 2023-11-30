/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"shellclient/pkg/cli"

	"github.com/spf13/cobra"
)

// resourceCmd represents the resource command
var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Resource model",
	Long:  `ICOS resource model.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Parent().Use == "get" {
			cli.GetResource()
		}
	},
}

func init() {
	var getResourceCmd = *resourceCmd
	getCmd.AddCommand(&getResourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
