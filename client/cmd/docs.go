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
	Short: "Docs in markdown format",
	Long:  `Generates the documentation of the icos-shell in markdown format.`,
	Run: func(cmd *cobra.Command, args []string) {
		pathDocs, _ := cmd.Flags().GetString("path")
		err := os.MkdirAll(pathDocs, 0755)
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
	updateCmd.AddCommand(docsCmd)
	deleteCmd.AddCommand(docsCmd)

	docsCmd.PersistentFlags().StringP("path", "", "", "path where to generate the documentation")
	docsCmd.MarkPersistentFlagRequired("path")
}
