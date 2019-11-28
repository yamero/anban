package service

import (
	"anban/models"
	"anban/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

// 获取学生信息
func GetStudentInfo(id int64, p map[string]interface{}) *models.Student {
	o := orm.NewOrm()
	student := &models.Student{}
	qs := o.QueryTable("Student")
	relation, _ := p["relation"].(bool)
	if relation {
		qs = qs.RelatedSel()
	}
	if id > 0 {
		qs = qs.Filter("id", id)
	}
	sn, _ := p["sn"].(string)
	if len(sn) > 0 {
		qs = qs.Filter("sn", sn)
	}
	qs.One(student)
	student.CreatedShow = student.Created.Format("2006-01-02 15:04:05")
	student.UpdatedShow = student.Updated.Format("2006-01-02 15:04:05")
	return student
}

// 添加学生信息
func AddStudent(input url.Values) (int64, error) {
	o := orm.NewOrm()
	school := &models.School{}
	if _, ok := input["school_id"]; ok {
		schoolId := utils.Atoi64(input["school_id"][0])
		school.Id = schoolId
		o.Read(school)
	}
	class := &models.Class{}
	if _, ok := input["class_id"]; ok {
		classId := utils.Atoi64(input["class_id"][0])
		class.Id = classId
		o.Read(class)
	}
	student := &models.Student{}
	student.School = school
	student.Class = class
	if _, ok := input["sn"]; ok {
		student.Sn = input["sn"][0]
	}
	if _, ok := input["real_name"]; ok {
		student.RealName = input["real_name"][0]
	}
	if _, ok := input["id_card"]; ok {
		student.IdCard = input["id_card"][0]
	}
	if _, ok := input["status"]; ok {
		student.Status, _ = strconv.Atoi(input["status"][0])
	}
	student.FamilySn = fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
	return o.Insert(student)
}

// 修改学生信息
func EditStudent(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	p["school_id"] = input["school_id"][0]
	p["class_id"] = input["class_id"][0]
	p["sn"] = input["sn"][0]
	p["real_name"] = input["real_name"][0]
	p["id_card"] = input["id_card"][0]
	p["status"] = input["status"][0]
	p["updated"] = time.Now().Format("2006-01-02 15:04:05")
	return o.QueryTable("Student").Filter("id", input["id"][0]).Update(p)
}

// 获取学生列表
func GetStudentList(p map[string]interface{}) (int64, []*models.Student) {
	var studentList []*models.Student
	o := orm.NewOrm()
	qs := o.QueryTable("Student")
	relation, _ := p["relation"].(bool)
	if relation {
		qs = qs.RelatedSel()
	}
	schoolId, _ := p["schoolId"].(int64)
	if schoolId > 0 {
		qs = qs.Filter("school_id", schoolId)
	}
	classId, _ := p["classId"].(int64)
	if classId > 0 {
		qs = qs.Filter("class_id", classId)
	}
	sn, _ := p["sn"].(string)
	if len(sn) > 0 {
		qs = qs.Filter("sn", sn)
	}
	realName, _ := p["realName"].(string)
	if len(realName) > 0 {
		qs = qs.Filter("real_name", realName)
	}
	idCard, _ := p["idCard"].(string)
	if len(idCard) > 0 {
		qs = qs.Filter("id_card", idCard)
	}
	totalCount, _ := qs.Count()
	curPage, ok := p["curPage"]
	perCount, _ := p["perCount"]
	if ok {
		start := perCount.(int) * (curPage.(int) - 1)
		qs = qs.Limit(perCount, start)
	}
	qs.OrderBy("-created").All(&studentList)
	for _, student := range studentList {
		student.StatusShow = models.StudentStatus[student.Status]
		student.CreatedShow = student.Created.Format("2006-01-02 15:04:05")
		student.UpdatedShow = student.Updated.Format("2006-01-02 15:04:05")
	}
	return totalCount, studentList
}