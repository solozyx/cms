package models

import (
	"encoding/json"
	"sort"

	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
)

// 菜单模型
type MenuModel struct {
	Mid    int    `orm:"pk;auto"` //菜单id 主键 自增
	Parent int    //父级菜单id
	Name   string `orm:"size(45)"` //varchar(45)
	Seq    int    //菜单排序
	Format string `orm:"size(2048);default({})"` //存储JSON数据
}

// 树形菜单
type MenuTree struct {
	MenuModel             // 本级菜单
	Child     []MenuModel // 子级菜单
}

// 设置表名
func (m *MenuModel) TableName() string {
	return TbNameMenu()
}

func MenuTreeStruct(user UserModel) map[int]MenuTree {
	query := orm.NewOrm().QueryTable(TbNameMenu())
	data := make([]*MenuModel, 0)
	// 树形菜单 parent父菜单id 排序字段倒序-seq
	query.OrderBy("parent", "-seq").Limit(1000).All(&data)
	// 菜单列表 menuId -> menu
	var menu = make(map[int]MenuTree)
	//auth
	if len(user.AuthStr) > 0 && len(data) > 0 {
		var authArr []int
		json.Unmarshal([]byte(user.AuthStr), &authArr)
		sort.Ints(authArr)

		for _, v := range data { //查询出来的数组
			if 0 == v.Parent { // 顶级父菜单 没有父菜单
				idx := sort.SearchInts(authArr, v.Mid)
				found := (idx < len(authArr) && authArr[idx] == v.Mid)
				if found {
					var tree = new(MenuTree)
					tree.MenuModel = *v
					menu[v.Mid] = *tree // menuId
				}
			} else {
				tmp, ok := menu[v.Parent]
				if ok { // 如果有父菜单
					tmp.Child = append(tmp.Child, *v)
					menu[v.Parent] = tmp
				}
			}
		}
	}

	return menu
}

func MenuList() ([]*MenuModel, int64) {
	query := orm.NewOrm().QueryTable(TbNameMenu())
	total, _ := query.Count()
	data := make([]*MenuModel, 0)
	// 父菜单升序 排序字段降序-seq
	query.OrderBy("parent", "-seq").Limit(1000).All(&data)
	return data, total
}

// 父菜单列表
func ParentMenuList() []*MenuModel {
	// parent = 0
	query := orm.NewOrm().QueryTable(TbNameMenu()).Filter("parent", 0)
	data := make([]*MenuModel, 0)
	query.OrderBy("-seq").Limit(1000).All(&data)
	return data
}

func MenuFormatStruct(mid int) *simplejson.Json {
	o := orm.NewOrm()
	menu := MenuModel{Mid: mid}
	err := o.Read(&menu)
	//fmt.Println(menu.Format)
	if nil == err {
		js, err2 := simplejson.NewJson([]byte(menu.Format))
		if nil == err2 {
			return js
		}
	}

	return nil
}

//func json2map(jsonstr string) {
//	//	aws, _ := js.Get("schema").Get("friends").Get("type").String()
//	//	fmt.Println(aws)
//}
