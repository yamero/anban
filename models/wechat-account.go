package models

import "time"

var WechatAccountStatus map[int]string

// 微信公众号
type WechatAccount struct {
	Id             int64     `orm:"description(自增主键)"`
	OriginalId     string    `orm:"description(公众号原始ID)"`
	Name           string    `orm:"description(公众号名称)"`
	AppId          string    `orm:"description(公众号AppId)"`
	AppSecret      string    `orm:"description(公众号AppSecret)"`
	Token          string    `orm:"description(公众号开发配置中设置的token)"`
	EncodingAesKey string    `orm:"description(公众号开发配置中设置的EncodingAESKey)"`
	NotifyUrl      string    `orm:"description(公众号开发配置中设置的通知url)"`
	Status         int       `orm:"description(状态，0禁用 1正常)"`
	StatusShow     string    `orm:"-"`
	Created        time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	CreatedShow    string    `orm:"-"`
	Updated        time.Time `orm:"auto_now;type(datetime);description(最后一次更新时间)"`
	UpdatedShow    string    `orm:"-"`
}

func init() {
	WechatAccountStatus = map[int]string{
		0: "禁用",
		1: "正常",
	}
}
