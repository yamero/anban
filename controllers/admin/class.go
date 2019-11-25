package admin

import (
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type ClassController struct {
	BaseController
}

func (c *ClassController) Add() {
	p := map[string]interface{}{}
	p["relation"] = false
	_, schoolList := service.GetSchoolList(p)
	c.Data["schoolList"] = schoolList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/class/add.html"
}

func (c *ClassController) DoAdd() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["school_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择所属学校", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["name"][0] == "" {
		res = utils.ResJson(0, "请输入班级名称", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.AddClass(input)
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

func (c *ClassController) Edit() {
	id := utils.Atoi64(c.Ctx.Input.Param(":id"))
	record := service.GetClassInfo(id)
	c.Data["record"] = record
	p := map[string]interface{}{}
	p["relation"] = false
	_, schoolList := service.GetSchoolList(p)
	c.Data["schoolList"] = schoolList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/class/edit.html"
}

func (c *ClassController) DoEdit() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["school_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择所属学校", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["name"][0] == "" {
		res = utils.ResJson(0, "请输入班级名称", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.EditClass(input)
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

func (c *ClassController) ShowList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	var schoolId int64
	c.Ctx.Input.Bind(&schoolId, "school_id")
	perCount, _ := beego.AppConfig.Int("percount")
	symPageCount, _ := beego.AppConfig.Int("symmetricpagecount")
	p := map[string]interface{}{}
	p["relation"] = true
	p["curPage"] = curPage
	p["perCount"] = perCount
	p["schoolId"] = schoolId
	totalCount, recordList := service.GetClassList(p)
	paginator := utils.NewPaginator(int(totalCount), perCount, symPageCount, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["recordList"] = recordList
	c.Data["schoolId"] = schoolId
	p = map[string]interface{}{}
	p["relation"] = false
	_, schoolList := service.GetSchoolList(p)
	c.Data["schoolList"] = schoolList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/class/list.html"
}