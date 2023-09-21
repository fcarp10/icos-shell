package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	openapi "shellclient/openapi"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

var client *openapi.APIClient

func main() {
	cfg := openapi.NewConfiguration()
	client = openapi.NewAPIClient(cfg)

	fileExists := func(filename string) bool {
		stat, _ := os.Stat(filename)
		return stat != nil
	}
	initConfigFileInputSource := func(configFlag string, flags []cli.Flag) cli.BeforeFunc {
		return func(context *cli.Context) error {
			configFile := context.String(configFlag)
			if context.IsSet(configFlag) && !fileExists(configFile) {
				return fmt.Errorf("config file %s does not exist", configFile)
			} else if !context.IsSet(configFlag) && !fileExists(configFile) {
				return nil
			}
			inputSource, err := altsrc.NewYamlSourceFromFile(configFile)
			if err != nil {
				return err
			}
			return altsrc.ApplyInputSourceValues(context, inputSource, flags)
		}
	}
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			EnvVars:     []string{"CONFIG_FILE"},
			Value:       "config.yaml",
			DefaultText: "config.yaml",
			Usage:       "config file",
		},
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "server",
			Value: "localhost:8080",
			Usage: "URL of the shell-backend",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "username",
			Value: "admin",
			Usage: "username",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "password",
			Usage: "password",
		}),
	}

	commands := []*cli.Command{
		{
			Name:    "controller",
			Aliases: []string{"c"},
			Usage:   "options for controllers",
			Subcommands: []*cli.Command{
				{
					Name:    "add",
					Aliases: []string{"a"},
					Usage:   "add a controller",
					Action: func(cCtx *cli.Context) error {
						cfg.Host = cCtx.String("server")
						controller := *openapi.NewController("controller_"+strconv.Itoa(rand.Intn(1000)), cCtx.Args().First())
						r, err := client.ControllerApi.AddController(context.Background()).Username(cCtx.String("username")).Password(cCtx.String("password")).Controller(controller).Execute()
						if err != nil {
							fmt.Fprintf(os.Stderr, "Error when adding a controller: %v\n", err)
							fmt.Fprintf(os.Stderr, "%v\n", r.Body)
						} else {
							fmt.Println("Controller added successfully!")
						}
						return nil
					},
					Before: initConfigFileInputSource("config", flags),
				},
				{
					Name:  "list",
					Usage: "list controllers",
					Action: func(cCtx *cli.Context) error {
						cfg.Host = cCtx.String("server")
						resp, _, err := client.ControllerApi.GetControllers(context.Background()).Execute()
						if err != nil {
							fmt.Fprintf(os.Stderr, "Error when retrieving controllers: %v\n", err)
						} else {
							for _, element := range resp {
								fmt.Println(element)
							}
						}

						return nil
					},
					Before: initConfigFileInputSource("config", flags),
				},
			},
		},
	}

	app := &cli.App{
		Version:              "v0.1",
		EnableBashCompletion: true,
		Name:                 "icos-cli",
		Usage:                "Shell",
		Before:               initConfigFileInputSource("config", flags),
		Flags:                flags,
		Commands:             commands,
		Action: func(cCtx *cli.Context) error {
			server := cCtx.String("server")
			if server != "" {
				fmt.Println("Trying to connect to", server)
				cfg.Host = server
			} else {
				fmt.Println("Server not specified, trying with ", server)
			}
			_, err := client.DefaultApi.GetHealthcheck(context.Background()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error connecting to server: %v\n", err)
			} else {
				fmt.Println("Server connection successful!")
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
