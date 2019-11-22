package models

import "time"

var UserAdminStatus map[int]string

// 后台管理员
type UserAdmin struct {
	Id            int64          `orm:"description(自增主键)"`
	Account       string         `orm:"description(登录账号)"`
	Password      string         `orm:"description(登录密码)"`
	Mobile        string         `orm:"description(手机号)"`
	RealName      string         `orm:"description(姓名)"`
	Email         string         `orm:"description(邮箱)"`
	UserAdminRole *UserAdminRole `orm:"null;default(0);rel(fk);description(对应角色)"`
	Status        int            `orm:"description(状态 0禁用 1正常)"`
	StatusShow    string         `orm:"-"`
	Created       time.Time      `orm:"auto_now_add;type(datetime);description(创建时间)"`
	CreatedShow   string         `orm:"-"`
	Updated       time.Time      `orm:"auto_now;type(datetime);description(最后一次更新时间)"`
	UpdatedShow   string         `orm:"-"`
}

func init() {
	UserAdminStatus = map[int]string{
		0: "禁用",
		1: "正常",
	}
}
