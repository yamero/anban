package models

// 后台管理员角色
type UserAdminRole struct {
	Id        int64  `orm:"description(自增主键)"`
	RightList string `orm:"null;default();type(text);description(权限Id 逗号分隔的字符串)"`
	RoleName  string `orm:"description(角色名称)"`
	Describe  string `orm:"description(角色描述)"`
}
