package baidumap

import (
	"github.com/luaxlou/gohttpclient"
)

type BatchReq struct {
	Reqs []BatchReqReq `json:"reqs"`
}

type BatchReqReq struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

type BatchReverseGeocodingRes struct {
	Status      int                   `json:"status"`
	BatchResult []ReverseGeocodingRes `json:"batch_result"`
}

type ReverseGeocodingRes struct {
	Status int                         `json:"status"`
	Result BatchReverseGeocodingResult `json:"result"`
}

type BatchReverseGeocodingResult struct {
	FormattedAddress string                           `json:"formatted_address"`
	AddressComponent ReverseGeocodingAddressComponent `json:"addressComponent"`
}


type ReverseGeocodingAddressComponent struct {
	Country      string `json:"country"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
}

func (m *BaiduMap) BatchReverseGeocoding(req BatchReq) (res BatchReverseGeocodingRes) {

	url := BASE_URL + "/batch"

	gohttpclient.PostBody(url, &req).Exec().RenderJSON(&res)
	return

}
