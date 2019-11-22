package service

import (
	"anban/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

// 添加省
func AddProvince(input url.Values) (int64, error) {
	o := orm.NewOrm()
	region := models.Region{}
	region.Name = input["name"][0]
	region.Code = input["code"][0]
	region.Level = 1
	region.Parent = &models.Region{}
	return o.Insert(&region)
}

// 获取省列表
func GetProvinceList(p map[string]interface{}) (int64, []*models.Region) {
	var provinceList []*models.Region
	o := orm.NewOrm()
	qs := o.QueryTable("Region").Filter("level", 1)
	totalCount, _ := qs.Count()
	curPage, ok := p["curPage"]
	perCount, _ := p["perCount"]
	if ok {
		start := perCount.(int) * (curPage.(int) - 1)
		qs = qs.Limit(perCount, start)
	}
	qs.All(&provinceList)
	for _, province := range provinceList {
		province.LevelShow = models.RegionLevel[province.Level]
	}
	return totalCount, provinceList
}