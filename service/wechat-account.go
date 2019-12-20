package service

import (
	"anban/models"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strconv"
	"time"
)

// 获取公众号信息
func GetWechatAccountInfo(id int64, p map[string]interface{}) *models.WechatAccount {
	o := orm.NewOrm()
	wechatAccount := &models.WechatAccount{}
	qs := o.QueryTable("WechatAccount")
	relation, _ := p["relation"].(bool)
	if relation {
		qs = qs.RelatedSel()
	}
	if id > 0 {
		qs = qs.Filter("id", id)
	}
	qs.One(wechatAccount)
	wechatAccount.StatusShow = models.WechatAccountStatus[wechatAccount.Status]
	wechatAccount.CreatedShow = wechatAccount.Created.Format("2006-01-02 15:04:05")
	wechatAccount.UpdatedShow = wechatAccount.Updated.Format("2006-01-02 15:04:05")
	return wechatAccount
}

// 添加公众号信息
func AddWechatAccount(input url.Values) (int64, error) {
	o := orm.NewOrm()
	wechatAccount := &models.WechatAccount{}
	if _, ok := input["original_id"]; ok {
		wechatAccount.OriginalId = input["original_id"][0]
	}
	if _, ok := input["name"]; ok {
		wechatAccount.Name = input["name"][0]
	}
	if _, ok := input["app_id"]; ok {
		wechatAccount.AppId = input["app_id"][0]
	}
	if _, ok := input["app_secret"]; ok {
		wechatAccount.AppSecret = input["app_secret"][0]
	}
	if _, ok := input["token"]; ok {
		wechatAccount.Token = input["token"][0]
	}
	if _, ok := input["encoding_aes_key"]; ok {
		wechatAccount.EncodingAesKey = input["encoding_aes_key"][0]
	}
	if _, ok := input["notify_url"]; ok {
		wechatAccount.NotifyUrl = input["notify_url"][0]
	}
	if _, ok := input["status"]; ok {
		wechatAccount.Status, _ = strconv.Atoi(input["status"][0])
	}
	return o.Insert(wechatAccount)
}

// 修改公众号信息
func EditWechatAccount(input url.Values) (int64, error) {
	o := orm.NewOrm()
	p := orm.Params{}
	/*if _, ok := input["original_id"]; ok {
		p["original_id"] = input["original_id"][0]
	}
	if _, ok := input["name"]; ok {
		p["name"] = input["name"][0]
	}
	if _, ok := input["app_id"]; ok {
		p["app_id"] = input["app_id"][0]
	}
	if _, ok := input["app_secret"]; ok {
		p["app_secret"] = input["app_secret"][0]
	}
	if _, ok := input["token"]; ok {
		p["token"] = input["token"][0]
	}
	if _, ok := input["status"]; ok {
		p["status"] = input["status"][0]
	}*/
	for k, v := range input {
		p[k] = v[0]
	}
	p["updated"] = time.Now().Format("2006-01-02 15:04:05")
	return o.QueryTable("WechatAccount").Filter("id", input["id"][0]).Update(p)
}

// 获取公众号列表
func GetWechatAccountList(p map[string]interface{}) (int64, []*models.WechatAccount) {
	var wechatAccountList []*models.WechatAccount
	o := orm.NewOrm()
	qs := o.QueryTable("WechatAccount")
	relation, _ := p["relation"].(bool)
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
	qs.OrderBy("-created").All(&wechatAccountList)
	for _, wechatAccount := range wechatAccountList {
		wechatAccount.StatusShow = models.WechatAccountStatus[wechatAccount.Status]
		wechatAccount.CreatedShow = wechatAccount.Created.Format("2006-01-02 15:04:05")
		wechatAccount.UpdatedShow = wechatAccount.Updated.Format("2006-01-02 15:04:05")
	}
	return totalCount, wechatAccountList
}