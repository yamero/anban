package models

import "time"

var (
	UserSubscribeStatus map[int]string
	UserStatus          map[int]string
	UserIdentity        map[int]string
)

// 会员
type User struct {
	Id                  int64      `orm:"description(自增主键)"`
	FamilySn            string     `orm:"description(家庭编号，来源于学生的FamilySn)"`
	Students            []*Student `orm:"rel(m2m)"`
	OpenId              string     `orm:"description(微信公众平台openid)"`
	UnionId             string     `orm:"description(微信开放平台unionid)"`
	RealName            string     `orm:"description(姓名)"`
	Phone               string     `orm:"description(电话)"`
	IdCard              string     `orm:"description(身份证号)"`
	ShortName           string     `orm:"description(简称，如李老师)"`
	NickName            string     `orm:"description(微信昵称)"`
	Status              int        `orm:"description(状态，0禁用 1正常)"`
	StatusShow          string     `orm:"-"`
	Identity            int        `orm:"description(身份 0未知 1家长 2老师 3家长&老师 4校领导)"`
	IdentityShow        string     `orm:"-"`
	SubscribeStatus     int        `orm:"description(关注公众号状态 0未关注 1已关注 2已取消关注)"`
	SubscribeStatusShow string     `orm:"-"`
	SubscribeTime       time.Time  `orm:"auto_now_add;type(datetime);description(关注公众号时间)"`
	SubscribeTimeShow   string     `orm:"-"`
	UnsubscribeTime     time.Time  `orm:"auto_now_add;type(datetime);description(取消关注公众号时间)"`
	UnsubscribeTimeShow string     `orm:"-"`
	Created             time.Time  `orm:"auto_now_add;type(datetime);description(创建时间)"`
	CreatedShow         string     `orm:"-"`
	Updated             time.Time  `orm:"auto_now;type(datetime);description(最后一次更新时间)"`
	UpdatedShow         string     `orm:"-"`
}

func init() {

	UserSubscribeStatus = map[int]string{
		0: "未关注",
		1: "已关注",
		2: "已取消关注",
	}

	UserStatus = map[int]string{
		0: "禁用",
		1: "正常",
	}

	UserIdentity = map[int]string{
		0: "未知",
		1: "家长",
		2: "老师",
		3: "家长&老师",
		4: "校领导",
	}

}
