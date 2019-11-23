package models

// 学生与家长关系，如爸爸、妈妈
type StudentRelation struct {
	Id   int64  `orm:"description(自增主键)"`
	Name string `orm:"description(关系名称)"`
}
