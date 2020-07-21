package iputils

import (
	"strings"

	"github.com/luaxlou/gohttpclient"
)

var wanIp string

func GetWanIp() string {

	if wanIp != "" {
		return wanIp
	}

	_, str, _ := gohttpclient.Get("http://icanhazip.com").Exec().String()


	return strings.Trim(str,"\n")
}
