package routers

import (
	"CLMS/controllers"
	"CLMS/controllers/auth"
	"CLMS/controllers/equips"
	"CLMS/controllers/labs"
	"CLMS/controllers/login"
	"CLMS/controllers/news"
	"CLMS/controllers/user"
	"github.com/astaxie/beego"
)

func init() {
	// 不需要登录既可请求的url
	beego.Router("/", &login.LoginController{})
	beego.Router("/main/user/log_out", &login.LoginController{}, "get:LogOut")
	beego.Router("/change_captcha", &login.LoginController{}, "get:ChangeCaptcha")

	// 必须登录才可请求的url

	// 后台首页
	beego.Router("/main/index", &controllers.HomeController{})
	beego.Router("/main/index/notify", &controllers.HomeController{}, "get:NotifyList")
	beego.Router("/main/index/read_notify", &controllers.HomeController{}, "get:ReadNotify")
	beego.Router("/main/index/delete_notify", &controllers.HomeController{}, "get:Delete")
	beego.Router("/main/welcome", &controllers.HomeController{}, "get:Welcome")

	// user模块
	beego.Router("/main/user/list", &user.UserController{}, "get:List")
	beego.Router("/main/user/to_add", &user.UserController{}, "get:ToAdd")
	beego.Router("/main/user/do_add", &user.UserController{}, "post:DoAdd")
	beego.Router("/main/user/is_active", &user.UserController{}, "post:IsActive")
	beego.Router("/main/user/delete", &user.UserController{}, "get:Delete")
	beego.Router("/main/user/reset_pwd", &user.UserController{}, "get:ResetPassword")
	beego.Router("/main/user/to_edit", &user.UserController{}, "get:ToUpdate")
	beego.Router("/main/user/do_edit", &user.UserController{}, "post:DoUpdate")
	beego.Router("/main/user/muli_delete", &user.UserController{}, "post:MuliDelete")

	// auth模块
	beego.Router("/main/auth/list", &auth.AuthController{}, "get:List")
	beego.Router("/main/auth/to_add", &auth.AuthController{}, "get:ToAdd")
	beego.Router("/main/auth/do_add", &auth.AuthController{}, "post:DoAdd")
	beego.Router("/main/auth/to_edit", &auth.AuthController{}, "get:ToUpdate")
	beego.Router("/main/auth/do_edit", &auth.AuthController{}, "post:DoUpdate")
	beego.Router("/main/auth/is_active", &auth.AuthController{}, "post:IsActive")
	// 角色模块
	beego.Router("/main/role/list", &auth.RoleController{}, "get:List")
	beego.Router("/main/role/to_add", &auth.RoleController{}, "get:ToAdd")
	beego.Router("/main/role/do_add", &auth.RoleController{}, "post:DoAdd")
	beego.Router("/main/role/muli_delete", &auth.RoleController{}, "post:MuliDelete")
	beego.Router("/main/role/is_active", &auth.RoleController{}, "post:IsActive")
	beego.Router("/main/role/to_edit", &auth.RoleController{}, "get:ToUpdate")
	beego.Router("/main/role/do_edit", &auth.RoleController{}, "post:DoUpdate")
	beego.Router("/main/role/delete", &auth.RoleController{}, "get:Delete")

	// 角色--用户
	beego.Router("/main/role/to_role_user_add", &auth.RoleController{}, "get:ToRoleUser")
	beego.Router("/main/role/do_role_user_add", &auth.RoleController{}, "post:DoRoleUser")

	// 角色--权限
	beego.Router("/main/role/to_role_auth_add", &auth.RoleController{}, "get:ToRoleAuth")
	beego.Router("/main/role/get_auth_json", &auth.RoleController{}, "get:GetAuthJson")
	beego.Router("/main/role/do_role_auth_add", &auth.RoleController{}, "post:DoRoleAuth")

	// 个人中心
	beego.Router("/main/user/my_center", &user.MyCenterController{})

	// 内容管理
	beego.Router("/main/news/category_list", &news.CategoryController{})
	beego.Router("/main/news/to_add_category", &news.CategoryController{}, "get:ToAdd")
	beego.Router("/main/news/do_add_category", &news.CategoryController{}, "post:DoAdd")
	beego.Router("/main/news/delete", &news.CategoryController{}, "get:Delete")
	beego.Router("/main/news/is_active", &news.CategoryController{}, "post:IsActive")
	beego.Router("/main/news/to_edit", &news.CategoryController{}, "get:ToUpdate")
	beego.Router("/main/news/do_edit", &news.CategoryController{}, "post:DoUpdate")
	beego.Router("/main/news/muli_delete", &news.CategoryController{}, "post:MuliDelete")

	beego.Router("/main/news/news_list", &news.NewsController{})
	beego.Router("/main/news/to_news_addt", &news.NewsController{}, "get:ToAdd")
	beego.Router("/main/news/do_news_addt", &news.NewsController{}, "post:DoAdd")
	beego.Router("/main/news/upload_img", &news.NewsController{}, "post:UploadImg")
	beego.Router("/main/news/newto_edit", &news.NewsController{}, "get:ToEdit")
	beego.Router("/main/news/newdo_edit", &news.NewsController{}, "post:DoEdit")
	beego.Router("/main/news/newdelete", &news.NewsController{}, "get:Delete")
	beego.Router("/main/news/newmuli_delete", &news.NewsController{}, "post:MuliDelete")
	beego.Router("/main/news/newfind", &news.NewsController{}, "get:Find")

	// 实验室管理模块
	beego.Router("/main/labs/lab_brand_list", &labs.LabsBrandController{})
	beego.Router("/main/labs/to_lab_brand_add", &labs.LabsBrandController{}, "get:ToAdd")
	beego.Router("/main/labs/do_lab_brand_add", &labs.LabsBrandController{}, "post:DoAdd")
	beego.Router("/main/labs/brandmuli_delete", &labs.LabsBrandController{}, "post:MuliDelete")
	beego.Router("/main/labs/branddelete", &labs.LabsBrandController{}, "get:Delete")
	beego.Router("/main/labs/brandis_active", &labs.LabsBrandController{}, "post:IsActive")
	beego.Router("/main/labs/brandto_edit", &labs.LabsBrandController{}, "get:ToUpdate")
	beego.Router("/main/labs/branddo_edit", &labs.LabsBrandController{}, "post:DoUpdate")

	beego.Router("/main/labs/labs_list", &labs.LabsController{})
	beego.Router("/main/labs/to_labs_add", &labs.LabsController{}, "get:ToAdd")
	beego.Router("/main/labs/do_labs_add", &labs.LabsController{}, "post:DoAdd")
	beego.Router("/main/labs/muli_delete", &labs.LabsController{}, "post:MuliDelete")
	beego.Router("/main/labs/delete", &labs.LabsController{}, "get:Delete")
	beego.Router("/main/labs/is_active", &labs.LabsController{}, "post:IsActive")
	beego.Router("/main/labs/to_edit", &labs.LabsController{}, "get:ToUpdate")
	beego.Router("/main/labs/do_edit", &labs.LabsController{}, "post:DoUpdate")

	beego.Router("/main/labs/labs_apply_list", &labs.LabsApplyController{})
	beego.Router("/main/labs/to_labs_apply", &labs.LabsApplyController{}, "get:ToApply")
	beego.Router("/main/labs/do_labs_apply", &labs.LabsApplyController{}, "post:DoApply")
	beego.Router("/main/labs/my_apply", &labs.LabsApplyController{}, "get:MyApply")
	beego.Router("/main/labs/audit_apply", &labs.LabsApplyController{}, "get:AuditApply")
	beego.Router("/main/labs/to_audit_apply", &labs.LabsApplyController{}, "get:ToAuditApply")
	beego.Router("/main/labs/do_audit_apply", &labs.LabsApplyController{}, "post:DoAuditApply")
	beego.Router("/main/labs/do_return", &labs.LabsApplyController{}, "get:DoReturn")

	//实验室器材管理模块
	beego.Router("/main/equips/equip_brand_list", &equips.EquipsBrandController{})
	beego.Router("/main/equips/to_equip_brand_add", &equips.EquipsBrandController{}, "get:ToAdd")
	beego.Router("/main/equips/do_equip_brand_add", &equips.EquipsBrandController{}, "post:DoAdd")
	beego.Router("/main/equips/brandmuli_delete", &equips.EquipsBrandController{}, "post:MuliDelete")
	beego.Router("/main/equips/branddelete", &equips.EquipsBrandController{}, "get:Delete")
	beego.Router("/main/equips/brandis_active", &equips.EquipsBrandController{}, "post:IsActive")
	beego.Router("/main/equips/brandto_edit", &equips.EquipsBrandController{}, "get:ToUpdate")
	beego.Router("/main/equips/branddo_edit", &equips.EquipsBrandController{}, "post:DoUpdate")

	beego.Router("/main/equips/equips_list", &equips.EquipsController{})
	beego.Router("/main/equips/to_equips_add", &equips.EquipsController{}, "get:ToAdd")
	beego.Router("/main/equips/do_equips_add", &equips.EquipsController{}, "post:DoAdd")
	beego.Router("/main/equips/muli_delete", &equips.EquipsController{}, "post:MuliDelete")
	beego.Router("/main/equips/delete", &equips.EquipsController{}, "get:Delete")
	beego.Router("/main/equips/is_active", &equips.EquipsController{}, "post:IsActive")
	beego.Router("/main/equips/to_edit", &equips.EquipsController{}, "get:ToUpdate")
	beego.Router("/main/equips/do_edit", &equips.EquipsController{}, "post:DoUpdate")

	beego.Router("/main/equips/equips_apply_list", &equips.EquipsApplyController{})
	beego.Router("/main/equips/to_equips_apply", &equips.EquipsApplyController{}, "get:ToApply")
	beego.Router("/main/equips/do_equips_apply", &equips.EquipsApplyController{}, "post:DoApply")
	beego.Router("/main/equips/my_apply", &equips.EquipsApplyController{}, "get:MyApply")
	beego.Router("/main/equips/audit_apply", &equips.EquipsApplyController{}, "get:AuditApply")
	beego.Router("/main/equips/to_audit_apply", &equips.EquipsApplyController{}, "get:ToAuditApply")
	beego.Router("/main/equips/do_audit_apply", &equips.EquipsApplyController{}, "post:DoAuditApply")
	beego.Router("/main/equips/do_return", &equips.EquipsApplyController{}, "get:DoReturn")

}
