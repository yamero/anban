package service

import (
	"anban/models"
	"anban/utils"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strconv"
)

// 获取区域信息
func GetRegionInfo(id int64) *models.Region {
	o := orm.NewOrm()
	region := &models.Region{}
	o.QueryTable("Region").Filter("id", id).RelatedSel("Parent").One(region)
	return region
}

// 修改区域信息
func EditRegion(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	p["name"] = input["name"][0]
	p["code"] = input["code"][0]
	p["parent_id"] = input["parent_id"][0]
	return o.QueryTable("Region").Filter("id", input["id"][0]).Update(p)
}

// 添加区域信息
func AddRegion(input url.Values) (int64, error) {
	o := orm.NewOrm()
	level, _ := strconv.Atoi(input["level"][0])
	parentId := utils.Atoi64(input["parent_id"][0])
	parentRegion := &models.Region{}
	if level > 1 {
		parentRegion.Id = parentId
		o.Read(parentRegion)
	}
	region := models.Region{}
	region.Name = input["name"][0]
	region.Code = input["code"][0]
	region.Level = level
	region.Parent = parentRegion
	return o.Insert(&region)
}

// 获取区域列表
func GetRegionList(p map[string]interface{}) (int64, []*models.Region) {
	var provinceList []*models.Region
	o := orm.NewOrm()
	qs := o.QueryTable("Region").Filter("level", p["level"].(int))
	relation, _ := p["relation"].(bool)
	if relation {
		qs = qs.RelatedSel("Parent")
	}
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