package lark

import (
	"github.com/parnurzeal/gorequest"
	"time"
)

type MsgContent struct {
	Tag      string `json:"tag"`
	UnEscape bool   `json:"un_escape"`
	Text     string `json:"text"`
}
type MsgData struct {
	Email   string `json:"email"`
	MsgType string `json:"msg_type"`
	Content struct {
		Post struct {
			ZhCn struct {
				Title   string         `json:"title"`
				Content [][]MsgContent `json:"content"`
			} `json:"zh_cn"`
		} `json:"post"`
	} `json:"content"`
}

func sendLarkTextNotify(url, title, text string) {
	if url == "" || text == "" {
		return
	}
	var data MsgData
	data.Email = ""
	data.MsgType = "post"
	data.Content.Post.ZhCn.Title = title
	data.Content.Post.ZhCn.Content = [][]MsgContent{
		{
			MsgContent{
				Tag:      "text",
				UnEscape: false,
				Text:     text,
			},
		},
	}
	_, body, errs := gorequest.New().Post(url).Timeout(time.Second * 10).SendStruct(&data).End()
	if len(errs) > 0 {
		log.Error("sendLarkTextNotify req err:", errs)
	} else {
		log.Info("sendLarkTextNotify req:", body)
	}
}
