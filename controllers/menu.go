package controllers

import (
	"fmt"

	"github.com/astaxie/beego/orm"

	"github.com/solozyx/cms/consts"
	"github.com/solozyx/cms/models"
)

type MenuController struct {
	// 继承base
	BaseController
}

func (c *MenuController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs.html"
	// c.setTpl("menu/index.html") 可省略 "menu/index.html" 走 base的default分支
	c.setTpl()
}

func (c *MenuController) List() {
	data, total := models.MenuList()
	// 父菜单名字 footer.js {field:'ParentName', title: '父菜单'}
	type MenuEx struct {
		models.MenuModel
		ParentName string
	}
	// menuId -> menuName
	var menu = make(map[int]string)
	menu[0] = "-"
	for _, v := range data {
		menu[v.Mid] = v.Name
	}

	var dataEx []MenuEx
	for _, v := range data {
		dataEx = append(dataEx, MenuEx{*v, menu[v.Parent]})
	}

	c.listJsonResult(consts.JRCodeSucc, "ok", total, dataEx)
}

func (c *MenuController) Add() {
	//选择父菜单数据
	data, _ := models.MenuList()
	var parentMenus []models.MenuModel
	for _, value := range data {
		if 0 == value.Parent {
			parentMenus = append(parentMenus, *value)
		}
	}
	c.Data["PMenus"] = parentMenus
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs_add.html"
	c.setTpl("menu/add.html", "common/layout_edit.html")
}
func (c *MenuController) AddDo() {
	var m models.MenuModel
	if err := c.ParseForm(&m); err == nil {
		id, _ := orm.NewOrm().Insert(&m)
		c.jsonResult(consts.JRCodeSucc, "ok", id)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}

func (c *MenuController) Edit() {
	c.Data["Mid"] = c.GetString("mid")
	c.Data["Parent"], _ = c.GetInt("parent")
	c.Data["Seq"] = c.GetString("seq")
	c.Data["Name"] = c.GetString("name")

	// 父菜单数据
	var parentMenus []models.MenuModel
	data, _ := models.MenuList()
	for _, value := range data {
		if 0 == value.Parent {
			parentMenus = append(parentMenus, *value)
		}
	}
	c.Data["PMenus"] = parentMenus

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs_edit.html"
	c.setTpl("menu/edit.html", "common/layout_edit.html")
}

func (c *MenuController) EditDo() {
	var m models.MenuModel
	if err := c.ParseForm(&m); err == nil {
		id, _ := orm.NewOrm().Update(&m)
		c.jsonResult(consts.JRCodeSucc, "ok", id)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}

func (c *MenuController) DeleteDo() {
	if mid, err := c.GetInt("mid"); err == nil {
		num, _ := orm.NewOrm().Delete(&models.MenuModel{Mid: mid})
		c.jsonResult(consts.JRCodeSucc, "ok", num)
	} else {
		fmt.Println(err, mid)
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}
