package models

// 班级
type Class struct {
	Id     int64   `orm:"description(自增主键)"`
	School *School `orm:"rel(fk);null;default(0);description(所属学校)"`
	Name   string  `orm:"description(班级名称)"`
	Sn     string  `orm:"description(班级编号，添加时生成的)"`
}
