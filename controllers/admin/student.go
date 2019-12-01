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

type StudentController struct {
	BaseController
}

func (c *StudentController) Add() {
	p := map[string]interface{}{}
	p["relation"] = false
	_, schoolList := service.GetSchoolList(p)
	c.Data["schoolList"] = schoolList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/student/add.html"
}

func (c *StudentController) DoAdd() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["school_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择学校", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["class_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择班级", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["sn"][0] == "" {
		res = utils.ResJson(0, "请输入唯一标识", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["real_name"][0] == "" {
		res = utils.ResJson(0, "请输入姓名", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["id_card"][0] == "" {
		res = utils.ResJson(0, "请输入身份证号", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.AddStudent(input)
	if err != nil {
		res = utils.ResJson(0, "添加失败", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	res = utils.ResJson(1, "添加成功", "")
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *StudentController) Edit() {
	id := utils.Atoi64(c.Ctx.Input.Param(":id"))
	p := map[string]interface{}{}
	p["relation"] = true
	record := service.GetStudentInfo(id, p)
	c.Data["record"] = record
	p["relation"] = false
	_, schoolList := service.GetSchoolList(p)
	c.Data["schoolList"] = schoolList
	p["schoolId"] = record.School.Id
	_, classList := service.GetClassList(p)
	c.Data["classList"] = classList
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/student/edit.html"
}

func (c *StudentController) DoEdit() {
	var res *utils.ResJsonStruct
	input := c.Input()
	if input["school_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择学校", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["class_id"][0] <= "0" {
		res = utils.ResJson(0, "请选择班级", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["sn"][0] == "" {
		res = utils.ResJson(0, "请输入唯一标识", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["real_name"][0] == "" {
		res = utils.ResJson(0, "请输入姓名", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	if input["id_card"][0] == "" {
		res = utils.ResJson(0, "请输入身份证号", "")
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	_, err := service.EditStudent(input)
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

func (c *StudentController) ShowList() {
	curPage, _ := strconv.Atoi(c.GetString("p"))
	if curPage <= 0 {
		curPage = 1
	}
	var schoolId int64
	var classId int64
	var sn string
	var realName string
	var idCard string
	var exportExcel int
	c.Ctx.Input.Bind(&schoolId, "school_id")
	c.Ctx.Input.Bind(&classId, "class_id")
	c.Ctx.Input.Bind(&sn, "sn")
	c.Ctx.Input.Bind(&realName, "real_name")
	c.Ctx.Input.Bind(&idCard, "id_card")
	c.Ctx.Input.Bind(&exportExcel, "export_excel")
	perCount, _ := beego.AppConfig.Int("percount")
	symPageCount, _ := beego.AppConfig.Int("symmetricpagecount")
	p := map[string]interface{}{}
	p["relation"] = true
	p["schoolId"] = schoolId
	p["classId"] = classId
	p["sn"] = sn
	p["realName"] = realName
	p["idCard"] = idCard
	if exportExcel == 0 {
		p["curPage"] = curPage
		p["perCount"] = perCount
	}
	totalCount, recordList := service.GetStudentList(p)
	if exportExcel == 1 && totalCount > 0 {
		c.ExportExcel(recordList)
		return
	}
	paginator := utils.NewPaginator(int(totalCount), perCount, symPageCount, curPage)
	c.Data["paginator"] = paginator.GetPageHtml()
	c.Data["recordList"] = recordList
	c.Data["schoolId"] = schoolId
	c.Data["classId"] = classId
	c.Data["sn"] = sn
	c.Data["realName"] = realName
	c.Data["idCard"] = idCard
	p = map[string]interface{}{}
	p["relation"] = false
	_, schoolList := service.GetSchoolList(p)
	c.Data["schoolList"] = schoolList
	if schoolId > 0 {
		p["schoolId"] = schoolId
		_, ClassList := service.GetClassList(p)
		c.Data["classList"] = ClassList
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/student/list.html"
}

func (c *StudentController) ExportExcel(recordList []*models.Student) {
	f := excelize.NewFile()
	sheetName := "sheet1"
	index := f.NewSheet(sheetName)
	col := 65
	header := []string{"唯一标识", "姓名", "身份证号", "就读学校", "班级", "状态", "注册时间"}
	for _, h := range header {
		f.SetCellValue(sheetName, string(col) + "1", h)
		col++
	}
	row := 2
	style, _ := f.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	f.SetColWidth(sheetName, "A", "G", 25)
	f.SetCellStyle(sheetName, "A1", fmt.Sprintf("G%d", len(recordList) + 1), style)
	for _, record := range recordList {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), record.Sn)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), record.RealName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), record.IdCard)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), record.School.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), record.Class.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), record.StatusShow)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), record.CreatedShow)
		row++
	}
	f.SetActiveSheet(index)
	d := "./static/export/" + time.Now().Format("20060102") + "/"
	if _, err := os.Stat(d); err != nil {
		os.MkdirAll(d, os.ModePerm)
	}
	saveName := d + "学生_" + time.Now().Format("20060102150405") + ".xlsx"
	err := f.SaveAs(saveName)
	if err != nil {
		c.Ctx.WriteString("导出失败")
		return
	}
	c.Ctx.Output.Download(saveName)
}