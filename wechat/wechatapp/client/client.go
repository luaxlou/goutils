package client

import (
	"errors"
	"fmt"
	"time"

	"github.com/luaxlou/gohttpclient"
)

const Host = "https://api.weixin.qq.com"

type Res struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type Client struct {
	appId         string
	secret        string
	accessToken   string
	lastTokenTime time.Time
}

func New(appId, secret string) *Client {

	c := Client{
		appId:         appId,
		secret:        secret,
		lastTokenTime: time.Now(),
	}

	return &c
}

type LoginRes struct {
	Res
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

func (c *Client) Login(code string) (LoginRes, error) {

	url := fmt.Sprintf(Host+"/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", c.appId, c.secret, code)

	var res GetUserInfoRes

	_, err := gohttpclient.Get(url).Exec().RenderJSON(&res)

	if err != nil {

		return "", err
	} else if res.Errcode != 0 {

		return "", errors.New(res.Errmsg)

	}

	return res, nil

}
