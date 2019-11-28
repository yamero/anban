package admin

import (
	"anban/models"
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type UserController struct {
	BaseController
}

func (c *UserController) Edit() {
	id := utils.Atoi64(c.Ctx.Input.Param(":id"))
	p := map[string]interface{}{}
	p["relation"] = false
	record := service.GetUserInfo(id, p)
	c.Data["record"] = record
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/user/edit.html"
}

func (c *UserController) DoEdit() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if len(input["real_name"][0]) <= 0 {
		res = utils.ResJson(0, "请输入姓名", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if len(input["id_card"][0]) <= 0 {
		res = utils.ResJson(0, "请输入身份证号", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if len(input["phone"][0]) <= 0 {
		res = utils.ResJson(0, "请输入电话", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.EditUser(input)
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

func (c *UserController) ShowList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	var realName string
	var idCard string
	var phone string
	var identity int
	c.Ctx.Input.Bind(&realName, "real_name")
	c.Ctx.Input.Bind(&idCard, "id_card")
	c.Ctx.Input.Bind(&phone, "phone")
	c.Ctx.Input.Bind(&identity, "identity")
	perCount, _ := beego.AppConfig.Int("percount")
	symPageCount, _ := beego.AppConfig.Int("symmetricpagecount")
	p := map[string]interface{}{}
	p["relation"] = false
	p["curPage"] = curPage
	p["perCount"] = perCount
	p["realName"] = realName
	p["idCard"] = idCard
	p["phone"] = phone
	p["identity"] = identity
	totalCount, recordList := service.GetUserList(p)
	paginator := utils.NewPaginator(int(totalCount), perCount, symPageCount, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["recordList"] = recordList
	c.Data["realName"] = realName
	c.Data["idCard"] = idCard
	c.Data["phone"] = phone
	c.Data["identity"] = identity
	c.Data["identityList"] = models.UserIdentity
 	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/user/list.html"
}