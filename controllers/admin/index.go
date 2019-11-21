package admin

// 后台主页
// @author lixin
type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	c.TplName = "admin/index.html"
}

func (c *IndexController) Welcome() {
	c.TplName = "admin/welcome.html"
}