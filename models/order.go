package models

import "time"

var (
	OrderPayType map[int]string
	OrderPayStatus map[int]string
	OrderPayMethod map[int]string
)

// 订单
type Order struct {
	Id            int64     `orm:"description(自增主键)"`
	Sn            string    `orm:"description(订单编号 自动生成)"`
	User          *User     `orm:"rel(fk);null;default(0);description(对应会员)"`
	Amount        float64   `orm:"digits(12);decimals(2);description(支付金额)"`
	PayType       int       `orm:"description(付费方式 0自动扣费 1会员支付)"`
	PayStatus     int       `orm:"支付状态 0未支付 1已支付 2支付失败"`
	PayMethod     int       `orm:"description(支付方式 0微信支付 1京东钱包 2qq钱包)"`
	TransactionSn string    `orm:"description(支付成功后支付平台返回的交易编号)"`
	Created       time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	CreatedShow   string    `orm:"-"`
	Updated       time.Time `orm:"auto_now;type(datetime);description(最后一次更新时间)"`
	UpdatedShow   string    `orm:"-"`
}

func init() {

	OrderPayType = map[int]string{
		0: "自动扣费",
		1: "会员支付",
	}

	OrderPayStatus = map[int]string{
		0: "未支付",
		1: "已支付",
		2: "支付失败",
	}

	OrderPayMethod = map[int]string{
		0: "微信支付",
		1: "京东钱包",
		2: "QQ钱包",
	}

}