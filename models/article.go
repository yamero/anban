package models

import "time"

// 文章
type Article struct {
	Id          int64        `orm:"description(自增主键)"`
	ArticleType *ArticleType `orm:"rel(fk);null;default(0);description(所属分类)"`
	Title       string       `orm:"description(文章标题)"`
	Content     string       `orm:"type(text);null;default();description(文章内容)"`
	Sorted      int          `orm:"description(排序，数字越大越靠前)"`
	Created     time.Time    `orm:"auto_now_add;type(datetime);description(创建时间)"`
	CreatedShow string       `orm:"-"`
	Updated     time.Time    `orm:"auto_now;type(datetime);description(最后一次更新时间)"`
	UpdatedShow string       `orm:"-"`
}
