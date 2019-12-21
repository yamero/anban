package admin

import (
	"anban/service"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	noLogin := map[string]string{
		"UserAdminController/Login": "1",
		"UserAdminController/DoLogin": "1",
		"UserAdminController/Logout": "1",
	}
	ctl, act := c.GetControllerAndAction()
	p := fmt.Sprintf("%s/%s", ctl, act)
	_, ok := noLogin[p]
	if !ok {
		userAdminId := c.GetSession("userAdminId")
		if userAdminId == nil {
			c.Redirect("/admin/login", 302)
			return
		}
		p := map[string]interface{}{}
		p["relation"] = false
		c.Data["loginUserAdmin"] = service.GetUserAdminInfo(userAdminId.(int64), p)
		c.Data["currentTime"] = time.Now().Format("2006-01-02 15:04")
	}
}