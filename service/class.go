package service

import (
	"anban/models"
	"anban/utils"
	"github.com/astaxie/beego/orm"
	"net/url"
)

// 获取班级信息
func GetClassInfo(id int64) *models.Class {
	o := orm.NewOrm()
	class := &models.Class{}
	o.QueryTable("Class").RelatedSel().Filter("id", id).One(class)
	return class
}

// 添加班级信息
func AddClass(input url.Values) (int64, error) {
	o := orm.NewOrm()
	schoolId := utils.Atoi64(input["school_id"][0])
	school := &models.School{Id: schoolId}
	o.Read(school)
	class := &models.Class{}
	class.School = school
	class.Name = input["name"][0]
	return o.Insert(class)
}

// 修改班级信息
func EditClass(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	p["school_id"] = input["school_id"][0]
	p["name"] = input["name"][0]
	return o.QueryTable("Class").Filter("id", input["id"][0]).Update(p)
}

// 获取班级列表
func GetClassList(p map[string]interface{}) (int64, []*models.Class) {
	var classList []*models.Class
	o := orm.NewOrm()
	qs := o.QueryTable("Class")
	relation, _ := p["relation"].(bool)
	if relation {
		qs = qs.RelatedSel()
	}
	schoolId, _ := p["schoolId"].(int64)
	if schoolId > 0 {
		qs = qs.Filter("school_id", schoolId)
	}
	totalCount, _ := qs.Count()
	curPage, ok := p["curPage"]
	perCount, _ := p["perCount"]
	if ok {
		start := perCount.(int) * (curPage.(int) - 1)
		qs = qs.Limit(perCount, start)
	}
	qs.All(&classList)
	return totalCount, classList
}