package controllers

import (
	"anban/utils"
	"github.com/astaxie/beego"
)

type DeviceController struct {
	beego.Controller
}

func (c *DeviceController) Prepare() {
	c.EnableXSRF = false
}

func (c *DeviceController) ReadCard() {
	var res *utils.ResJsonStruct
	card := c.GetString("card")
	door := c.GetString("door")
	status := c.GetString("status")
	res = utils.ResJson(1, "成功", map[string]string{"card": card, "door": door, "status": status})
	c.Data["json"] = res
	c.ServeJSON()
}
