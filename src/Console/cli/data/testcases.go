package data

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/moti/console/build"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func AllocateTestCaseInstructions() bool{
	var cmd *exec.Cmd
	color.Set(color.FgMagenta, color.Bold)
	fmt.Print("\nProvisioning Resources for test....... \n")
	color.Unset()

	//run build command
	cmd = exec.Command("terraform", "apply", "-auto-approve")
	_, cmdErr := build.StdnWriter(cmd)
	if len(strings.TrimSpace(cmdErr)) != 0{
		color.Set(color.FgRed, color.Bold)
		fmt.Print("\nError: Moti Test has been affected, couldn't provision resources. will deprovision created resources \n")
		color.Unset()
		return false
	}
	return true
}

func GetAllTargetFiles() []string{
	files, err := ioutil.ReadDir(".")
	sourceFiles := make([]string, 0)
	if err != nil {
		color.Set(color.FgRed, color.Bold)
		log.Fatal("\r\nERROR: Cant read files, Could be a permissions issue\r\n")
		color.Unset()
	}
	for _, f := range files {
		if strings.Contains(f.Name(), ".tf"){
			sourceFiles = append(sourceFiles, f.Name())
		}
	}
	fmt.Print(sourceFiles)
	return sourceFiles
}
