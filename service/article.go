package service

import (
	"anban/models"
	"anban/utils"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strconv"
	"time"
)

// 获取文章信息
func GetArticleInfo(id int64) *models.Article {
	o := orm.NewOrm()
	article := &models.Article{}
	o.QueryTable("Article").RelatedSel().Filter("id", id).One(article)
	article.CreatedShow = article.Created.Format("2006-01-02 15:04:05")
	article.UpdatedShow = article.Updated.Format("2006-01-02 15:04:05")
	return article
}

// 添加文章信息
func AddArticle(input url.Values) (int64, error) {
	o := orm.NewOrm()
	articleTypeId := utils.Atoi64(input["article_type_id"][0])
	articleType := &models.ArticleType{Id: articleTypeId}
	o.Read(articleType)
	article := &models.Article{}
	article.ArticleType = articleType
	article.Title = input["title"][0]
	article.Content = input["content"][0]
	article.Sorted, _ = strconv.Atoi(input["sorted"][0])
	return o.Insert(article)
}

// 修改文章信息
func EditArticle(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	p["article_type_id"] = input["article_type_id"][0]
	p["title"] = input["title"][0]
	p["sorted"] = input["sorted"][0]
	p["content"] = input["content"][0]
	p["updated"] = time.Now().Format("2006-01-02 15:04:05")
	return o.QueryTable("Article").Filter("id", input["id"][0]).Update(p)
}

// 获取文章列表
func GetArticleList(p map[string]interface{}) (int64, []*models.Article) {
	var articleList []*models.Article
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	relation, _ := p["relation"].(bool)
	if relation {
		qs = qs.RelatedSel()
	}
	articleTypeId, _ := p["articleTypeId"].(int64)
	if articleTypeId > 0 {
		qs = qs.Filter("article_type_id", articleTypeId)
	}
	totalCount, _ := qs.Count()
	curPage, ok := p["curPage"]
	perCount, _ := p["perCount"]
	if ok {
		start := perCount.(int) * (curPage.(int) - 1)
		qs = qs.Limit(perCount, start)
	}
	qs.OrderBy("-sorted").All(&articleList)
	for _, article := range articleList {
		article.CreatedShow = article.Created.Format("2006-01-02 15:04:05")
		article.UpdatedShow = article.Updated.Format("2006-01-02 15:04:05")
	}
	return totalCount, articleList
}