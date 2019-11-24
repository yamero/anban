package models

// 文章分类
type ArticleType struct {
	Id   int64  `orm:"description(自增主键)"`
	Name string `orm:"description(文章分类名称)"`
}
