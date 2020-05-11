package baidumap

import "os"

const BASE_URL = "http://api.map.baidu.com"

type BaiduMap struct {
	AK string
}

var instance *BaiduMap

func init() {

	ak := os.Getenv("BAIDUMAP_AK")

	if ak != "" {
		instance = New(ak)
	}

}

func Instance() *BaiduMap {

	return instance

}

func New(ak string) *BaiduMap {

	return &BaiduMap{
		AK:ak,
	}
}
