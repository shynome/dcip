package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func PrintCmd(cmd *exec.Cmd) {
	if !CLI.Debug {
		return
	}
	fmt.Printf("+ %+v \n", cmd)
}

func RunCommand(command string) *exec.Cmd {
	cmd := exec.Command("bash", "-c", command)
	return cmd
}

func RunSSHCommand(host []string, command string) *exec.Cmd {
	var cmd *exec.Cmd
	params := host
	params = append(params, command)
	cmd = exec.Command("ssh", params...)
	return cmd
}

func ParseHostOptions(host string) []string {
	if !strings.HasPrefix(host, " ") {
		return []string{host}
	}
	return strings.Split(host, " ")[1:]
}
