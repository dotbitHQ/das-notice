package lark

import (
	"fmt"
	"testing"
)

func TestDoBlockParserNotifyLark(t *testing.T) {
	DoBlockParserNotifyLark(ParamsDoBlockParserNotify{
		Action:      "apply_register",
		BlockNumber: 5718189,
		Hash:        "0xe9116d651c371662b6e29e2102422e23f90656b8619df82c48b782ff4db43a37",
		WebhookUrl:  "https://open.larksuite.com/open-apis/bot/v2/hook/a5225cf9-7865-486e-917d-2284b0395e98",
	})
}

func TestSendLarkTextNotify(t *testing.T) {
	msg := fmt.Sprintf("The count of not serial account is %d", 2)
	err := SendLarkTextNotify(
		"https://open.larksuite.com/open-apis/bot/v2/hook/a5225cf9-7865-486e-917d-2284b0395e98",
		"Account not serial",
		msg,
	)
	if err != nil {
		log.Error("SendLarkTextNotify err: ", err.Error())
	}
}
