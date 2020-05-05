package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
	)

	//cmd = exec.Command("bin/bash","-c","echo 1; echo2;")
	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "whereis ls")

	err = cmd.Run()

	fmt.Println(err)
}
