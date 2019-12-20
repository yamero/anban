package wechat

import "encoding/xml"

// 微信发送过来的事件消息
type EventMsg struct {
	XMLName      xml.Name `xml:"xml"`          // XML根节点名字
	ToUserName   string   `xml:"ToUserName"`   // 公众号原始ID
	FromUserName string   `xml:"FromUserName"` // 用户openid
	CreateTime   int      `xml:"CreateTime"`   // 微信发消息的时间
	MsgType      string   `xml:"MsgType"`      // 消息类型，事件消息为event
	Event        string   `xml:"Event"`        // 事件类型，如用户关注、取消关注事件

	EventKey string `xml:"EventKey"` // 用户扫描带参数二维码时携带的参数
	Ticket   string `xml:"Ticket"`   // 二维码的ticket，可用来换取二维码图片

	Content string `xml:"Content"` // 用户文字消息内容

	PicUrl  string `xml:"PicUrl"`  // 用户图片消息的图片链接
	MediaId int64  `xml:"MediaId"` // 图片消息媒体id，可以调用获取临时素材接口拉取数据

	MsgId int64 `xml:"MsgId"` // 用户消息ID

	Latitude  float64 `xml:"Latitude"`  // 上报地理位置事件，位置纬度
	Longitude float64 `xml:"Longitude"` // 上报地理位置事件，位置经度
	Precision float64 `xml:"Precision"` // 上报地理位置事件，位置精度
}
