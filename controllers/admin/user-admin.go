package admin

import (
	"anban/models"
	"anban/service"
	"anban/utils"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"html/template"
	"strconv"
)

// 后台管理员
// @author lixin
type UserAdminController struct {
	BaseController
}

func (c *UserAdminController) Add() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/user-admin-add.html"
}

func (c *UserAdminController) DoAdd() {
	var res *utils.ResJsonStruct
	input := c.Input()
	logs.Info("请求数据：", c.GetStrings(""))
	if input["account"][0] == "" {
		res = utils.ResJson(0, "登录名不能为空", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if len(input["password"][0]) < 6 {
		res = utils.ResJson(0, "密码至少6个字符", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["password"][0] != input["repassword"][0] {
		res = utils.ResJson(0, "两次密码输入不一致", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.AddUserAdmin(input)
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

func (c *UserAdminController) Edit() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	o := orm.NewOrm()
	userAdmin := models.UserAdmin{Id: id}
	o.Read(&userAdmin)
	c.Data["userAdmin"] = userAdmin
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/user-admin-edit.html"
}

func (c *UserAdminController) DoEdit() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["account"][0] == "" {
		res = utils.ResJson(0, "登录名不能为空", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if len(input["password"][0]) > 0 {
		if len(input["password"][0]) < 6 {
			res = utils.ResJson(0, "密码至少6个字符", "")
			c.Data["json"] = res
			c.ServeJSON()
			return
		}
		if input["password"][0] != input["repassword"][0] {
			res = utils.ResJson(0, "两次密码输入不一致", "")
			c.Data["json"] = res
			c.ServeJSON()
			return
		}
	}
	_, err := service.EditUserAdmin(input)
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

func (c *UserAdminController) ShowList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	perCount := 1
	p := map[string]interface{}{}
	p["curPage"] = curPage
	p["perCount"] = perCount
	totalCount, userAdminList := service.GetUserAdminList(p)
	paginator := utils.NewPaginator(int(totalCount), perCount, 5, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["userAdminList"] = userAdminList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/user-admin-list.html"
}

func (c *UserAdminController) StatusSwitch() {
	var res *utils.ResJsonStruct
	o := orm.NewOrm()
	p := orm.Params{"status": c.GetString("status")}
	_, err := o.QueryTable("UserAdmin").Filter("id", c.GetString("id")).Update(p)
	if err != nil {
		res = utils.ResJson(0, "操作失败", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	res = utils.ResJson(1, "操作成功", "")
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *UserAdminController) Login() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/login.html"
}

func (c *UserAdminController) DoLogin() {
	var res *utils.ResJsonStruct
	account := c.GetString("account")
	password := c.GetString("password")
	if len(account) <= 0 {
		res = utils.ResJson(0, "请输入用户名", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if len(password) <= 0 {
		res = utils.ResJson(0, "请输入密码", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	userAdmin := models.UserAdmin{}
	err := o.QueryTable("UserAdmin").Filter("account", account).One(&userAdmin)
	if err != nil {
		res = utils.ResJson(0, "登录失败，请检查用户名或密码是否正确", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if !utils.Check(password, userAdmin.Password) {
		res = utils.ResJson(0, "登录失败，请检查用户名或密码是否正确", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if userAdmin.Status == 0 {
		res = utils.ResJson(0, "登录失败，此账号已被禁用", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	c.SetSession("userAdminId", userAdmin.Id)
	res = utils.ResJson(1, "登录成功", "")
	c.Data["json"] = res
	c.ServeJSON()
	return
}

func (c *UserAdminController) Logout() {
	c.DestroySession()
	c.Redirect("/admin/login", 302)
}