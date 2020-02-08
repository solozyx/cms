package models

import (
	"github.com/astaxie/beego/orm"
)

// init 初始化 导入本包就会调用init函数
func init() {
	// 注册model 调用orm创建数据表
	orm.RegisterModel(new(MenuModel), new(UserModel), new(DataModel))
}

func TbNameMenu() string {
	return "cms_menu"
}

func TbNameUser() string {
	return "cms_user"
}

func TbNameData() string {
	return "cms_data"
}
