package wechat

import (
	"anban/service"
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"sort"
	"strings"
)

// 与微信通信的签名算法
func GetWechatSignature(timestamp, nonce string) string {
	id, _ := beego.AppConfig.Int64("anbanid")
	p := map[string]interface{}{}
	record := service.GetWechatAccountInfo(id, p)
	//1. 将 token、timestamp、nonce三个参数进行字典序排序
	sl := []string{record.Token, timestamp, nonce}
	sort.Strings(sl)
	//2. 将三个参数字符串拼接成一个字符串进行sha1加密
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
