package lark

import (
	"fmt"
	"testing"
	"time"
)

func TestSendLarkTextNotify(t *testing.T) {
	msg := "> Block number：%d\n> Action：%s\n> Timestamp：%s\n> Transaction hash：%s"
	msg = fmt.Sprintf(msg, 3763163, "edit_offer", time.Now().Format("2006-01-02 15:04:05"), "0x31db70391cb1b541cccbd5146e77870c095b48bc4dc8d763f222c9b0afe19424")
	err := SendLarkTextNotify(
		"https://open.larksuite.com/open-apis/bot/v2/hook/a5225cf9-7865-486e-917d-2284b0395e98",
		"Block monitor",
		msg,
	)
	if err != nil {
		t.Log("SendLarkTextNotify err:", err.Error())
	}
}

func TestSendLarkTextAllNotify(t *testing.T) {
	msg := fmt.Sprintf("\nThe count of not serial account is %d", 2)
	err := SendLarkTextAllNotify(
		"https://open.larksuite.com/open-apis/bot/v2/hook/a5225cf9-7865-486e-917d-2284b0395e98",
		"Account not serial",
		msg,
	)
	if err != nil {
		t.Log("SendLarkTextAllNotify err:", err.Error())
	}
}
