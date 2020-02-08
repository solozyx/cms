package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"

	"github.com/solozyx/cms/models"
	"github.com/solozyx/cms/utils"
)

var (
	sessionName = "cmsuser"
)

type LoginController struct {
	// NOTICE 不是继承 BaseController
	beego.Controller
}

func (c *LoginController) Index() {
	if c.Ctx.Request.Method == "POST" {
		userkey := strings.TrimSpace(c.GetString("userkey"))
		password := strings.TrimSpace(c.GetString("password"))
		if len(userkey) > 0 && len(password) > 0 {
			passmd5 := utils.Md5([]byte(password))
			user := models.GetUserByName(userkey)
			if passmd5 == user.Password {
				user.Password = ""
				c.SetSession(sessionName, user)
				fmt.Println("login ok")
				c.Redirect("/", 302)
				c.StopRun()
			}
		}
	}
	c.TplName = "login/index.html"
}
