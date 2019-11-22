package admin

import (
	"anban/models"
	"anban/service"
	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
}

func (c *CommonController) GetRegionListByParent() {
	id, _ := c.GetInt64("id")
	level, _ := c.GetInt("level")
	c.Data["levelShow"] = "请选择" + models.RegionLevel[level]
	c.Data["regionList"] = service.GetRegionListByParent(id)
	c.TplName = "admin/common-option.html"
}
