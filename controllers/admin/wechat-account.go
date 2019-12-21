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
	if input["name"][0] == "" {
		res = utils.ResJson(0, "请输入公众号名称", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["original_id"][0] == "" {
		res = utils.ResJson(0, "请输入公众号原始ID", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["app_id"][0] == "" {
		res = utils.ResJson(0, "请输入公众号AppId", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["app_secret"][0] == "" {
		res = utils.ResJson(0, "请输入公众号AppSecret", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["token"][0] == "" {
		res = utils.ResJson(0, "请输入公众号Token", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["encoding_aes_key"][0] == "" {
		res = utils.ResJson(0, "请输入公众号EncodingAESKey", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["notify_url"][0] == "" {
		res = utils.ResJson(0, "请输入公众号通知url", "")
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
	if input["name"][0] == "" {
		res = utils.ResJson(0, "请输入公众号名称", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["original_id"][0] == "" {
		res = utils.ResJson(0, "请输入公众号原始ID", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["app_id"][0] == "" {
		res = utils.ResJson(0, "请输入公众号AppId", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["app_secret"][0] == "" {
		res = utils.ResJson(0, "请输入公众号AppSecret", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["token"][0] == "" {
		res = utils.ResJson(0, "请输入公众号Token", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["encoding_aes_key"][0] == "" {
		res = utils.ResJson(0, "请输入公众号EncodingAESKey", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["notify_url"][0] == "" {
		res = utils.ResJson(0, "请输入公众号通知url", "")
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