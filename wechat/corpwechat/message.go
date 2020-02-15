package corpwechat

import (
	"errors"
	"fmt"

	"github.com/luaxlou/gohttpclient"
)

type SendRes struct {
	Res
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
	Safe                 int         `json:"safe"`
	EnableIdTrans        int         `json:"enable_id_trans"`
	EnableDuplicateCheck int         `json:"enable_duplicate_check"`
	Text                 TextMessage `json:"text"`
}

func (c *App) SendTextMessage(req SendTextMessageReq) error {

	req.MsgType = "text"
	req.AgentId = c.agentId

	token := c.GetAccessToken()
	url := fmt.Sprintf(Host+"/cgi-bin/message/send?access_token=%s", token)

	var res SendRes

	_, err := gohttpclient.PostBody(url, &req).Exec().RenderJSON(&res)

	if err != nil {

		return err
	} else if res.Errcode != 0 {

		return errors.New(res.Errmsg)

	}

	return nil
}
