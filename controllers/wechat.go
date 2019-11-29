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

func (c *WechatController) Index() {
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	signature := c.GetString("signature")
	genSignature := utils.GetWechatSignature(timestamp, nonce)
	if signature != genSignature {
		c.Ctx.WriteString("")
		return
	}
	echoStr := c.GetString("echostr")
	c.Ctx.WriteString(echoStr)
}