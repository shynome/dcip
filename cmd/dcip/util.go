package main

import (
	"fmt"
	"os/exec"
)

func PrintCmd(cmd *exec.Cmd) {
	if !CLI.Debug {
		return
	}
	fmt.Print("+ ")
	fmt.Println(cmd)
}

func RunCommand(command string) *exec.Cmd {
	cmd := exec.Command("bash", "-c", command)
	return cmd
}

func RunSSHCommand(host interface{}, command string) *exec.Cmd {
	var cmd *exec.Cmd
	switch host.(type) {
	case string:
		cmd = exec.Command("ssh", host.(string), command)
	default:
		params := host.([]string)
		params = append(params, command)
		cmd = exec.Command("ssh", params...)
	}
	return cmd
}
