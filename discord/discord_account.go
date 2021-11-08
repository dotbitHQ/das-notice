package discord

import (
	"fmt"
	"github.com/scorpiotzh/mylog"
	"math/rand"
	"strings"
)

var log = mylog.NewLogger("discord", mylog.LevelDebug)

//=================== payment account ===================

// RegisterAccountNotice 注册账户
func RegisterAccountNotice(webHookUrl, content string) {
	log.Info("RegisterAccountNotice:", content)

	err := sendNotifyDiscordContent(webHookUrl, content)
	if err != nil {
		log.Error("RegisterAccountNotice err: ", err.Error())
	}
}

//=================== transaction account ===================

type BuyAccountNoticeParams struct {
	Account    string
	PriceCkb   string
	PriceUsd   string
	WebhookUrl string
}

func numFormat(str string) string {
	strArr := strings.Split(str, ".")
	length := len(strArr[0])
	if length < 4 {
		return str
	}
	resStr := strArr[0]
	count := (len(resStr) - 1) / 3
	for i := 0; i < count; i++ {
		resStr = resStr[:length-(i+1)*3] + "," + resStr[length-(i+1)*3:]
	}
	if len(strArr) > 1 {
		resStr += "." + strArr[1]
	}
	return resStr
}

// BuyAccountNotice 购买账户
func BuyAccountNotice(p *BuyAccountNoticeParams) {
	log.Info("BuyAccountNotice:", p.Account, p.PriceCkb, p.PriceUsd)

	title := fmt.Sprintf("** %s ** bought for %s CKB($%s).", p.Account, numFormat(p.PriceCkb), numFormat(p.PriceUsd))

	err := sendNotifyDiscordContent(p.WebhookUrl, title)
	if err != nil {
		log.Error("BuyAccountNotice err: ", err.Error())
	}
}

type StartAccountNoticeParams struct {
	Account     string
	PriceCkb    string
	PriceUsd    string
	EmbedUrl    string
	Description string
	WebhookUrl  string
}

// StartAccountNotice 上架账户，修改上架账户
func StartAccountNotice(p *StartAccountNoticeParams) {
	log.Info("StartAccountNotice:", p.Account, p.PriceCkb, p.PriceUsd)

	title := fmt.Sprintf("** %s **listed for %s CKB($%s). Buy Now", p.Account, numFormat(p.PriceCkb), numFormat(p.PriceUsd))
	embed := Embed{
		Title:       title,
		URL:         p.EmbedUrl,
		Description: p.Description,
		Color:       rand.Int63n(16777215),
	}
	err := sendNotifyDiscordEmbeds(p.WebhookUrl, embed)
	if err != nil {
		log.Error("StartAccountNotice err: ", err.Error())
	}
}
