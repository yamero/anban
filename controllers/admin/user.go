package admin

import (
	"anban/models"
	"anban/service"
	"anban/utils"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"html/template"
	"os"
	"strconv"
	"time"
)

type UserController struct {
	BaseController
}

func (c *UserController) Edit() {
	id := utils.Atoi64(c.Ctx.Input.Param(":id"))
	p := map[string]interface{}{}
	p["relation"] = false
	record := service.GetUserInfo(id, p)
	c.Data["record"] = record
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/user/edit.html"
}

func (c *UserController) DoEdit() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if len(input["real_name"][0]) <= 0 {
		res = utils.ResJson(0, "请输入姓名", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if len(input["id_card"][0]) <= 0 {
		res = utils.ResJson(0, "请输入身份证号", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if len(input["phone"][0]) <= 0 {
		res = utils.ResJson(0, "请输入电话", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.EditUser(input)
	if err != nil {
		res = utils.ResJson(0, "修改失败", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	res = utils.ResJson(1, "修改成功", "")
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *UserController) ShowList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	var realName string
	var idCard string
	var phone string
	var identity int
	var exportExcel int
	if c.GetString("identity") == "" {
		identity = -1
	} else {
		identity, _ = strconv.Atoi(c.GetString("identity"))
	}
	c.Ctx.Input.Bind(&realName, "real_name")
	c.Ctx.Input.Bind(&idCard, "id_card")
	c.Ctx.Input.Bind(&phone, "phone")
	c.Ctx.Input.Bind(&exportExcel, "export_excel")
	perCount, _ := beego.AppConfig.Int("percount")
	symPageCount, _ := beego.AppConfig.Int("symmetricpagecount")
	p := map[string]interface{}{}
	p["relation"] = false
	p["realName"] = realName
	p["idCard"] = idCard
	p["phone"] = phone
	p["identity"] = identity
	if exportExcel == 0 {
		p["curPage"] = curPage
		p["perCount"] = perCount
	}
	totalCount, recordList := service.GetUserList(p)
	if exportExcel == 1 && totalCount > 0 {
		c.ExportExcel(recordList)
		return
	}
	paginator := utils.NewPaginator(int(totalCount), perCount, symPageCount, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["recordList"] = recordList
	c.Data["realName"] = realName
	c.Data["idCard"] = idCard
	c.Data["phone"] = phone
	c.Data["identity"] = identity
	c.Data["identityList"] = models.UserIdentity
 	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/user/list.html"
}

func (c *UserController) ExportExcel(recordList []*models.User) {
	f := excelize.NewFile()
	sheetName := "sheet1"
	index := f.NewSheet(sheetName)
	col := 65
	header := []string{"姓名", "身份", "电话", "身份证号", "微信昵称", "关注状态", "状态", "注册时间"}
	for _, h := range header {
		f.SetCellValue(sheetName, string(col) + "1", h)
		col++
	}
	row := 2
	style, _ := f.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	f.SetColWidth(sheetName, "A", "H", 25)
	f.SetCellStyle(sheetName, "A1", fmt.Sprintf("H%d", len(recordList) + 1), style)
	for _, record := range recordList {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), record.RealName)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), record.IdentityShow)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), record.Phone)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), record.IdCard)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), record.NickName)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), record.SubscribeStatusShow)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), record.StatusShow)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), record.CreatedShow)
		row++
	}
	f.SetActiveSheet(index)
	d := "./static/export/" + time.Now().Format("20060102") + "/"
	if _, err := os.Stat(d); err != nil {
		os.MkdirAll(d, os.ModePerm)
	}
	saveName := d + "会员_" + time.Now().Format("20060102150405") + ".xlsx"
	err := f.SaveAs(saveName)
	if err != nil {
		c.Ctx.WriteString("导出失败")
		return
	}
	c.Ctx.Output.Download(saveName)
}