package models

import "time"

var (
	SwipeCardDoor map[string]string
	SwipeCardStatus map[string]string
)

// 刷卡记录
type SwipeCardRecord struct {
	Id          int64     `orm:"description(自增主键)"`
	Student     *Student  `orm:"rel(fk);null;default(0);description(对应学生)"`
	Card        string    `orm:"description(卡号)"`
	Door        string    `orm:"description(门号)"`
	Status      string    `orm:"description(状态)"`
	Created     time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	CreatedShow string    `orm:"-"`
}

func init() {

	SwipeCardDoor = map[string]string{
		"1": "进门",
		"2": "进门",
		"3": "出门",
		"4": "出门",
	}

	SwipeCardStatus = map[string]string{
		"1": "合法开门",
		"2": "密码开门",
		"3": "卡加密码开门",
		"4": "手动输入卡加密码开门",
		"5": "首卡开门",
		"6": "门常开",
		"7": "多卡开门",
		"8": "重复读卡",
		"9": "有效期过期",
		"10": "不在开门时间段",
		"11": "节假日无效",
		"12": "非法卡",
		"13": "巡更卡不开门",
		"14": "探测锁定",
		"15": "无有效次数",
		"16": "防潜回",
		"17": "密码错误",
		"18": "密码加卡模式密码错误",
		"19": "锁定时或开门",
		"20": "锁定时",
		"21": "首卡未开门",
		"22": "挂失卡",
		"23": "黑名单卡",
		"24": "门内上限已满，禁止入门",
		"25": "开启防盗布防状态",
		"26": "撤销防盗布防状态",
		"27": "开启防盗布防状态",
		"28": "撤销防盗布防状态",
		"29": "互锁时或开门",
		"30": "互锁时",
		"31": "全卡开门",
	}

}
