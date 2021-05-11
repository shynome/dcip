package dcip

import "fmt"

// ssh -NT -L 0.0.0.0:5432:172.17.0.5:5432 debian@example.host
func MakeForwardPortCmd(host []string, cport string, lport string) []string {
	cmd := []string{
		"-NT",
		"-L", fmt.Sprintf("%s:%s", lport, cport),
	}
	cmd = append(cmd, host...)
	return cmd
}
