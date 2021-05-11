package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/shynome/dcip"
)

var CLI struct {
	Of struct {
		Host      string `arg:"" name:"[host|name]" help:"remote host optional or just a container name"`
		Container string `arg:"" name:"name" optional:"" help:"container name"`
	} `cmd:"" passthrough:"" help:"get container ip."`
	Export struct {
		Host          string `arg:"" name:"host" help:"ssh host. example: debian@example.host"`
		ContainerPort string `arg:"" name:"cport" help:"remote container name and port. example: pg:5432"`
		LocalAddr     string `arg:"" name:"lport" optional:"" help:"local bind address and port. default bind address is 0.0.0.0, default port is remote container port. example: 127.0.0.1:5432 or 5432"`
	} `cmd:"" passthrough:"" help:"export remote container port to local host."`
	Debug   bool             `name:"debug" short:"D" optional:""`
	Version kong.VersionFlag `short:"V"`
}

func reportError(err error) {
	os.Stderr.WriteString(err.Error())
	os.Stderr.WriteString("\r\n")
	os.Exit(1)
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Vars{"version": "0.1.0"},
	)
	switch ctx.Command() {
	case "of <[host|name]>":
		fallthrough
	case "of <[host|name]> <name>":
		params := CLI.Of
		if params.Container == "" {
			params.Container = params.Host
			params.Host = ""
		}
		cmdStr := dcip.MakeGetContainerIPCmd(params.Container)
		var cmd *exec.Cmd
		if params.Host == "" {
			cmd = RunCommand(cmdStr)
		} else {
			cmd = RunSSHCommand(ParseHostOptions(params.Host), cmdStr)
		}
		PrintCmd(cmd)
		result, err := cmd.CombinedOutput()
		if err != nil {
			reportError(err)
			return
		}
		fmt.Print(string(result))
	case "export <host> <cport>":
		fallthrough
	case "export <host> <cport> <lport>":
		params := CLI.Export
		var container string
		var cport string
		var lbind string = "0.0.0.0"
		var lport string
		ContainerPortArr := strings.Split(params.ContainerPort, ":")
		container = ContainerPortArr[0]
		cport = ContainerPortArr[1]
		if container == "" || cport == "" {
			reportError(fmt.Errorf("container name and port is required"))
			return
		}
		localAddrArr := strings.Split(params.LocalAddr, ":")
		if params.LocalAddr == "" {
			lport = cport
		} else if len(localAddrArr) == 1 {
			lport = localAddrArr[0]
		} else {
			lbind = localAddrArr[0]
			lport = localAddrArr[1]
		}
		getIPCmd := RunSSHCommand(ParseHostOptions(params.Host), dcip.MakeGetContainerIPCmd(container))
		PrintCmd(getIPCmd)
		cipBytes, err := getIPCmd.CombinedOutput()
		cip := strings.Replace(string(cipBytes), "\n", "", 1)
		if err != nil {
			reportError(err)
			return
		}
		cmdStr := dcip.MakeForwardPortCmd(ParseHostOptions(params.Host), string(cip)+":"+cport, lbind+":"+lport)
		cmd := RunSSHCommand(cmdStr, "")
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		PrintCmd(cmd)
		if err := cmd.Run(); err != nil {
			reportError(err)
			return
		}
	default:
		panic(ctx.Command())
	}
}
