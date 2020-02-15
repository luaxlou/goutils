package wechatapp

import (
	"errors"
	"fmt"

	"github.com/luaxlou/gohttpclient"
)

const Host = "https://api.weixin.qq.com"

type Res struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type App struct {
	appId  string
	secret string
}

func New(appId, secret string) *App {

	c := App{
		appId:  appId,
		secret: secret,
	}

	return &c
}

type LoginRes struct {
	Res
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

func (c *App) Login(code string) (LoginRes, error) {

	url := fmt.Sprintf(Host+"/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", c.appId, c.secret, code)

	var res LoginRes

	_, err := gohttpclient.Get(url).Exec().RenderJSON(&res)

	if err != nil {

		return res, err
	} else if res.Errcode != 0 {

		return res, errors.New(res.Errmsg)

	}

	return res, nil

}
