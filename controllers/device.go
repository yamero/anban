package controllers

import (
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type DeviceController struct {
	beego.Controller
}

func (c *DeviceController) Prepare() {
	c.EnableXSRF = false
}

func (c *DeviceController) ReadCard() {
	var res *utils.ResJsonStruct
	k := c.GetString("k")
	key := beego.AppConfig.String("swipecardrecordkey")
	if k != key {
		res = utils.ResJson(0, "非法访问", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	card := c.GetString("card")
	door := c.GetString("door")
	status := c.GetString("status")
	p := map[string]interface{}{}
	p["relation"] = true
	p["sn"] = card
	student := service.GetStudentInfo(0, p)
	if student.Id <= 0 {

	}
	go c.SaveSwipeCardRecord()
	res = utils.ResJson(1, "成功", map[string]string{"card": card, "door": door, "status": status})
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *DeviceController) SaveSwipeCardRecord() {
	card := c.GetString("card")
	door := c.GetString("door")
	status := c.GetString("status")
	logs.Info("goroutine：", card, " ", door, " ", status)
}
