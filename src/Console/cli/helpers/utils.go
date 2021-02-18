package helpers

import (
	"fmt"
	"os"
)

//try different file names
func TryFile(files []string) string {
	_, wd := GetCurrentWorkingDirectory()
	var validFile string = ""
	for _, file := range files {
		configFile, err := os.Open(fmt.Sprintf("%s/%s", wd, file))
		defer configFile.Close()
		if err == nil {
			validFile = file
		}
	}
	return validFile
}

//get current working directory
func GetCurrentWorkingDirectory() (error, string) {
	dir, err := os.Getwd()
	if err != nil {
		return err, ""
	}
	return nil, dir
}
