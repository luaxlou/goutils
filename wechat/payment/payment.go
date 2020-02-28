package payment

import (
	"encoding/xml"
	"log"

	"github.com/clbanning/mxj"
	"github.com/liyoung1992/wechatpay"
)

const (
	TRADE_TYPE_WEAPP = "JSAPI"
)

type Payment struct {
	client *wechatpay.WechatPay
}

func New(appId string, merchantId string, apiKey string, weichatKey, wechatCert []byte) *Payment {

	c := wechatpay.New(appId, merchantId, apiKey, weichatKey, wechatCert)

	return &Payment{
		client: c,
	}

}

func (c *Payment) Pay(tradeNo string, totalFee int, subject string, notifyUrl string, tradeType string, openId string, clientIp string) (*wechatpay.UnifyOrderResult, error) {
	var pay_data wechatpay.UnitOrder
	pay_data.NotifyUrl = notifyUrl
	pay_data.TradeType = tradeType
	pay_data.Body = subject
	pay_data.SpbillCreateIp = clientIp
	pay_data.TotalFee = totalFee
	pay_data.OutTradeNo = tradeNo
	pay_data.Openid = openId

	return c.client.Pay(pay_data)
}

func (w *Payment) VerifyNotify(body []byte) (wechatpay.PayNotifyResult, bool) {

	var result wechatpay.PayNotifyResult

	err := xml.Unmarshal(body, &result)
	if err != nil {
		log.Println("read http body xml failed! err :" + err.Error())
	}
	mv, err := mxj.NewMapXml(body)
	if err != nil {
		log.Println(err.Error())
	}

	if w.client.VerifySign(mv["xml"].(map[string]interface{}), mv["xml"].(map[string]interface{})["sign"].(string)) {

		return result, true
	}

	return result, false

}

// 小程序 客户端唤起支付签名
func (c *Payment) GetPaySign(nonceStr string, prepayId string, ts string) string {

	m := map[string]interface{}{
		"appId":     c.client.AppId,
		"nonceStr":  nonceStr,
		"package":   "prepay_id=" + prepayId,
		"signType":  "MD5",
		"timeStamp": ts,
	}

	sign := wechatpay.GetSign(m, c.client.ApiKey)

	return sign
}
