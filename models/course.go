package models

// 学校科目，如语文、数字
type Course struct {
	Id   int64  `orm:"description(自增主键)"`
	Name string `orm:"description(科目名称)"`
}