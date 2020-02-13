package message

import (
	"errors"
	"fmt"

	"github.com/luaxlou/gohttpclient"
	"github.com/luaxlou/goutils/tools/logutils"
	"github.com/luaxlou/goutils/wechat/corpwechat"
	"github.com/luaxlou/goutils/wechat/corpwechat/client"
)

type Message struct {
	client  *client.Client `json:"message"`
	agentId int64
}

func New(c *client.Client, agentId int64) *Message {

	msg := Message{
		client:  c,
		agentId: agentId,
	}

	return &msg
}

type SendRes struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

type TextMessage struct {
	Content string `json:"content"`
}

type SendTextMessageReq struct {
	ToUser               string      `json:"touser"`
	ToParty              string      `json:"toparty"`
	ToTag                string      `json:"totag"`
	MsgType              string      `json:"msgtype"`
	AgentId              int64       `json:"agentid"`
	Text                 TextMessage `json:"text"`
	Safe                 int         `json:"safe"`
	EnableIdTrans        int         `json:"enable_id_trans"`
	EnableDuplicateCheck int         `json:"enable_duplicate_check"`
}

func (c *Message) SendTextMessage(req SendTextMessageReq) error {

	req.MsgType = "text"
	req.AgentId = c.agentId

	token := c.client.GetAccessToken()
	url := fmt.Sprintf(corpwechat.Host+"/cgi-bin/message/send?access_token=%s", token)

	var res SendRes

	_, err := gohttpclient.PostBody(url, &req).Exec().RenderJSON(&res)

	logutils.PrintObj(req)
	logutils.PrintObj(res)

	if err != nil {

		return err
	} else if res.Errcode != 0 {

		return errors.New(res.Errmsg)

	}

	return nil
}
