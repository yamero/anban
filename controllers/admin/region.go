package admin

import (
	"anban/service"
	"anban/utils"
	"html/template"
	"strconv"
)

type RegionController struct {
	BaseController
}

func (c *RegionController) AddProvince() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/province-add.html"
}

func (c *RegionController) DoAddProvince() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["name"][0] == "" {
		res = utils.ResJson(0, "省名称不能为空", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.AddProvince(input)
	if err != nil {
		res = utils.ResJson(0, "添加失败", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	res = utils.ResJson(1, "添加成功", "")
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *RegionController) EditProvince() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	c.Data["record"] = service.GetRegionInfo(id)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/province-edit.html"
}

func (c *RegionController) DoEditProvince() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["name"][0] == "" {
		res = utils.ResJson(0, "省名称不能为空", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.EditRegion(input)
	if err != nil {
		res = utils.ResJson(0, "修改失败", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	res = utils.ResJson(1, "修改成功", "")
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *RegionController) ShowProvinceList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	perCount := 1
	p := map[string]interface{}{}
	p["curPage"] = curPage
	p["perCount"] = perCount
	totalCount, recordList := service.GetProvinceList(p)
	paginator := utils.NewPaginator(int(totalCount), perCount, 5, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["recordList"] = recordList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/province-list.html"
}