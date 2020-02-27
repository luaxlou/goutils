package payment

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/clbanning/mxj"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/liyoung1992/wechatpay"
	"github.com/luaxlou/goutils/tools/logutils"
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

func (w *Payment) HandleNotify(c *gin.Context, onSucc, onFail func(rawBody string)) {

	log.Info("wechat notify start")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error(err, "read http body failed！error msg:"+err.Error())
	}

	rawBody := string(body)

	log.Info("wechat pay notify body :" + rawBody)

	var wx_notify_req wechatpay.PayNotifyResult

	err = xml.Unmarshal(body, &wx_notify_req)
	if err != nil {
		log.Error(err, "read http body xml failed! err :"+err.Error())
	}
	mv, err := mxj.NewMapXml(body)
	if err != nil {
		log.Error(err, err.Error())
	}

	if w.client.VerifySign(mv["xml"].(map[string]interface{}), mv["xml"].(map[string]interface{})["sign"].(string)) {
		record, err := json.Marshal(wx_notify_req)

		logutils.PrintObj(record)
		if err != nil {
			log.Error(err, "wechat pay marshal err :"+err.Error())
		}
		c.XML(http.StatusOK, gin.H{
			"return_code": "SUCCESS",
			"return_msg":  "OK",
		})
	} else {
		c.XML(http.StatusOK, gin.H{
			"return_code": "FAIL",
			"return_msg":  "failed to verify sign, please retry!",
		})
	}
	return
}

const PaySignTemplate = "appId=%s&nonceStr=%s&package=prepay_id=%s&signType=MD5&timeStamp=%d&key=%s"

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
