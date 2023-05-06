package auth

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Auth struct {
	Id         int       `orm:"pk;auto"`
	AuthName   string    `orm:"description(权限名称);size(64)"`
	UrlFor     string    `orm:"description(url反转);size(255)"`
	Pid        int       `orm:"description(父节点id)"`
	Desc       string    `orm:"description(描述);size(255)"`
	CreateTime time.Time `orm:"type(datetime);auto_now;description(创建时间)"`
	IsActive   int       `orm:"column(is_active);description(1启用，0停用)"`
	IsDelete   int       `orm:"columns(is_delete);description(1删除，0未删除)"`
	Weight     int       `orm:"description(权重，数值越大，权重越大)"`
	Role       []*Role   `orm:"reverse(many)"`
}

type Role struct {
	Id         int       `orm:"pk;auto"`
	RoleName   string    `orm:"size(64)"`
	Desc       string    `orm:"size(255)"`
	IsActive   int       `orm:"column(is_active)"`
	IsDelete   int       `orm:"column(is_delete)"`
	CreateTime time.Time `orm:"type(datetime);auto_now;description(创建时间)"`
	Auth       []*Auth   `orm:"rel(m2m)"`
	User       []*User   `orm:"rel(m2m)"`
}

type User struct {
	Id            int              `orm:"pk;auto"`
	CardId        string           `orm:"size(64);column(card_id);description(用户工号)"`
	UserName      string           `orm:"unique;column(user_name);size(64);description(用户名)"`
	Password      string           `orm:"size(32);description(密码)"`
	Age           int              `orm:"null;description(年龄)"`
	Gender        int              `orm:"null;description(1:男,2:女,3:未知)"`
	Phone         int64            `orm:"null;description(电话号码)"`
	Addr          string           `orm:"null;size(255);description(地址)"`
	IsActive      int              `orm:"description(1启用，0停用);default(1)"`
	IsDelete      int              `orm:"description(1删除，0未删除);default(0)"`
	CreateTime    time.Time        `orm:"auto_now;type(datetime);description(创建时间);null"`
	Role          []*Role          `orm:"reverse(many)"`
	LabsApply     []*LabsApply     `orm:"reverse(many)"`
	MessageNotify []*MessageNotify `orm:"reverse(many)"`
}

// 实验室归属
type LabsUse struct {
	Id         int       `orm:"pk;auto"`
	Name       string    `orm:"description(实验室用法名称);size(64)"`
	Desc       string    `orm:"description(实验室用法描述);size(255)"`
	Labs       []*Labs   `orm:"reverse(many)"`
	IsActive   int       `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete   int       `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime time.Time `orm:"description(创建时间);type(datetime);auto_now"`
}

// 实验室列表
type Labs struct {
	Id          int            `orm:"pk;auto"`
	Name        string         `orm:"description(实验室名称);size(64)"`
	LabsUse     *LabsUse       `orm:"rel(fk);description(实验室用法外键)"`
	LabsApply   []*LabsApply   `orm:"reverse(many)"`
	Status      int            `orm:"default(0);description(0:可借,1:不可借)"`
	IsActive    int            `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete    int            `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime  time.Time      `orm:"description(创建时间);type(datetime);auto_now"`
	EquipsApply []*EquipsApply `orm:"reverse(many)"`
}

// 实验室申请、审核
type LabsApply struct {
	Id           int       `orm:"pk;auto"`
	User         *User     `orm:"rel(fk)"`
	Labs         *Labs     `orm:"rel(fk)"`
	Reason       string    `orm:"description(申请理由);size(255)"`
	Destination  string    `orm:"description(目的地);size(64)"`
	ReturnDate   time.Time `orm:"type(date);auto_now;description(归还日期)"`
	ReturnStatus int       `orm:"description(1:已归还，0：未归还);default(0)"`
	AuditStatus  int       `orm:"description(1:同意，2:未同意，3:未审批);default(3)"`
	AuditOption  string    `orm:"description(审批意见);size(255)"`
	IsActive     int       `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete     int       `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime   time.Time `orm:"description(创建时间);type(datetime);auto_now"`
	NotifyTag    int       `orm:"description(1:已发送通知，0：未发送通知);default(0)"`
}

// 器材用途
type EquipsUse struct {
	Id         int       `orm:"pk;auto"`
	Name       string    `orm:"description(器材用途名称);size(64)"`
	Desc       string    `orm:"description(器材用法描述);size(255)"`
	Equips     []*Equips `orm:"reverse(many)"`
	IsActive   int       `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete   int       `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime time.Time `orm:"description(创建时间);type(datetime);auto_now"`
}

// 器材列表
type Equips struct {
	Id          int            `orm:"pk;auto"`
	Name        string         `orm:"description(器材名称);size(64)"`
	EquipsUse   *EquipsUse     `orm:"rel(fk);description(器材用法外键)"`
	Mount       int            `orm:"description(器材数量);default(0)"`
	EquipsApply []*EquipsApply `orm:"reverse(many)"`
	Status      int            `orm:"default(0);description(0:可借,1:不可借)"`
	IsActive    int            `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete    int            `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime  time.Time      `orm:"description(创建时间);type(datetime);auto_now"`
}

// 器材申请、审核
type EquipsApply struct {
	Id           int       `orm:"pk;auto"`
	User         *User     `orm:"rel(fk)"`
	Equips       *Equips   `orm:"rel(fk)"`
	Labs         *Labs     `orm:"rel(fk)"`
	Mount        int       `orm:"description(申请数量);default(0)"`
	Reason       string    `orm:"description(申请理由);size(255)"`
	ReturnDate   time.Time `orm:"type(date);auto_now;description(归还日期)"`
	ReturnStatus int       `orm:"description(1:已归还，0：未归还);default(0)"`
	AuditStatus  int       `orm:"description(1:同意，2:未同意，3:未审批);default(3)"`
	AuditOption  string    `orm:"description(审批意见);size(255)"`
	IsActive     int       `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete     int       `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime   time.Time `orm:"description(创建时间);type(datetime);auto_now"`
	NotifyTag    int       `orm:"description(1:已发送通知，0：未发送通知);default(0)"`
}

// 消息通知
type MessageNotify struct {
	Id      int    `orm:"pk;auto"`
	Flag    int    `orm:"description(1:实验室借用逾期，2:所有通知);default(1)"`
	Title   string `orm:"size(64);description(消息标题)"`
	Content string `orm:"type(text);description(消息内容)"`
	User    *User  `orm:"rel(fk);description(用户外键)"`
	ReadTag int    `orm:"description(1:已读，0:未读)"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func (r *Auth) TableName() string {
	return "sys_auth"

}

func (r *Role) TableName() string {
	return "sys_role"

}

func (u *LabsUse) TableName() string {
	return "sys_labs_use"
}

func (u *Labs) TableName() string {
	return "sys_labs"
}

func (u *LabsApply) TableName() string {
	return "sys_labs_apply"
}

func (u *EquipsUse) TableName() string {
	return "sys_equips_use"
}

func (u *Equips) TableName() string {
	return "sys_equips"
}

func (u *EquipsApply) TableName() string {
	return "sys_equips_apply"
}
func (u *MessageNotify) TableName() string {
	return "sys_message_notify"
}

func init() {
	orm.RegisterModel(
		new(Auth),
		new(Role),
		new(User),
		new(LabsUse),
		new(Labs),
		new(LabsApply),
		new(EquipsUse),
		new(Equips),
		new(EquipsApply),
		new(MessageNotify),
	)
}
