package build

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os/exec"
	"strings"
)


//select selected workspace
func SelectTerraformWorkspace(workspace string) bool {
	var cmd *exec.Cmd
	color.Set(color.FgMagenta, color.Bold)
	fmt.Print("\nSelecting Terraform Workspace....... \n")
	color.Unset()

	cmd = exec.Command("terraform", "workspace", "select", workspace)
	_, cmdErr := StdnWriter(cmd)
	if len(strings.TrimSpace(cmdErr)) != 0{
		color.Set(color.FgRed, color.Bold)
		log.Fatal("\nError: Moti Test has been aborted, Could not switch to workspace \n")
		color.Unset()
	}
	return true
}

//initialize terraform worspace
func InitializeTerraformWorkspace() bool {
	var cmd *exec.Cmd
	color.Set(color.FgMagenta, color.Bold)
	fmt.Print("\nInitializing Terraform Files....... \n")
	color.Unset()

	//run build command
	cmd = exec.Command("terraform", "init")
	_, cmdErr := StdnWriter(cmd)
	if len(strings.TrimSpace(cmdErr)) != 0{
		color.Set(color.FgRed, color.Bold)
		log.Fatal("\nError: Moti Test has been aborted, resolve issues first \n")
		color.Unset()
	}
	return true
}

//validate terraform files
func ValidateTerraformFiles() bool {
	var cmd *exec.Cmd
	color.Set(color.FgMagenta, color.Bold)
	fmt.Print("\nValidating Terraform Files....... \n")
	color.Unset()

	//run build command
	cmd = exec.Command("terraform", "validate", ".", "-recursive")
	_, cmdErr := StdnWriter(cmd)
	if len(strings.TrimSpace(cmdErr)) != 0{
		color.Set(color.FgRed, color.Bold)
		log.Fatal("\nError: Moti Test has been aborted, resolve syntax issues first \n")
		color.Unset()
	}
	return true
}

//provision resources
func ProvisionResources() bool {
	var cmd *exec.Cmd
	color.Set(color.FgMagenta, color.Bold)
	fmt.Print("\nProvisioning Resources for test....... \n")
	color.Unset()

	//run build command
	cmd = exec.Command("terraform", "apply", "-auto-approve")
	_, cmdErr := StdnWriter(cmd)
	if len(strings.TrimSpace(cmdErr)) != 0{
		color.Set(color.FgRed, color.Bold)
		fmt.Print("\nError: Moti Test has been affected, couldn't provision resources. will deprovision created resources \n")
		color.Unset()
		return false
	}
	return true
}

//deprovision resources
func DeProvisionResources() bool {
	var cmd *exec.Cmd
	color.Set(color.FgMagenta, color.Bold)
	fmt.Print("\nProvisioning Resources for test....... \n")
	color.Unset()

	//run build command
	cmd = exec.Command("terraform", "destroy", "-auto-approve")
	_, cmdErr := StdnWriter(cmd)
	if len(strings.TrimSpace(cmdErr)) != 0{
		color.Set(color.FgRed, color.Bold)
		log.Fatal("\nError: Moti Test has been aborted, couldn't provision resources \n")
		color.Unset()
	}
	return true
}