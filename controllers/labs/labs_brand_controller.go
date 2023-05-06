package labs

import (
	"CLMS/models/auth"
	"CLMS/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"strings"
)

type LabsBrandController struct {
	beego.Controller
}

func (c *LabsBrandController) Get() {

	o := orm.NewOrm()

	qs := o.QueryTable("sys_labs_use")

	labs_u := []auth.LabsUse{}

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
		qs.Filter("is_delete", 0).Filter("name__contains", kw).Limit(pagePerNum).Offset(offsetNum).All(&labs_u)
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).All(&labs_u)

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
	c.Data["labs_u"] = labs_u
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw
	c.TplName = "labs/labs_brand_list.html"

}

func (c *LabsBrandController) ToAdd() {
	c.TplName = "labs/labs_brand_add.html"

}

func (c *LabsBrandController) DoAdd() {

	name := c.GetString("name")
	desc := c.GetString("desc")
	is_active, _ := c.GetInt("is_active")

	o := orm.NewOrm()
	Labs_U := auth.LabsUse{
		Name:     name,
		Desc:     desc,
		IsActive: is_active,
	}
	_, err := o.Insert(&Labs_U)

	message_map := map[string]interface{}{}

	if err != nil {
		message_map["code"] = 10001
		message_map["msg"] = "添加失败"
	}

	message_map["code"] = 200
	message_map["msg"] = "添加成功"

	c.Data["json"] = message_map
	c.ServeJSON()

}

func (u *LabsBrandController) MuliDelete() {

	ids := u.GetString("ids")
	//"3,7,8"
	new_ids := ids[1 : len(ids)-1]
	id_arr := strings.Split(new_ids, ",")

	o := orm.NewOrm()
	qs := o.QueryTable("sys_labs_use")
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

func (u *LabsBrandController) Delete() {
	id, _ := u.GetInt("id")

	o := orm.NewOrm()
	o.QueryTable("sys_labs_use").Filter("id", id).Update(orm.Params{
		"is_delete": 1,
	})
	ret := fmt.Sprintf("用户id:%d,删除成功", id)
	logs.Info(ret)
	u.Redirect(beego.URLFor("LabsBrandController.Get"), 302)

}

func (u *LabsBrandController) ToUpdate() {
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	user_data := auth.LabsUse{}
	o.QueryTable("sys_labs_use").Filter("id", id).One(&user_data)
	u.Data["labsuse"] = user_data
	ret := fmt.Sprintf("用户信息修改，用户id:%d", id)
	logs.Info(ret)
	u.TplName = "labs/labs_brand_edit.html"
}

func (u *LabsBrandController) DoUpdate() {
	uid, _ := u.GetInt("uid")
	labsusename := u.GetString("labsusename")
	desc := u.GetString("desc")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_labs_use").Filter("id", uid)

	message_map := map[string]interface{}{}

	_, err := qs.Update(orm.Params{
		"name": labsusename,
		"desc": desc,
	})
	if err != nil {
		ret := fmt.Sprintf("更新失败，角色id:%d", uid)
		logs.Error(ret)
		message_map["code"] = 10001
		message_map["msg"] = "更新失败"
	} else {
		ret := fmt.Sprintf("更新成功，角色id:%d", uid)
		logs.Info(ret)
		message_map["code"] = 200
		message_map["msg"] = "更新成功"
	}

	u.Data["json"] = message_map
	u.ServeJSON()

}

func (u *LabsBrandController) IsActive() {
	is_active, _ := u.GetInt("is_active_val")
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_labs_use").Filter("id", id)

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
