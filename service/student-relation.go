package service

import (
	"anban/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

// 获取学校信息
func GetStudentRelationInfo(id int64) *models.StudentRelation {
	o := orm.NewOrm()
	studentRelation := &models.StudentRelation{}
	o.QueryTable("StudentRelation").Filter("id", id).One(studentRelation)
	return studentRelation
}

// 添加关系信息
func AddStudentRelation(input url.Values) (int64, error) {
	o := orm.NewOrm()
	studentRelation := &models.StudentRelation{}
	studentRelation.Name = input["name"][0]
	return o.Insert(studentRelation)
}

// 修改关系信息
func EditStudentRelation(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	p["name"] = input["name"][0]
	return o.QueryTable("StudentRelation").Filter("id", input["id"][0]).Update(p)
}

// 获取关系列表
func GetStudentRelationList(p map[string]interface{}) (int64, []*models.StudentRelation) {
	var studentRelationList []*models.StudentRelation
	o := orm.NewOrm()
	qs := o.QueryTable("StudentRelation")
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