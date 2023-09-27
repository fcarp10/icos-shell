/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generates the documentation of icos-shell in markdown files",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pathDocs, _ := cmd.Flags().GetString("path")
		err := os.Mkdir(pathDocs, 0755)
		if err != nil {
			log.Fatal(err)
		}
		err = doc.GenMarkdownTree(rootCmd, pathDocs)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	createCmd.AddCommand(docsCmd)

	docsCmd.PersistentFlags().StringP("path", "", "", "path where to generate the documentation")
	docsCmd.MarkPersistentFlagRequired("path")
}
