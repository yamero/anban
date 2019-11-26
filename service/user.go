package service

import (
	"anban/models"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strconv"
	"time"
)

// 获取会员信息
func GetUserInfo(id int64) *models.User {
	o := orm.NewOrm()
	user := &models.User{}
	o.QueryTable("User").Filter("id", id).One(user)
	user.IdentityShow = models.UserIdentity[user.Identity]
	user.SubscribeStatusShow = models.UserSubscribeStatus[user.SubscribeStatus]
	user.SubscribeTimeShow = user.SubscribeTime.Format("2006-01-02 15:04:05")
	user.UnsubscribeTimeShow = user.UnsubscribeTime.Format("2006-01-02 15:04:05")
	user.CreatedShow = user.Created.Format("2006-01-02 15:04:05")
	user.UpdatedShow = user.Updated.Format("2006-01-02 15:04:05")
	return user
}

// 添加会员信息
func AddUser(input url.Values) (int64, error) {
	o := orm.NewOrm()
	user := &models.User{}
	if _, ok := input["open_id"]; ok {
		user.OpenId = input["open_id"][0]
	}
	if _, ok := input["union_id"]; ok {
		user.UnionId = input["union_id"][0]
	}
	if _, ok := input["status"]; ok {
		user.Status, _ = strconv.Atoi(input["status"][0])
	}
	return o.Insert(user)
}

// 修改会员信息
func EditUser(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	if _, ok := input["real_name"]; ok {
		p["real_name"] = input["real_name"][0]
	}
	if _, ok := input["id_card"]; ok {
		p["id_card"] = input["id_card"][0]
	}
	if _, ok := input["phone"]; ok {
		p["phone"] = input["phone"][0]
	}
	if _, ok := input["short_name"]; ok {
		p["short_name"] = input["short_name"][0]
	}
	if _, ok := input["status"]; ok {
		p["status"] = input["status"][0]
	}
	p["updated"] = time.Now().Format("2006-01-02 15:04:05")
	return o.QueryTable("User").Filter("id", input["id"][0]).Update(p)
}

// 获取会员列表
func GetUserList(p map[string]interface{}) (int64, []*models.User) {
	var userList []*models.User
	o := orm.NewOrm()
	qs := o.QueryTable("User")
	relation, _ := p["relation"].(bool)
	if relation {
		qs = qs.RelatedSel()
	}
	realName, _ := p["realName"].(string)
	if len(realName) > 0 {
		qs = qs.Filter("real_name", realName)
	}
	idCard, _ := p["idCard"].(string)
	if len(idCard) > 0 {
		qs = qs.Filter("id_card", idCard)
	}
	phone, _ := p["phone"].(string)
	if len(phone) > 0 {
		qs = qs.Filter("phone", phone)
	}
	identity, _ := p["identity"].(int)
	if identity > 0 {
		qs = qs.Filter("identity", identity)
	}
	totalCount, _ := qs.Count()
	curPage, ok := p["curPage"]
	perCount, _ := p["perCount"]
	if ok {
		start := perCount.(int) * (curPage.(int) - 1)
		qs = qs.Limit(perCount, start)
	}
	qs.OrderBy("-created").All(&userList)
	for _, user := range userList {
		user.StatusShow = models.UserStatus[user.Status]
		user.IdentityShow = models.UserIdentity[user.Identity]
		user.SubscribeStatusShow = models.UserSubscribeStatus[user.SubscribeStatus]
		user.SubscribeTimeShow = user.SubscribeTime.Format("2006-01-02 15:04:05")
		user.UnsubscribeTimeShow = user.UnsubscribeTime.Format("2006-01-02 15:04:05")
		user.CreatedShow = user.Created.Format("2006-01-02 15:04:05")
		user.UpdatedShow = user.Updated.Format("2006-01-02 15:04:05")
	}
	return totalCount, userList
}