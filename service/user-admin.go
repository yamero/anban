package service

import (
	"anban/models"
	"anban/utils"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strconv"
	"time"
)

// 获取管理员信息
func GetUserAdminInfo(id int64) *models.UserAdmin {
	o := orm.NewOrm()
	userAdmin := &models.UserAdmin{}
	o.QueryTable("UserAdmin").Filter("id", id).One(userAdmin)
	userAdmin.CreatedShow = userAdmin.Created.Format("2006-01-02 15:04:05")
	userAdmin.UpdatedShow = userAdmin.Updated.Format("2006-01-02 15:04:05")
	return userAdmin
}

// 获取管理员列表
func GetUserAdminList(p map[string]interface{}) (int64, []*models.UserAdmin) {
	var userAdminList []*models.UserAdmin
	o := orm.NewOrm()
	qs := o.QueryTable("UserAdmin")
	totalCount, _ := qs.Count()
	curPage, ok := p["curPage"]
	perCount, _ := p["perCount"]
	if ok {
		start := perCount.(int) * (curPage.(int) - 1)
		qs = qs.Limit(perCount, start)
	}
	qs.RelatedSel("UserAdminRole").All(&userAdminList)
	for _, userAdmin := range userAdminList {
		userAdmin.StatusShow = models.UserAdminStatus[userAdmin.Status]
		userAdmin.CreatedShow = userAdmin.Created.Format("2006-01-02 15:04:05")
		userAdmin.UpdatedShow = userAdmin.Updated.Format("2006-01-02 15:04:05")
	}
	return totalCount, userAdminList
}

// 添加管理员
func AddUserAdmin(input url.Values) (int64, error) {
	o := orm.NewOrm()
	roleId := utils.Atoi64(input["role_id"][0])
	userAdminRole := models.UserAdminRole{}
	if roleId > 0 {
		userAdminRole.Id = roleId
		o.Read(&userAdminRole)
	}
	userAdmin := models.UserAdmin{}
	userAdmin.Account = input["account"][0]
	userAdmin.Password = utils.Encode(input["password"][0])
	if _, ok := input["mobile"]; ok {
		userAdmin.Mobile = input["mobile"][0]
	}
	if _, ok := input["email"]; ok {
		userAdmin.Email = input["email"][0]
	}
	if _, ok := input["real_name"]; ok {
		userAdmin.RealName = input["real_name"][0]
	}
	if _, ok := input["status"]; ok {
		status, _ := strconv.Atoi(input["status"][0])
		userAdmin.Status = status
	}
	userAdmin.UserAdminRole = &userAdminRole
	return o.Insert(&userAdmin)
}

// 编辑管理员
func EditUserAdmin(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	p["user_admin_role_id"] = input["role_id"][0]
	p["account"] = input["account"][0]
	if len(input["password"][0]) >= 6 {
		p["password"] = utils.Encode(input["password"][0])
	}
	p["mobile"] = input["mobile"][0]
	p["email"] = input["email"][0]
	p["real_name"] = input["real_name"][0]
	p["status"] = input["status"][0]
	p["updated"] = time.Now().Format("2006-01-02 15:04:05")
	return o.QueryTable("UserAdmin").Filter("id", input["id"][0]).Update(p)
}
