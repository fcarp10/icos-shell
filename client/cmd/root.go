/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"shellclient/pkg/cli"
	"shellclient/pkg/openapi"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "icos-shell",
	Short: "icos-shell - a CLI to interface the ICOS Shell",
	Long: `icos-shell - a CLI to interface the ICOS Shell
   
The icos-shell can be used to modify or inspect resources in the ICOS controller from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("controller") != "" {
			openapi.Init(viper.GetString("controller"))
			cli.GetHealthcheck()
		} else {
			if viper.GetString("lighthouse") != "" {
				fmt.Fprintln(os.Stderr, "Retrieving controllers from lighthouse...")
				openapi.Init(viper.GetString("lighthouse"))
				cli.GetController()
				fmt.Fprintln(os.Stderr, "Please, add a controller to the config file. Exiting.")
			} else {
				fmt.Fprintln(os.Stderr, "Lighthouse not defined")
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing icos-shell '%s'", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$XDG_CONFIG_HOME/icos-shell/config.yaml", "config file")
	viper.BindEnv("controller", "CONTROLLER")
	viper.BindEnv("auth_token", "ICOS_AUTH_TOKEN")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".client" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".client")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Config:", viper.ConfigFileUsed())
	}

	if viper.GetString("controller") == "" {
		fmt.Fprintln(os.Stderr, "Controller not defined in the config file, trying lighthouse instead...")
		openapi.Init(viper.GetString("lighthouse"))
	} else if viper.GetString("controller") != "" && viper.GetString("auth_token") == "" {
		fmt.Fprintln(os.Stderr, "Controller is defined, but no token found")
		openapi.Init(viper.GetString("controller"))
	} else {
		fmt.Fprintln(os.Stderr, "Controller:", viper.GetString("controller"))
		token := viper.GetString("auth_token")
		token = strings.ReplaceAll(token, "\n", "")
		viper.Set("auth_token", token)
		fmt.Fprintln(os.Stderr, "Token found:", viper.GetString("auth_token"))
		openapi.Init(viper.GetString("controller"))
	}
}
