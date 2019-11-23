package service

import (
	"anban/models"
	"anban/utils"
	"github.com/astaxie/beego/orm"
	"net/url"
)

// 获取学校信息
func GetSchoolInfo(id int64) *models.School {
	o := orm.NewOrm()
	school := &models.School{}
	o.QueryTable("School").Filter("id", id).RelatedSel().One(school)
	return school
}

// 修改学校信息
func EditSchool(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	p["province_id"] = input["province_id"][0]
	p["city_id"] = input["city_id"][0]
	p["district_id"] = input["district_id"][0]
	p["name"] = input["name"][0]
	return o.QueryTable("School").Filter("id", input["id"][0]).Update(p)
}

// 添加学校
func AddSchool(input url.Values) (int64, error) {
	o := orm.NewOrm()
	provinceId := utils.Atoi64(input["province_id"][0])
	province := &models.Region{Id: provinceId}
	o.Read(province)
	cityId := utils.Atoi64(input["city_id"][0])
	city := &models.Region{Id: cityId}
	o.Read(city)
	districtId := utils.Atoi64(input["district_id"][0])
	district := &models.Region{Id: districtId}
	o.Read(district)
	school := &models.School{}
	school.Name = input["name"][0]
	school.Province = province
	school.City = city
	school.District = district
	return o.Insert(school)
}

// 获取学校列表
func GetSchoolList(p map[string]interface{}) (int64, []*models.School) {
	var schoolList []*models.School
	o := orm.NewOrm()
	qs := o.QueryTable("School")
	relation, _ := p["relation"].(bool)
	provinceId, _ := p["provinceId"].(int64)
	if provinceId > 0 {
		qs = qs.Filter("province_id", provinceId)
	}
	cityId, _ := p["cityId"].(int64)
	if cityId > 0 {
		qs = qs.Filter("city_id", cityId)
	}
	districtId, _ := p["districtId"].(int64)
	if districtId > 0 {
		qs = qs.Filter("district_id", districtId)
	}
	name, _ := p["name"].(string)
	if len(name) > 0 {
		qs = qs.Filter("name", name)
	}
	if relation {
		qs = qs.RelatedSel()
	}
	totalCount, _ := qs.Count()
	curPage, ok := p["curPage"]
	perCount, _ := p["perCount"]
	if ok {
		start := perCount.(int) * (curPage.(int) - 1)
		qs = qs.Limit(perCount, start)
	}
	qs.All(&schoolList)
	return totalCount, schoolList
}
