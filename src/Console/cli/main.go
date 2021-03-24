package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/moti/console/build"
	"github.com/moti/console/data"
	"log"
	"os"
	"strings"
	"time"

	"github.com/moti/console/helpers"
	"github.com/moti/console/validation"
	"github.com/urfave/cli"
)

// main
func main() {
	app := cli.NewApp()
	app.Name = "Moti"
	app.Version = "0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Moti | Tangent Solutions",
			Email: "info@tangentsolutions.co.za",
		},
	}
	app.Copyright = fmt.Sprintf("(c) %d Moti", time.Now().Year())
	app.UsageText = "moti [global options] command [command options] [arguments...]"
	app.Usage = "IAC tests, push based provisioning and scheduled auto provisioning, scaling and deprovisioning"
	app.Commands = []cli.Command{
		{
			Name:  "test",
			Usage: "validates and tests terraform code and provisioned resources",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "folder, f"},
				cli.StringFlag{Name: "output, o"},
				cli.StringFlag{Name: "workspace, w"},
			},
			Action: func(c *cli.Context) error {
				//set default config
				_, wd := helpers.GetCurrentWorkingDirectory()
				var MotiMap helpers.Artifact
				SourceDirectory := wd
				Output := "CONSOLE"
				Workspace := "default"


				//set config file to parsed file
				if len(strings.TrimSpace(c.String("folder"))) !=0 {
					SourceDirectory = strings.TrimSpace(c.String("folder"))
				}
				if len(strings.TrimSpace(c.String("output"))) != 0{
					Output = strings.TrimSpace(c.String("output"))
				}
				if len(strings.TrimSpace(c.String("workspace"))) != 0{
					Workspace = strings.TrimSpace(c.String("workspace"))
				}

				MotiMap = helpers.Artifact{
					Folder:         SourceDirectory,
					Output:         Output,
					Worspace: Workspace,
				}

				//validate installation and source
				validation.Validate(MotiMap.Folder)

				//initialize terraform workspace
				build.InitializeTerraformWorkspace()
				color.Set(color.FgGreen, color.Bold)
				fmt.Print("\nSuccess: Terraform Workspace initialized, Proceeding to the next Step...... \n")
				color.Unset()

				//initialize terraform workspace
				build.SelectTerraformWorkspace(MotiMap.Worspace)
				color.Set(color.FgGreen, color.Bold)
				fmt.Print("\nSucess: Terraform Workspace selected, Proceeding to the next Step...... \n")
				color.Unset()

				//validate terraform files
				build.ValidateTerraformFiles()
				color.Set(color.FgGreen, color.Bold)
				fmt.Print("\nSuccess: Terraform files Valid, Proceeding to the next Step...... \n")
				color.Unset()

				//provision resources
				build.ProvisionResources()
				color.Set(color.FgGreen, color.Bold)
				fmt.Print("\nSuccess: Resources Provisioned, Proceeding to the next Step...... \n")
				color.Unset()

				/////////////////////////////////////////////////////////////////////////////////
				//Tests begin
				canContinue, _ := data.OrchestrateGatherTestData()
				if !canContinue {
					color.Set(color.FgRed, color.Bold)
					fmt.Print("\nError: Error(s) Occured Reading Terraform files...... \n")
					color.Unset()
					build.DeProvisionResources()
				}


				//provision resources
				build.DeProvisionResources()
				color.Set(color.FgGreen, color.Bold)
				fmt.Print("Resources Deprovisioned, Please find the results of the tests in your Wordk Directory...... \n")
				color.Unset()

				return nil
			},
		},
	}

	//run cli
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
