package models

import "time"

type Student struct {
	Id int64 `orm:"description(自增主键)"`
	Sn string `orm:"description(唯一标识，如卡号)"`
	RealName string `orm:"description(真实姓名)"`
	IdCard string `orm:"description(身份证号)"`
	Created       time.Time      `orm:"auto_now_add;type(datetime);description(创建时间)"`
	CreatedShow   string         `orm:"-"`
	Updated       time.Time      `orm:"auto_now;type(datetime);description(最后一次更新时间)"`
	UpdatedShow   string         `orm:"-"`
}
