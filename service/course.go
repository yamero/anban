package service

import (
	"anban/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

// 获取科目信息
func GetCourseInfo(id int64) *models.Course {
	o := orm.NewOrm()
	studentRelation := &models.Course{}
	o.QueryTable("Course").Filter("id", id).One(studentRelation)
	return studentRelation
}

// 添加科目信息
func AddCourse(input url.Values) (int64, error) {
	o := orm.NewOrm()
	studentRelation := &models.Course{}
	studentRelation.Name = input["name"][0]
	return o.Insert(studentRelation)
}

// 修改科目信息
func EditCourse(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	p["name"] = input["name"][0]
	return o.QueryTable("Course").Filter("id", input["id"][0]).Update(p)
}

// 获取科目列表
func GetCourseList(p map[string]interface{}) (int64, []*models.Course) {
	var studentRelationList []*models.Course
	o := orm.NewOrm()
	qs := o.QueryTable("Course")
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