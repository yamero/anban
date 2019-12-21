package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentyun/cos-go-sdk-v5"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ResJsonStruct struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// 响应json数据
func ResJson(status int, msg string, data interface{}) *ResJsonStruct {
	return &ResJsonStruct{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
}

// md5字符串比较
func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}

// md5加密
func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// 字符串转int64
func Atoi64(a string) int64 {
	i, _ := strconv.ParseInt(a, 10, 64)
	return i
}

// 分页结构体
type Paginator struct {
	CurrentPage        int // 当前页码
	TotalCount         int // 总记录条数
	TotalPageCount     int // 总页数
	SymmetricPageCount int // 左右对称页码数
	PerCount           int // 每页显示记录条数
	Total              int // 总页码数
}

func NewPaginator(totalCount int, perCount int, symPageCount int, curPage int) *Paginator {
	totalPageCount := int(math.Ceil(float64(totalCount) / float64(perCount)))
	total := 2 * symPageCount + 1
	return &Paginator{
		CurrentPage:        curPage,
		TotalCount:         totalCount,
		TotalPageCount:     totalPageCount,
		SymmetricPageCount: symPageCount,
		PerCount:           perCount,
		Total:              total,
	}
}

// 获取分页页码
func (p *Paginator) GetPageHtml() string {
	if p.TotalPageCount <= 0 {
		return ""
	}
	h := fmt.Sprintf("<span style=\"margin-right:10px;border:none;\">共 %d 页</span>", p.TotalPageCount)
	if p.TotalPageCount == 1 {
		h += fmt.Sprintf("<span style=\"margin-left:10px;border:none;\">共 %d 条数据</span>", p.TotalCount)
		return h
	}
	prePage := p.CurrentPage - 1
	if prePage <= 1 {
		prePage = 1
	}
	nextPage := p.CurrentPage + 1
	if nextPage >= p.TotalPageCount {
		nextPage = p.TotalPageCount
	}
	start := p.CurrentPage - p.SymmetricPageCount
	end := p.CurrentPage + p.SymmetricPageCount
	if p.CurrentPage <= p.SymmetricPageCount {
		start = 1
		end = p.Total
	}
	if p.CurrentPage >= p.TotalPageCount-p.SymmetricPageCount {
		start = p.TotalPageCount - p.Total + 1
		end = p.TotalPageCount
	}
	if p.TotalPageCount <= p.Total {
		start = 1
		end = p.TotalPageCount
	}
	h += fmt.Sprintf("<a class=\"prev\" href=\"javascript:void(0)\" page-num=\"%d\">&lt;&lt;</a>", prePage)
	for i := start; i <= end; i++ {
		if i == p.CurrentPage {
			h += fmt.Sprintf("<span class=\"current\">%d</span>", i)
		} else {
			h += fmt.Sprintf("<a class=\"num\" href=\"javascript:void(0)\" page-num=\"%d\">%d</a>", i, i)
		}
	}
	h += fmt.Sprintf("<a class=\"next\" href=\"javascript:void(0)\" page-num=\"%d\">&gt;&gt;</a>", nextPage)
	h += fmt.Sprintf("<span style=\"margin-left:10px;border:none;\">共 %d 条数据</span>", p.TotalCount)
	return h
}

// 获取腾讯云存储对象API client
func GetStorageObj() *cos.Client {
	bucketDomain := beego.AppConfig.String("bucketdomain")
	secretId := beego.AppConfig.String("secretid")
	secretKey := beego.AppConfig.String("secretkey")
	u, _ := url.Parse(bucketDomain)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: secretId,
			SecretKey: secretKey,
		},
	})
	return c
}