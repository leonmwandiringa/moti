package validation

import (
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

//validate source is in check
func Validate(source string){
	ValidateTerraformInstallation()
	ValidateSourceHasTerraformFiles(source)
}
//deprecated, nolonger in use now but might be useful
func PruneWorskpace(source string){
	err := os.RemoveAll("./.terraform")
	if err != nil {
		color.Set(color.FgRed, color.Bold)
		log.Fatal("\r\nERROR: An Error Ocurred Pruning workspace, Could be permissions\r\n")
		color.Unset()
	}
}

func ValidateTerraformInstallation() bool {
	_, terraformErr := exec.LookPath("terraform")
	if terraformErr != nil {
		color.Set(color.FgRed, color.Bold)
		log.Fatal("\r\nERROR: Terraform is not installed on this machine, please install Terraform\r\n")
		color.Unset()
	}
	return terraformErr == nil
}

func ValidateSourceHasTerraformFiles(source string) bool{
	files, err := ioutil.ReadDir(source)
	if err != nil {
		color.Set(color.FgRed, color.Bold)
		log.Fatal("\r\nERROR: No such folder/source was found, please provide valid source target\r\n")
		color.Unset()
	}
	for _, f := range files {
		if strings.Contains(f.Name(), ".tf"){
			//navigate to working directory
			NavigateToSource(source)
			return true
		}
	}
	color.Set(color.FgRed, color.Bold)
	log.Fatal("\r\nERROR: No Terraform files were found, please provide a source with terraform files for testing\r\n")
	color.Unset()
	return false
}

//navigate into the source diectory to utilize already initialized source plugins
func NavigateToSource(source string){
	os.Chdir(source)
	_, err := os.Getwd()
	if err != nil {
		color.Set(color.FgRed, color.Bold)
		log.Fatal("\r\nERROR: Could navigate into source, could be a permission issue\r\n")
		color.Unset()
	}
}