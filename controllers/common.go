package controllers

import (
	"anban/models"
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego"
	"strconv"
)

type CommonController struct {
	beego.Controller
}

func (c *CommonController) GetRegionListByParent() {
	id := utils.Atoi64(c.GetString("id"))
	level, _ := strconv.Atoi(c.GetString("level"))
	c.Data["levelShow"] = "请选择" + models.RegionLevel[level]
	if id > 0 {
		c.Data["regionList"] = service.GetRegionListByParent(id)
	}
	c.TplName = "admin/common-option.html"
}
