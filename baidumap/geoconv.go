package baidumap

import (
	"fmt"
	"github.com/luaxlou/gohttpclient"
	"github.com/shopspring/decimal"
)

type GeoConvRes struct {
	Status int           `json:"status"`
	Result GeoConvResult `json:"result"`
}

type GeoConvResult struct {
	X decimal.Decimal `json:"x"`
	Y decimal.Decimal `json:"y"`
}

func (m *BaiduMap) GeoConv(lng, lat decimal.Decimal, from int, to int) (res GeoConvRes) {

	url := fmt.Sprintf(BASE_URL+"/geoconv/v1/?coords=%s,%s&from=%d&to=%d&ak=%s", lng, lat, from, to,m.AK)

	gohttpclient.Get(url).Exec().RenderJSON(&res)
	return

}
