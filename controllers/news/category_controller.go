package news

import (
	"CLMS/models/news"
	"CLMS/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"strings"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {

	o := orm.NewOrm()

	qs := o.QueryTable("sys_category")

	categrories := []news.Category{}
	// 每页显示的条数
	pagePerNum := 8
	// 当前页
	currentPage, err := c.GetInt("page")
	if err != nil { // 说明没有获取到当前页
		currentPage = 1
	}

	offsetNum := pagePerNum * (currentPage - 1)

	kw := c.GetString("kw")

	var count int64 = 0

	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	if kw != "" { // 有查询条件的
		// 总数
		count, _ = qs.Filter("is_delete", 0).Filter("name__contains", kw).Count()
		qs.Filter("is_delete", 0).Filter("name__contains", kw).Limit(pagePerNum).Offset(offsetNum).All(&categrories)
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).All(&categrories)

	}

	// 总页数
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))

	prePage := 1
	if currentPage == 1 {
		prePage = currentPage
	} else if currentPage > 1 {
		prePage = currentPage - 1
	}

	nextPage := 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else if currentPage >= countPage {
		nextPage = currentPage
	}

	page_map := utils.Paginator(currentPage, pagePerNum, count)

	c.Data["categrories"] = categrories
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw

	c.TplName = "news/category_list.html"

}

func (c *CategoryController) ToAdd() {
	c.TplName = "news/category_add.html"

}

func (c *CategoryController) DoAdd() {

	name := c.GetString("name")
	desc := c.GetString("desc")
	is_active, _ := c.GetInt("is_active")

	o := orm.NewOrm()
	category := news.Category{
		Name:     name,
		Desc:     desc,
		IsActive: is_active,
	}
	_, err := o.Insert(&category)

	message_map := map[string]interface{}{}
	if err != nil {
		message_map["code"] = 10001
		message_map["msg"] = "添加栏目失败"

	}

	message_map["code"] = 200
	message_map["msg"] = "添加成功"

	c.Data["json"] = message_map
	c.ServeJSON()

}

func (u *CategoryController) Delete() {
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	o.QueryTable("sys_category").Filter("id", id).Update(orm.Params{
		"is_delete": 1,
	})
	ret := fmt.Sprintf("用户id:%d,删除成功", id)
	logs.Info(ret)
	u.Redirect(beego.URLFor("CategoryController.Get"), 302)

}

func (u *CategoryController) IsActive() {
	is_active, _ := u.GetInt("is_active_val")
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_category").Filter("id", id)

	message_map := map[string]interface{}{}
	if is_active == 1 {
		qs.Update(orm.Params{
			"is_active": 0,
		})
		ret := fmt.Sprintf("用户id:%d,停用成功", id)
		logs.Info(ret)
		message_map["msg"] = "停用成功"
	} else if is_active == 0 {
		qs.Update(orm.Params{
			"is_active": 1,
		})
		ret := fmt.Sprintf("用户id:%d,启用成功", id)
		logs.Info(ret)
		message_map["msg"] = "启用成功"
	}

	u.Data["json"] = message_map
	u.ServeJSON()
}
func (u *CategoryController) ToUpdate() {
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	user_data := news.Category{}
	o.QueryTable("sys_category").Filter("id", id).One(&user_data)
	u.Data["category"] = user_data
	ret := fmt.Sprintf("栏目信息修改，栏目id:%d", id)
	logs.Info(ret)
	u.TplName = "news/category_edit.html"
}

func (u *CategoryController) DoUpdate() {
	uid, _ := u.GetInt("uid")
	name := u.GetString("name")
	desc := u.GetString("desc")
	is_active, _ := u.GetInt("is_active")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_category").Filter("id", uid)

	message_map := map[string]interface{}{}

	_, err := qs.Update(orm.Params{
		"name":      name,
		"desc":      desc,
		"is_active": is_active,
	})
	if err != nil {
		ret := fmt.Sprintf("更新失败，栏目id:%d", uid)
		logs.Error(ret)
		message_map["code"] = 10001
		message_map["msg"] = "更新失败"
	} else {
		ret := fmt.Sprintf("更新成功，栏目id:%d", uid)
		logs.Info(ret)
		message_map["code"] = 200
		message_map["msg"] = "更新成功"
	}

	u.Data["json"] = message_map
	u.ServeJSON()

}
func (u *CategoryController) MuliDelete() {

	ids := u.GetString("ids")
	//"3,7,8"
	new_ids := ids[1 : len(ids)-1]
	id_arr := strings.Split(new_ids, ",")

	o := orm.NewOrm()
	qs := o.QueryTable("sys_category")
	for _, v := range id_arr {
		id_int := utils.StrToInt(v)
		qs.Filter("id", id_int).Update(orm.Params{
			"is_delete": 1,
		})

	}

	ret := fmt.Sprintf("批量删除成功，栏目ids:%d", ids)
	logs.Info(ret)

	u.Data["json"] = map[string]interface{}{"code": 200, "msg": "批量删除成功"}
	u.ServeJSON()

}
