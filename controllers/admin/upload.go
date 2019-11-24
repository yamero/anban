package admin

import (
	"anban/utils"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

type UploadController struct {
	BaseController
}

func (c *UploadController) Prepare() {
	c.EnableXSRF = false
}

func (c *UploadController) Editor() {
	res := struct{
		Errno int `json:"errno"`
		Msg string `json:"msg"`
		Data []string `json:"data"`
	}{}
	fileList, err := c.GetFiles("editorFile")
	if err != nil {
		res.Errno = 1
		res.Msg = "获取文件失败" + err.Error()
		c.Data["json"] = &res
		c.ServeJSON()
		return
	}
	maxSize := int64(2 << 20)
	for _, h := range fileList {
		if h.Size > maxSize {
			res.Errno = 1
			res.Msg = "上传的每个文件大小不能超过2M"
			c.Data["json"] = &res
			c.ServeJSON()
			return
		}
	}
	for _, fh := range fileList {
		f, err := fh.Open()
		if err != nil {
			res.Errno = 1
			res.Msg = "有文件上传失败"
			c.Data["json"] = &res
			c.ServeJSON()
			return
		}
		d := "./static/upload/" + time.Now().Format("20060102")
		_, err = os.Stat(d)
		if err != nil {
			os.MkdirAll(d, os.ModePerm)
		}
		ext := path.Ext(fh.Filename)
		n := fmt.Sprintf("%d", time.Now().UnixNano())
		toFile := d + "/" + utils.Encode(n) + ext
		dst, err := os.OpenFile(toFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			res.Errno = 1
			res.Msg = "有文件上传失败"
			c.Data["json"] = &res
			c.ServeJSON()
			return
		}
		io.Copy(dst, f)
		res.Data = append(res.Data, strings.TrimLeft(toFile, "."))
		f.Close()
		dst.Close()
	}
	res.Errno = 0
	c.Data["json"] = &res
	c.ServeJSON()
}

/*
// 单个文件上传
func (c *UploadController) Editor1() {
	res := struct{
		Errno int `json:"errno"`
		Msg string `json:"msg"`
		Data []string `json:"data"`
	}{}
	f, h, err := c.GetFile("editorFile")
	if err != nil {
		res.Errno = 1
		res.Msg = "上传文件出错" + err.Error()
		c.Data["json"] = &res
		c.ServeJSON()
		return
	}
	defer f.Close()
	maxSize := int64(2 << 20)
	if h.Size > maxSize {
		res.Errno = 1
		res.Msg = "上传文件大小不能超过2M"
		c.Data["json"] = &res
		c.ServeJSON()
		return
	}
	ext := path.Ext(h.Filename)
	enableExt := map[string]bool{
		".jpg": true,
		".JPG": true,
		".jpeg": true,
		".JPEG": true,
		".png": true,
		".PNG": true,
		".gif": true,
		".GIF": true,
	}
	_, ok := enableExt[ext]
	if !ok {
		res.Errno = 1
		res.Msg = "上传文件格式不正确"
		c.Data["json"] = &res
		c.ServeJSON()
		return
	}
	t := time.Now().Format("20060102")
	d := "./static/upload/" + t
	_, err = os.Stat(d)
	if err != nil {
		os.MkdirAll(d, os.ModePerm)
	}
	n := fmt.Sprintf("%d", time.Now().UnixNano())
	toFile := d + "/" + utils.Encode(n) + ext
	err = c.SaveToFile("editorFile", toFile)
	if err != nil {
		res.Errno = 1
		res.Msg = "上传文件出错" + err.Error()
		c.Data["json"] = &res
		c.ServeJSON()
		return
	}
	res.Errno = 0
	res.Data = append(res.Data, strings.TrimLeft(toFile, "."))
	c.Data["json"] = &res
	c.ServeJSON()
}*/
