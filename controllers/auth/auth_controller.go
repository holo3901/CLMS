package auth

import (
	"CLMS/models/auth"
	"CLMS/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"strings"
	"time"
)

type AuthController struct {
	beego.Controller
}

func (a *AuthController) List() {

	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")

	auths := []auth.Auth{}

	// 每页显示的条数
	pagePerNum := 8

	// 当前页
	currentPage, err := a.GetInt("page")
	offsetNum := pagePerNum * (currentPage - 1)
	kw := a.GetString("kw")

	var count int64 = 0
	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)

	logs.Info(ret)
	if kw != "" { // 有查询条件的
		// 总数
		count, _ = qs.Filter("is_delete", 0).Filter("auth_name__contains", kw).Count()
		qs.Filter("is_delete", 0).Filter("auth_name__contains", kw).Limit(pagePerNum).Offset(offsetNum).All(&auths)
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).All(&auths)

	}
	if err != nil { // 说明没有获取到当前页
		currentPage = 1
	}
	// 总页数
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))

	prePage := 1
	if currentPage == 1 {
		prePage = currentPage
	} else if currentPage > 1 {
		prePage = currentPage - 1
	}

	// 当前页码小于总页数，下一页可以+1
	nextPage := 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else if currentPage >= countPage { // 不能+1
		nextPage = currentPage
	}

	// 当前页面=1，不能-1
	page_map := utils.Paginator(currentPage, pagePerNum, count)

	a.Data["page_map"] = page_map
	a.Data["countPage"] = countPage
	a.Data["count"] = count
	a.Data["auths"] = auths
	a.Data["currentPage"] = currentPage
	a.Data["prePage"] = prePage
	a.Data["nextPage"] = nextPage
	a.TplName = "auth/auth-list.html"

}

func (a *AuthController) ToAdd() {

	auths := []auth.Auth{}

	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	qs.Filter("is_delete", 0).All(&auths)
	a.Data["auths"] = auths
	a.TplName = "auth/auth-add.html"

}

func (a *AuthController) DoAdd() {
	auth_parent_id, _ := a.GetInt("auth_parent_id")
	auth_name := a.GetString("auth_name")
	auth_url := a.GetString("auth_url")
	auth_desc := a.GetString("auth_desc")
	is_active, _ := a.GetInt("is_active")
	auth_weight, _ := a.GetInt("auth_weight")

	o := orm.NewOrm()

	auth_data := auth.Auth{AuthName: auth_name, UrlFor: auth_url, Pid: auth_parent_id, Desc: auth_desc, CreateTime: time.Now(), IsActive: is_active, Weight: auth_weight}
	o.Insert(&auth_data)

	a.Data["json"] = map[string]interface{}{"code": 200, "msg": "添加成功"}
	a.ServeJSON()

}

func (u *AuthController) IsActive() {
	is_active, _ := u.GetInt("is_active_val")
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth").Filter("id", id)

	message_map := map[string]interface{}{}
	if is_active == 1 {
		qs.Update(orm.Params{
			"is_active": 0,
		})
		ret := fmt.Sprintf("权限id:%d,停用成功", id)
		rs := o.QueryTable("sys_auth").Filter("pid", id)
		rs.Update(orm.Params{
			"is_active": 0,
		})
		logs.Info(ret)
		message_map["msg"] = "停用成功"
	} else if is_active == 0 {
		qs.Update(orm.Params{
			"is_active": 1,
		})
		ret := fmt.Sprintf("权限id:%d,启用成功", id)
		logs.Info(ret)
		message_map["msg"] = "启用成功"
	}

	u.Data["json"] = message_map
	u.ServeJSON()
}

func (u *AuthController) MuliDelete() {

	ids := u.GetString("ids")
	//"3,7,8"
	new_ids := ids[1 : len(ids)-1]
	id_arr := strings.Split(new_ids, ",")

	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	for _, v := range id_arr {
		id_int := utils.StrToInt(v)
		qs.Filter("id", id_int).Update(orm.Params{
			"is_delete": 1,
		})

	}

	ret := fmt.Sprintf("批量删除成功，用户ids:%d", ids)
	logs.Info(ret)

	u.Data["json"] = map[string]interface{}{"code": 200, "msg": "批量删除成功"}
	u.ServeJSON()

}

func (u *AuthController) ToUpdate() {
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	user_data := auth.Auth{}
	o.QueryTable("sys_auth").Filter("id", id).One(&user_data)
	u.Data["auth"] = user_data
	ret := fmt.Sprintf("权限信息修改，权限id:%d", id)
	logs.Info(ret)
	u.TplName = "auth/auth_edit.html"
}

func (u *AuthController) DoUpdate() {
	uid, _ := u.GetInt("uid")
	authname := u.GetString("authname")
	desc := u.GetString("desc")
	is_active, _ := u.GetInt("is_active")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth").Filter("id", uid)

	message_map := map[string]interface{}{}

	_, err := qs.Update(orm.Params{
		"authname":  authname,
		"desc":      desc,
		"is_active": is_active,
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
