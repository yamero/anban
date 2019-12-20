package admin

import (
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type WechatAccountController struct {
	BaseController
}

func (c *WechatAccountController) Add() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/wechat_account/add.html"
}

func (c *WechatAccountController) DoAdd() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["school_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择学校", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["class_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择班级", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["sn"][0] == "" {
		res = utils.ResJson(0, "请输入唯一标识", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["real_name"][0] == "" {
		res = utils.ResJson(0, "请输入姓名", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["id_card"][0] == "" {
		res = utils.ResJson(0, "请输入身份证号", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.AddWechatAccount(input)
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

func (c *WechatAccountController) Edit() {
	id := utils.Atoi64(c.Ctx.Input.Param(":id"))
	p := map[string]interface{}{}
	p["relation"] = false
	record := service.GetWechatAccountInfo(id, p)
	c.Data["record"] = record
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/wechat_account/edit.html"
}

func (c *WechatAccountController) DoEdit() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["school_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择学校", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["class_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择班级", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["sn"][0] == "" {
		res = utils.ResJson(0, "请输入唯一标识", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["real_name"][0] == "" {
		res = utils.ResJson(0, "请输入姓名", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["id_card"][0] == "" {
		res = utils.ResJson(0, "请输入身份证号", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.EditWechatAccount(input)
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

func (c *WechatAccountController) ShowList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	perCount, _ := beego.AppConfig.Int("percount")
	symPageCount, _ := beego.AppConfig.Int("symmetricpagecount")
	p := map[string]interface{}{}
	p["relation"] = false
	p["curPage"] = curPage
	p["perCount"] = perCount
	totalCount, recordList := service.GetWechatAccountList(p)
	paginator := utils.NewPaginator(int(totalCount), perCount, symPageCount, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["recordList"] = recordList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/wechat_account/list.html"
}