package routers

import (
	"anban/controllers"
	"anban/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	// 虽然beego已经支持注解路由，但注解路由会用到反射机制，降低性能，所以这里还是使用注册路由

	// 前端路由
	beego.Router("/", &controllers.HomeController{})
    beego.Router("/main", &controllers.MainController{})

	// 后端路由
	beego.Router("/admin", &admin.IndexController{}, "get:Index")
	beego.Router("/admin/login", &admin.UserAdminController{}, "get:Login;post:DoLogin")
	beego.Router("/admin/logout", &admin.UserAdminController{}, "get:Logout")
	beego.Router("/admin/welcome", &admin.IndexController{}, "get:Welcome")
	beego.Router("/admin/user_admin", &admin.UserAdminController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/user_admin/add", &admin.UserAdminController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/user_admin/edit/?:id", &admin.UserAdminController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/user_admin/status_switch", &admin.UserAdminController{}, "post:StatusSwitch")
}
