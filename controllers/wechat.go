package controllers

import (
	"anban/utils/wechat"
	"encoding/xml"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	genSignature := wechat.GetWechatSignature(timestamp, nonce)
	if signature != genSignature { // 验证是否是微信发过来的消息，如果不是，直接返回
		c.Ctx.WriteString("")
		return
	}
	var p wechat.EventMsg
	if err := xml.Unmarshal(c.Ctx.Input.RequestBody, &p); err != nil {
		logs.Info("微信返回出错：", err)
	}
	logs.Info("微信返回：", p)
	c.Ctx.WriteString("success")
}