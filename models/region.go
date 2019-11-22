package models

var RegionLevel map[int]string

// 区域
type Region struct {
	Id        int64   `orm:"description(自增主键)"`
	Name      string  `orm:"description(区域名称)"`
	Parent    *Region `orm:"null;default(0);rel(fk);description(上级Id)"`
	Level     int     `orm:"description(区域级别 1省 2市 3区/县 4乡/镇 5村)"`
	LevelShow string  `orm:"-"`
	Code      string  `orm:"description(区域行政代码)"`
}

func init() {
	RegionLevel = map[int]string{
		1: "省",
		2: "市",
		3: "区/县",
		4: "乡/镇",
		5: "村",
	}
}
