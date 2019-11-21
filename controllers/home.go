package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.SetSession("isLogin", "登录")
	/*o := orm.NewOrm()
	adminUser := new(models.AdminUser)
	adminUser.Account = "admin"
	adminUser.Password = "123456"
	n, _ := o.Insert(adminUser)
	logs.Info("插入了", n, "条数据")*/
	this.Data["name"] = "黎昕"
	this.TplName = "home.html"
}