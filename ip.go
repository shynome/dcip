package dcip

import "fmt"

var getNetworks = "docker network ls  --filter=driver=bridge --format='{{.Name}}'"
var getAllConatinersIP = fmt.Sprintf("for dn in $(%s);do docker network inspect $dn --format '{{range $k,$c:=.Containers}}{{$k}}/{{.IPv4Address}}{{println}}{{end}}';done", getNetworks)
var getContainerID = "docker ps --latest -q --no-trunc --filter='name=%s'"
var getIPOnly = "awk -F '/' '{print $2}'"
var getContainerIPCmdFormat = fmt.Sprintf("%s | grep $(%s) | %s", getAllConatinersIP, getContainerID, getIPOnly)

func MakeGetContainerIPCmd(name string) string {
	return fmt.Sprintf(getContainerIPCmdFormat, name)
}
