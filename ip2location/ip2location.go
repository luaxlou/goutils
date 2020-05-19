package ip2location

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"github.com/luaxlou/goutils/tools/fileutils"
)

var instance *ip2region.Ip2Region

var pathname = "./ip2region.db"

func init() {
	downloadIfNotExists()
	r, err := ip2region.New(pathname)

	if err != nil {
		panic(err)
	}

	instance = r
}

var dataUrl = "https://github.com/lionsoul2014/ip2region/raw/master/data/ip2region.db"
//var dataUrl = "https://gitee.com/lionsoul/ip2region/raw/master/data/ip2region.db"

func downloadIfNotExists() {


	if fileutils.Exists(pathname) {

		return
	}

	log.Println("Downloading ip2region.db from", dataUrl)

	res, err := http.Get(dataUrl)

	if err != nil {
		panic(err)
	}

	f, err := os.Create(pathname)

	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)

	f.Close()
}

func GetLocation(ip string) (ip2region.IpInfo, error) {

	return instance.MemorySearch(ip)

}

func Close() {
	instance.Close()

}
