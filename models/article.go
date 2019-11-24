package models

// 文章
type Article struct {
	Id int64 `orm:"description(自增主键)"`
	Title string `orm:"description(文章标题)"`
	Content string `orm:"type(text);null;default();description(文章内容)"`
}
