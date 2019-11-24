package service

import (
	"anban/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

// 获取文章分类信息
func GetArticleTypeInfo(id int64) *models.ArticleType {
	o := orm.NewOrm()
	studentRelation := &models.ArticleType{}
	o.QueryTable("ArticleType").Filter("id", id).One(studentRelation)
	return studentRelation
}

// 添加文章分类信息
func AddArticleType(input url.Values) (int64, error) {
	o := orm.NewOrm()
	studentRelation := &models.ArticleType{}
	studentRelation.Name = input["name"][0]
	return o.Insert(studentRelation)
}

// 修改文章分类信息
func EditArticleType(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	p["name"] = input["name"][0]
	return o.QueryTable("ArticleType").Filter("id", input["id"][0]).Update(p)
}

// 获取文章分类列表
func GetArticleTypeList(p map[string]interface{}) (int64, []*models.ArticleType) {
	var studentRelationList []*models.ArticleType
	o := orm.NewOrm()
	qs := o.QueryTable("ArticleType")
	totalCount, _ := qs.Count()
	curPage, ok := p["curPage"]
	perCount, _ := p["perCount"]
	if ok {
		start := perCount.(int) * (curPage.(int) - 1)
		qs = qs.Limit(perCount, start)
	}
	qs.All(&studentRelationList)
	return totalCount, studentRelationList
}