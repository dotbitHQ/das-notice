package lark

import (
	"fmt"
)

//=================== payment error ===================

type ParamsDoOrderNotRefundedNotify struct {
	Count      int64
	WebhookUrl string
}

func DoOrderNotRefundedNotifyLark(p ParamsDoOrderNotRefundedNotify) {
	msg := fmt.Sprintf("未退款订单数：%d", p.Count)
	sendLarkTextNotify(p.WebhookUrl, "", msg)
}

type ParamsDoErrorHandleNotify struct {
	FuncName   string
	KeyInfo    string
	ErrInfo    string
	WebhookUrl string
}

func getLarkTextNotifyStr(funcName, keyInfo, errInfo string) string {
	msg := fmt.Sprintf(`方法名：%s
关键信息：%s
错误信息：%s`, funcName, keyInfo, errInfo)
	return msg
}

func DoHedgeDepositNotifyLark(p ParamsDoErrorHandleNotify) {
	msg := getLarkTextNotifyStr(p.FuncName, p.KeyInfo, p.ErrInfo)
	sendLarkTextNotify(p.WebhookUrl, "对冲失败", msg)
}
func DoUpdateOrderNotifyLark(p ParamsDoErrorHandleNotify) {
	msg := getLarkTextNotifyStr(p.FuncName, p.KeyInfo, p.ErrInfo)
	sendLarkTextNotify(p.WebhookUrl, "更新订单CkbHash失败", msg)
}
func DoPayCallbackNotifyLark(p ParamsDoErrorHandleNotify) {
	msg := getLarkTextNotifyStr(p.FuncName, p.KeyInfo, p.ErrInfo)
	sendLarkTextNotify(p.WebhookUrl, "支付回调发交易失败", msg)
}
func DoNodeMonitorNotifyLark(p ParamsDoErrorHandleNotify) {
	msg := getLarkTextNotifyStr(p.FuncName, p.KeyInfo, p.ErrInfo)
	sendLarkTextNotify(p.WebhookUrl, "节点监控", msg)
}

//=================== register error ===================

type ParamsDoOrderTxRejectedNotify struct {
	Account    string
	Status     int
	SinceMin   float64
	WebhookUrl string
}

func DoPreRegisterNotifyLark(p ParamsDoErrorHandleNotify) {
	msg := getLarkTextNotifyStr(p.FuncName, p.KeyInfo, p.ErrInfo)
	sendLarkTextNotify(p.WebhookUrl, "发预注册交易失败", msg)
}

func DoOrderTxRejectedNotifyLark(p ParamsDoOrderTxRejectedNotify) {
	action := "申请注册"
	if p.Status == 3 {
		action = "预注册"
	}
	msg := `> 账号：%s
> 步骤: %s
> 时间：%.2f 分钟前`
	msg = fmt.Sprintf(msg, p.Account, action, p.SinceMin)
	sendLarkTextNotify(p.WebhookUrl, "Rejected 交易监听", msg)
}
