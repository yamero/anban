package controllers

import (
	"anban/utils"
	"github.com/astaxie/beego"
)

// 接收微信通知
type WechatController struct {
	beego.Controller
}

func (c *WechatController) Prepare() {
	c.EnableXSRF = false
}

// 安伴公众号
func (c *WechatController) AnBan() {
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	signature := c.GetString("signature")
	genSignature := utils.GetWechatSignature(timestamp, nonce)
	if signature != genSignature { // 验证是否是微信发过来的消息，如果不是，直接返回
		c.Ctx.WriteString("")
		return
	}
	echoStr := c.GetString("echostr")
	c.Ctx.WriteString(echoStr)
}