package main

import (
	"fmt"
	"os/exec"
)

func ExeCommand(commandStr string) string {
	commands := "echo hello"
	cmd := exec.CommandContext(nil,"/bin/bash", "-c", commands)
	if output, err := cmd.CombinedOutput(); err != nil {
		return err.Error()
	} else {
		return string(output)
	}
}

func main() {
	output := ExeCommand("echo hello")
	fmt.Println(output)
}
