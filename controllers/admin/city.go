package admin

import (
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type CityController struct {
	BaseController
}

func (c *CityController) AddCity() {
	p := map[string]interface{}{}
	p["level"] = 1
	_, provinceList := service.GetRegionList(p)
	c.Data["provinceList"] = provinceList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/city/add.html"
}

func (c *CityController) DoAddCity() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["name"][0] == "" {
		res = utils.ResJson(0, "市名称不能为空", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.AddRegion(input)
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

func (c *CityController) EditCity() {
	id := utils.Atoi64(c.Ctx.Input.Param(":id"))
	p := map[string]interface{}{}
	p["level"] = 1
	_, provinceList := service.GetRegionList(p)
	c.Data["provinceList"] = provinceList
	c.Data["record"] = service.GetRegionInfo(id)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/city/edit.html"
}

func (c *CityController) DoEditCity() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["name"][0] == "" {
		res = utils.ResJson(0, "市名称不能为空", "")
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

func (c *CityController) ShowCityList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	perCount, _ := beego.AppConfig.Int("percount")
	symPageCount, _ := beego.AppConfig.Int("symmetricpagecount")
	p := map[string]interface{}{}
	p["curPage"] = curPage
	p["perCount"] = perCount
	p["level"] = 2
	p["relation"] = true
	p["parent_id"] = utils.Atoi64(c.GetString("parent_id"))
	c.Data["parentId"] = p["parent_id"]
	totalCount, recordList := service.GetRegionList(p)
	paginator := utils.NewPaginator(int(totalCount), perCount, symPageCount, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["recordList"] = recordList
	p = map[string]interface{}{}
	p["level"] = 1
	_, provinceList := service.GetRegionList(p)
	c.Data["provinceList"] = provinceList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/city/list.html"
}