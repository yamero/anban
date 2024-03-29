package routers

import (
	"anban/controllers"
	"anban/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	// 虽然beego已经支持注解路由，但注解路由会用到反射机制，降低性能，所以这里还是使用注册路由

	// 初始化第一个后台管理员
	beego.Router("/init_admin_user", &controllers.MyTestController{}, "get:InitAdminUser")

	// 测试路由
	beego.Router("/test_json", &controllers.MyTestController{}, "post:TestJson")
	beego.Router("/test_xml", &controllers.MyTestController{}, "post:TestXml")

	// 安伴公众号接收微信消息路由
	beego.Router("/wechat/anban.html", &controllers.WechatController{}, "get:AnBan;post:AnBan")

	// 前端路由
	beego.Router("/", &admin.IndexController{}, "get:Index")
    beego.Router("/main", &controllers.MainController{})

	// 后端路由
	beego.Router("/common/get_region_by_parent", &controllers.CommonController{}, "post:GetRegionListByParent")
	beego.Router("/common/get_class_by_school", &controllers.CommonController{}, "post:GetClassBySchool")
	beego.Router("/admin", &admin.IndexController{}, "get:Index")
	beego.Router("/admin/login", &admin.UserAdminController{}, "get:Login;post:DoLogin")
	beego.Router("/admin/logout", &admin.UserAdminController{}, "get:Logout")
	beego.Router("/admin/welcome", &admin.IndexController{}, "get:Welcome")
	beego.Router("/admin/editor_upload", &admin.UploadController{}, "post:Editor")

	beego.Router("/admin/user_admin", &admin.UserAdminController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/user_admin/add", &admin.UserAdminController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/user_admin/edit/?:id", &admin.UserAdminController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/user_admin/status_switch", &admin.UserAdminController{}, "post:StatusSwitch")

	beego.Router("/admin/province", &admin.ProvinceController{}, "get:ShowProvinceList;post:ShowProvinceList")
	beego.Router("/admin/province/add", &admin.ProvinceController{}, "get:AddProvince;post:DoAddProvince")
	beego.Router("/admin/province/edit/?:id", &admin.ProvinceController{}, "get:EditProvince;post:DoEditProvince")

	beego.Router("/admin/city", &admin.CityController{}, "get:ShowCityList;post:ShowCityList")
	beego.Router("/admin/city/add", &admin.CityController{}, "get:AddCity;post:DoAddCity")
	beego.Router("/admin/city/edit/?:id", &admin.CityController{}, "get:EditCity;post:DoEditCity")

	beego.Router("/admin/district", &admin.DistrictController{}, "get:ShowDistrictList;post:ShowDistrictList")
	beego.Router("/admin/district/add", &admin.DistrictController{}, "get:AddDistrict;post:DoAddDistrict")
	beego.Router("/admin/district/edit/?:id", &admin.DistrictController{}, "get:EditDistrict;post:DoEditDistrict")

	beego.Router("/admin/school", &admin.SchoolController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/school/add", &admin.SchoolController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/school/edit/?:id", &admin.SchoolController{}, "get:Edit;post:DoEdit")

	beego.Router("/admin/class", &admin.ClassController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/class/add", &admin.ClassController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/class/edit/?:id", &admin.ClassController{}, "get:Edit;post:DoEdit")

	beego.Router("/admin/student_relation", &admin.StudentRelationController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/student_relation/add", &admin.StudentRelationController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/student_relation/edit/?:id", &admin.StudentRelationController{}, "get:Edit;post:DoEdit")

	beego.Router("/admin/course", &admin.CourseController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/course/add", &admin.CourseController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/course/edit/?:id", &admin.CourseController{}, "get:Edit;post:DoEdit")

	beego.Router("/admin/article_type", &admin.ArticleTypeController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/article_type/add", &admin.ArticleTypeController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/article_type/edit/?:id", &admin.ArticleTypeController{}, "get:Edit;post:DoEdit")

	beego.Router("/admin/article", &admin.ArticleController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/article/edit/?:id", &admin.ArticleController{}, "get:Edit;post:DoEdit")

	beego.Router("/admin/student", &admin.StudentController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/student/add", &admin.StudentController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/student/edit/?:id", &admin.StudentController{}, "get:Edit;post:DoEdit")

	beego.Router("/admin/user", &admin.UserController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/user/edit/?:id", &admin.UserController{}, "get:Edit;post:DoEdit")

	beego.Router("/admin/wechat_account", &admin.WechatAccountController{}, "get:ShowList;post:ShowList")
	beego.Router("/admin/wechat_account/add", &admin.WechatAccountController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/wechat_account/edit/?:id", &admin.WechatAccountController{}, "get:Edit;post:DoEdit")

	// 与设备通信
	beego.Router("/read_card", &controllers.DeviceController{}, "post:ReadCard")
}
