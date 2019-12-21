package controllers

import (
	"anban/service"
	"anban/utils"
	"anban/utils/Device"
	"anban/utils/wechat"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"net/url"
)

type MyTestController struct {
	beego.Controller
}

func (c *MyTestController) Prepare() {
	c.EnableXSRF = false;
}

func (c *MyTestController) TestJson() {
	var res *utils.ResJsonStruct
	var p Device.DeviceLogin
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &p); err != nil {
		res = utils.ResJson(0, "解析失败", err.Error())
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	for _, p1 := range p.Place {
		fmt.Printf("Tid：%s Name：%s\n", p1.Tid, p1.Name)
		for _, p2 := range p1.Child {
			fmt.Printf("	Tid：%s Name：%s\n", p2.Tid, p2.Name)
		}
	}
	res = utils.ResJson(0, "解析成功", p.Place)
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *MyTestController) TestXml() {
	var res *utils.ResJsonStruct
	var p wechat.EventMsg
	if err := xml.Unmarshal(c.Ctx.Input.RequestBody, &p); err != nil {
		res = utils.ResJson(0, "解析失败", err.Error())
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	fmt.Printf("ToUserName：%s FromUserName：%s CreateTime：%d MsgType：%s Event：%s EventKey：%s Ticket：%s\n",
		p.ToUserName, p.FromUserName, p.CreateTime, p.MsgType, p.Event, p.EventKey, p.Ticket)
	res = utils.ResJson(0, "解析成功", p)
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *MyTestController) InitAdminUser() {
	input := url.Values{
		"account": {"admin"},
		"password": {"anban123"},
		"status": {"1"},
	}
	service.AddUserAdmin(input)
	c.Ctx.WriteString("管理员初始化成功，密码:anban123，登录后请务必修改为更复杂的密码")
}
