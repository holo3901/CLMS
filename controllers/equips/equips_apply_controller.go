package equips

import (
	"CLMS/models/auth"
	"CLMS/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
	"time"
)

type EquipsApplyController struct {
	beego.Controller
}

func (c *EquipsApplyController) Get() {
	o := orm.NewOrm()

	qs := o.QueryTable("sys_equips")

	equips_data := []auth.Equips{}

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
		qs.Filter("is_delete", 0).Filter("name__contains", kw).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&equips_data)
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&equips_data)

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
	c.Data["equips_data"] = equips_data
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw

	c.TplName = "equips/equips_apply_list.html"

}

func (c *EquipsApplyController) ToApply() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_labs")
	labs := []auth.Labs{}
	qs.Filter("is_delete", 0).All(&labs)
	id, _ := c.GetInt("id")
	c.Data["id"] = id
	c.Data["labs"] = labs
	c.TplName = "equips/equips_apply.html"

}

func (c *EquipsApplyController) DoApply() {
	mount, _ := c.GetInt("mount")
	reason := c.GetString("reason")
	labs_id := c.GetString("labs")
	return_date := c.GetString("return_date")
	return_date_new, _ := time.Parse("2006-01-02", return_date)
	equips_id, _ := c.GetInt("equips_id")
	uid := c.GetSession("id")

	// interface --> int
	user := auth.User{Id: uid.(int)}
	equips_date := auth.Equips{Id: equips_id}
	i, _ := strconv.Atoi(labs_id)
	Labs := auth.Labs{Id: i}
	o := orm.NewOrm()
	// 默认：ReturnStatus=0，AuditStatus=3，IsActive=1
	equips_apply := auth.EquipsApply{
		Mount:        mount,
		User:         &user,
		Equips:       &equips_date,
		Reason:       reason,
		Labs:         &Labs,
		ReturnDate:   return_date_new,
		ReturnStatus: 0,
		AuditStatus:  3,
		IsActive:     1,
	}
	r := o.Read(&equips_date)
	if r != nil {

	}
	x := equips_date.Mount
	_, err := o.Insert(&equips_apply)
	o.QueryTable("sys_equips").Filter("id", labs_id).Update(orm.Params{
		"mount": x - mount,
	})

	message_map := map[string]interface{}{}
	if err != nil {
		message_map["code"] = 10001
		message_map["msg"] = "申请失败"
	}

	message_map["code"] = 200
	message_map["msg"] = "申请成功"

	c.Data["json"] = message_map
	c.ServeJSON()

}

func (c *EquipsApplyController) MyApply() {
	o := orm.NewOrm()

	qs := o.QueryTable("sys_equips_apply")

	equips_apply := []auth.EquipsApply{}

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

	uid := c.GetSession("id")

	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	if kw != "" { // 有查询条件的
		// 总数
		count, _ = qs.Filter("is_delete", 0).Filter("equips__name__contains", kw).Filter("user_id", uid.(int)).Count()
		qs.Filter("is_delete", 0).Filter("equips__name__contains", kw).Limit(pagePerNum).Offset(offsetNum).RelatedSel().Filter("user_id", uid.(int)).All(&equips_apply)
	} else {
		count, _ = qs.Filter("is_delete", 0).Filter("user_id", uid.(int)).Count()
		qs.Filter("is_delete", 0).Limit(pagePerNum).Filter("user_id", uid.(int)).Offset(offsetNum).RelatedSel().All(&equips_apply)

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
	c.Data["equips_apply"] = equips_apply
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw
	c.TplName = "equips/my_apply_list.html"
}

func (c *EquipsApplyController) AuditApply() {
	o := orm.NewOrm()

	qs := o.QueryTable("sys_equips_apply")

	equips_apply := []auth.EquipsApply{}

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
		count, _ = qs.Filter("is_delete", 0).Filter("equips__name__contains", kw).Count()
		qs.Filter("is_delete", 0).Filter("equips__name__contains", kw).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&equips_apply)
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&equips_apply)

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
	c.Data["equips_apply"] = equips_apply
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw
	c.TplName = "equips/audit_apply_list.html"

}

func (c *EquipsApplyController) ToAuditApply() {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_equips_apply")
	equips_apply := auth.EquipsApply{}
	labs := auth.Labs{}
	qs.Filter("id", id).One(&equips_apply)
	c.Data["equips_apply"] = equips_apply
	c.Data["labs"] = labs
	c.TplName = "equips/audit_apply.html"

}

func (c *EquipsApplyController) DoAuditApply() {

	option := c.GetString("option")
	audit_status, _ := c.GetInt("audit_status")

	id, _ := c.GetInt("id")

	o := orm.NewOrm()

	qs := o.QueryTable("sys_equips_apply")
	_, err := qs.Filter("id", id).Update(orm.Params{
		"audit_option": option,
		"audit_status": audit_status,
	})

	message_map := map[string]interface{}{}
	if err != nil {
		message_map["code"] = 10001
		message_map["msg"] = "审核失败"
	}

	message_map["code"] = 200
	message_map["msg"] = "审核成功"

	c.Data["json"] = message_map
	c.ServeJSON()

}

func (c *EquipsApplyController) DoReturn() {

	id, _ := c.GetInt("id")
	mount, _ := c.GetInt("mount")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_equips_apply")
	qs.Filter("id", id).Update(orm.Params{
		"return_status": 1,
	})
	x := auth.EquipsApply{}
	qs.Filter("id", id).RelatedSel().One(&x)
	a := x.Equips.Mount

	equips_apply := auth.EquipsApply{}
	qs.Filter("id", id).One(&equips_apply)
	o.QueryTable("sys_equips").Filter("id", equips_apply.Labs.Id).Update(orm.Params{
		"mount": a + mount,
	})

	c.Redirect(beego.URLFor("EquipsApplyController.MyApply"), 302)

}
