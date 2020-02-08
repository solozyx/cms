package controllers

import (
	"strings"

	"github.com/astaxie/beego"

	"github.com/solozyx/cms/consts"
	"github.com/solozyx/cms/models"
)

// 所有controller的基类
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

// beego.Controller 的 Prepare 方法
// 所有controller方法执行前都会执行Prepare
func (c *BaseController) Prepare() {
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	beego.Informational(c.controllerName, c.actionName)
	// 用户权限验证 执行控制器的action 都会先执行Prepare
	user := c.auth()
	// 根据用户权限 展示菜单列表
	c.Data["Menu"] = models.MenuTreeStruct(user)
}

// 设置模板
// 第一个参数模板 第二个参数为layout
func (c *BaseController) setTpl(template ...string) {
	var tplName string
	// 默认模板路径 /views/common/layout.html
	layout := "common/layout.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		// 模板
		tplName = template[0]
		// 布局
		layout = template[1]
	default:
		// 不要"Controller"这个10个字母
		ctrlName := strings.ToLower(c.controllerName[0 : len(c.controllerName)-10])
		actionName := strings.ToLower(c.actionName)
		// 如 menu/index.html
		tplName = ctrlName + "/" + actionName + ".html"
	}

	_, found := c.Data["Footer"]
	if !found {
		c.Data["Footer"] = "menu/footerjs.html"
	}
	c.Layout = layout
	c.TplName = tplName
}

func (c *BaseController) jsonResult(code consts.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) listJsonResult(code consts.JsonResultCode, msg string, count int64, obj interface{}) {
	r := &models.ListJsonResult{code, msg, count, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) auth() models.UserModel {
	user := c.GetSession(sessionName)
	beego.Debug("base auth" + c.controllerName)
	beego.Debug(user)
	if user == nil {
		c.Redirect("/login", 302)
		c.StopRun()
		return models.UserModel{}
	} else {
		return user.(models.UserModel)
	}

}
