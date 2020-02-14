package client

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/luaxlou/gohttpclient"
)

var Host = "https://qyapi.weixin.qq.com"

type Client struct {
	corpId        string
	secret        string
	accessToken   string
	lastTokenTime time.Time
}

func New(corpId, secret string) *Client {

	c := Client{
		corpId:        corpId,
		secret:        secret,
		lastTokenTime: time.Now(),
	}

	return &c
}

type GetAccessTokenRes struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) GetAccessToken() string {

	t := time.Now()

	if c.accessToken == "" || t.Sub(c.lastTokenTime).Seconds() > 7200 {

		url := fmt.Sprintf(Host+"/cgi-bin/gettoken?corpid=%s&corpsecret=%s", c.corpId, c.secret)

		var res GetAccessTokenRes

		_, err := gohttpclient.Get(url).Exec().RenderJSON(&res)

		if err != nil {
			log.Println(err.Error())
		} else {

			c.accessToken = res.AccessToken

			c.lastTokenTime = t

		}

	}

	return c.accessToken

}

type GetUserInfoRes struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	UserId  string `json:"UserId"`
}

func (c *Client) GetUserInfo(code string) (string, error) {

	token := c.GetAccessToken()
	url := fmt.Sprintf(Host+"/cgi-bin/user/getuserinfo?access_token=%s&code=%s", token, code)

	var res GetUserInfoRes

	_, err := gohttpclient.Get(url).Exec().RenderJSON(&res)

	if err != nil {

		return "", err
	} else if res.Errcode != 0 {

		return "", errors.New(res.Errmsg)

	}

	return res.UserId, nil

}
