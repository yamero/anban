package models

import "time"

var StudentStatus map[int]string

// 学生
type Student struct {
	Id          int64     `orm:"description(自增主键)"`
	Sn          string    `orm:"description(唯一标识，如卡号)"`
	FamilySn    string    `orm:"description(家庭编号，创建学生时随机生成)"`
	School      *School   `orm:"rel(fk);null;default(0);description(所属学校)"`
	Class       *Class    `orm:"rel(fk);null;default(0);description(所在班级)"`
	RealName    string    `orm:"description(真实姓名)"`
	IdCard      string    `orm:"description(身份证号)"`
	Status      int       `orm:"description(状态，0禁用 1启用，禁用后不再给学生家长发送作业、安全等通知)"`
	StatusShow  string    `orm:"-"`
	Created     time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	CreatedShow string    `orm:"-"`
	Updated     time.Time `orm:"auto_now;type(datetime);description(最后一次更新时间)"`
	UpdatedShow string    `orm:"-"`
}

func init() {
	StudentStatus = map[int]string{
		0: "禁用",
		1: "正常",
	}
}
