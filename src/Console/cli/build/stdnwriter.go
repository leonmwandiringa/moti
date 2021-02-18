package build

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

func StdnWriter(cmd *exec.Cmd) (cmdOut string, cmdErr string){
	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		color.Set(color.FgRed, color.Bold)
		fmt.Printf("cmd.Start() failed with '%s'\n", err)
		color.Unset()
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		color.Set(color.FgRed, color.Bold)
		fmt.Sprintf("cmd.Run() failed with %s\n", err)
		color.Unset()
	}
	if errStdout != nil || errStderr != nil {
		color.Set(color.FgRed, color.Bold)
		log.Fatal("failed to capture stdout or stderr\n")
		color.Unset()
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Sprintf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	return outStr, errStr
}

