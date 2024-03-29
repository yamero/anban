package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dataSource := beego.AppConfig.String("database::datasource")
	orm.RegisterDataBase("default", "mysql", dataSource)
	orm.RegisterModel(
		new(UserAdmin),
		new(UserAdminRole),
		new(Region),
		new(School),
		new(StudentRelation),
		new(Course),
		new(ArticleType),
		new(Article),
		new(Class),
		new(Student),
		new(User),
		new(WechatAccount),
	)
	orm.RunSyncdb("default", false, true)
}
