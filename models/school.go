package models

type School struct {
	Id int64 `orm:"description(自增主键)"`
	Name string `orm:"description(学校名称)"`
	Province *Region `orm:"null;default(0);rel(fk);description(所属省)"`
	City *Region `orm:"null;default(0);rel(fk);description(所属市)"`
	District *Region `orm:"null;default(0);rel(fk);description(所属县区)"`
}