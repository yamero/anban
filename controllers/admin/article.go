package admin

import (
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Add() {
	_, articleTypeList := service.GetArticleTypeList(map[string]interface{}{})
	c.Data["articleTypeList"] = articleTypeList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/article/add.html"
}

func (c *ArticleController) DoAdd() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["article_type_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择文章分类", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["title"][0] == "" {
		res = utils.ResJson(0, "请输入文章标题", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["content"][0] == "" {
		res = utils.ResJson(0, "文章内容过少", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.AddArticle(input)
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

func (c *ArticleController) Edit() {
	id := utils.Atoi64(c.Ctx.Input.Param(":id"))
	record := service.GetArticleInfo(id)
	c.Data["record"] = record
	_, articleTypeList := service.GetArticleTypeList(map[string]interface{}{})
	c.Data["articleTypeList"] = articleTypeList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/article/edit.html"
}

func (c *ArticleController) DoEdit() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["article_type_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择文章分类", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["title"][0] == "" {
		res = utils.ResJson(0, "请输入文章标题", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["content"][0] == "" {
		res = utils.ResJson(0, "文章内容过少", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.EditArticle(input)
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

func (c *ArticleController) ShowList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	var articleTypeId int64
	c.Ctx.Input.Bind(&articleTypeId, "article_type_id")
	perCount, _ := beego.AppConfig.Int("percount")
	symPageCount, _ := beego.AppConfig.Int("symmetricpagecount")
	p := map[string]interface{}{}
	p["relation"] = true
	p["curPage"] = curPage
	p["perCount"] = perCount
	p["articleTypeId"] = articleTypeId
	totalCount, recordList := service.GetArticleList(p)
	paginator := utils.NewPaginator(int(totalCount), perCount, symPageCount, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["recordList"] = recordList
	c.Data["articleTypeId"] = articleTypeId
	_, articleTypeList := service.GetArticleTypeList(map[string]interface{}{})
	c.Data["articleTypeList"] = articleTypeList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/article/list.html"
}