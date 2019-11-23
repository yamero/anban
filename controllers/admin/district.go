package admin

import (
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type DistrictController struct {
	BaseController
}

func (c *DistrictController) AddDistrict() {
	p := map[string]interface{}{}
	p["level"] = 1
	_, provinceList := service.GetRegionList(p)
	c.Data["provinceList"] = provinceList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/district-add.html"
}

func (c *DistrictController) DoAddDistrict() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["name"][0] == "" {
		res = utils.ResJson(0, "区/县名称不能为空", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["parent_id"][0] == "0" {
		res = utils.ResJson(0, "请选择所属市", "")
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

func (c *DistrictController) EditDistrict() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	p := map[string]interface{}{}
	p["level"] = 1
	_, provinceList := service.GetRegionList(p)
	p["level"] = 2
	_, cityList := service.GetRegionList(p)
	c.Data["provinceList"] = provinceList
	c.Data["cityList"] = cityList
	c.Data["record"] = service.GetRegionInfo(id)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/district-edit.html"
}

func (c *DistrictController) DoEditDistrict() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["name"][0] == "" {
		res = utils.ResJson(0, "区/县名称不能为空", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["parent_id"][0] == "0" {
		res = utils.ResJson(0, "请选择所属市", "")
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

func (c *DistrictController) ShowDistrictList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	perCount, _ := beego.AppConfig.Int("percount")
	symPageCount, _ := beego.AppConfig.Int("symmetricpagecount")
	p := map[string]interface{}{}
	p["curPage"] = curPage
	p["perCount"] = perCount
	p["level"] = 3
	p["relation"] = true
	p["parent_id"] = utils.Atoi64(c.GetString("parent_id"))
	c.Data["parentId"] = p["parent_id"]
	provinceId := utils.Atoi64(c.GetString("province_id"))
	c.Data["provinceId"] = provinceId
	totalCount, recordList := service.GetRegionList(p)
	paginator := utils.NewPaginator(int(totalCount), perCount, symPageCount, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["recordList"] = recordList
	p = map[string]interface{}{}
	p["level"] = 1
	_, provinceList := service.GetRegionList(p)
	c.Data["provinceList"] = provinceList
	if provinceId > 0 {
		p = map[string]interface{}{}
		p["level"] = 2
		p["parent_id"] = provinceId
		_, cityList := service.GetRegionList(p)
		c.Data["cityList"] = cityList
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/district-list.html"
}