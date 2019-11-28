package models

import "time"

type Order struct {
	Id int64 `orm:"description(自增主键)"`
	User *User `orm:"rel(fk);null;default(0);description(对应会员)"`
	PayStatus int `orm:"支付状态 0未支付 1已支付 2支付失败"`
	Created     time.Time    `orm:"auto_now_add;type(datetime);description(创建时间)"`
	CreatedShow string       `orm:"-"`
	Updated     time.Time    `orm:"auto_now;type(datetime);description(最后一次更新时间)"`
	UpdatedShow string       `orm:"-"`
}
