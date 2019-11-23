package admin

import (
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type SchoolController struct {
	BaseController
}

func (c *SchoolController) Add() {
	p := map[string]interface{}{}
	p["level"] = 1
	_, provinceList := service.GetRegionList(p)
	c.Data["provinceList"] = provinceList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/school-add.html"
}

func (c *SchoolController) DoAdd() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["province_id"][0] == "0" {
		res = utils.ResJson(0, "请选择所属省", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["city_id"][0] == "0" {
		res = utils.ResJson(0, "请选择所属市", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["district_id"][0] == "0" {
		res = utils.ResJson(0, "请选择所属区/县", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["name"][0] == "" {
		res = utils.ResJson(0, "请输入学校名称", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.AddSchool(input)
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

func (c *SchoolController) Edit() {
	id := utils.Atoi64(c.Ctx.Input.Param(":id"))
	p := map[string]interface{}{}
	p["level"] = 1
	_, provinceList := service.GetRegionList(p)
	c.Data["provinceList"] = provinceList
	record := service.GetSchoolInfo(id)
	c.Data["record"] = record
	p["level"] = 2
	p["parent_id"] = record.Province.Id
	_, cityList := service.GetRegionList(p)
	c.Data["cityList"] = cityList
	p["level"] = 3
	p["parent_id"] = record.City.Id
	_, districtList := service.GetRegionList(p)
	c.Data["districtList"] = districtList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/school-edit.html"
}

func (c *SchoolController) DoEdit() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["province_id"][0] == "0" {
		res = utils.ResJson(0, "请选择所属省", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["city_id"][0] == "0" {
		res = utils.ResJson(0, "请选择所属市", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["district_id"][0] == "0" {
		res = utils.ResJson(0, "请选择所属区/县", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["name"][0] == "" {
		res = utils.ResJson(0, "请输入学校名称", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.EditSchool(input)
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

func (c *SchoolController) ShowList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	perCount, _ := beego.AppConfig.Int("percount")
	symPageCount, _ := beego.AppConfig.Int("symmetricpagecount")
	p := map[string]interface{}{}
	p["curPage"] = curPage
	p["perCount"] = perCount
	p["relation"] = true
	provinceId := utils.Atoi64(c.GetString("province_id"))
	c.Data["provinceId"] = provinceId
	p["provinceId"] = provinceId
	cityId := utils.Atoi64(c.GetString("city_id"))
	c.Data["cityId"] = cityId
	p["cityId"] = cityId
 	districtId := utils.Atoi64(c.GetString("district_id"))
	c.Data["districtId"] = districtId
	p["districtId"] = districtId
	name := c.GetString("name")
	c.Data["name"] = name
	p["name"] = name
	totalCount, recordList := service.GetSchoolList(p)
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
	if cityId > 0 {
		p = map[string]interface{}{}
		p["level"] = 3
		p["parent_id"] = cityId
		_, districtList := service.GetRegionList(p)
		c.Data["districtList"] = districtList
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/school-list.html"
}