package admin

import (
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type StudentRelationController struct {
	BaseController
}

func (c *StudentRelationController) Add() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/student_relation/add.html"
}

func (c *StudentRelationController) DoAdd() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["name"][0] == "" {
		res = utils.ResJson(0, "请输入关系名称", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.AddStudentRelation(input)
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

func (c *StudentRelationController) Edit() {
	id := utils.Atoi64(c.Ctx.Input.Param(":id"))
	record := service.GetStudentRelationInfo(id)
	c.Data["record"] = record
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/student_relation/edit.html"
}

func (c *StudentRelationController) DoEdit() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["name"][0] == "" {
		res = utils.ResJson(0, "请输入关系名称", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.EditStudentRelation(input)
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

func (c *StudentRelationController) ShowList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	perCount, _ := beego.AppConfig.Int("percount")
	symPageCount, _ := beego.AppConfig.Int("symmetricpagecount")
	p := map[string]interface{}{}
	p["curPage"] = curPage
	p["perCount"] = perCount
	totalCount, recordList := service.GetStudentRelationList(p)
	paginator := utils.NewPaginator(int(totalCount), perCount, symPageCount, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["recordList"] = recordList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/student_relation/list.html"
}