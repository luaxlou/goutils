package dockerutils

import (
	"github.com/lixiangzhong/dnsutil"
	"github.com/luaxlou/goutils/tools/logutils"
)

func GetMachineIp() string {

	var dig dnsutil.Dig
	dig.SetDNS("8.8.8.8")
	a, err := dig.A("host.docker.internal")
	//a, err := dig.A("www.baidu.com")

	if err != nil  || a == nil{
		return ""
	}

	logutils.PrintObj(a)

	return a[0].A.String()

}
