package system

import (
	"github.com/jasvtfvan/oms-admin/server/model/common"
)

type SysUser struct {
	common.BaseModel
	Username  string     `json:"username" gorm:"index;not null;comment:用户名"`
	Password  string     `json:"-"  gorm:"not null;comment:密码"`
	NickName  string     `json:"nickName" gorm:"default:'';comment:用户昵称"`
	Avatar    string     `json:"avatar" gorm:"default:https://foruda.gitee.com/avatar/1710471233758250270/2074074_jasvtfvan_1710471233.png!avatar200;comment:头像"`
	Phone     string     `json:"phone"  gorm:"default:'';comment:手机号"`
	Email     string     `json:"email"  gorm:"default:'';comment:邮箱"`
	Enable    bool       `json:"enable" gorm:"default:true;comment:是否可用"`
	SysGroups []SysGroup `gorm:"many2many:sys_user_group;"`
	SysRoles  []SysRole  `gorm:"many2many:sys_user_role;"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
