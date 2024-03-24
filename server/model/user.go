package model

type SysUser struct {
	Basic
	Uuid     uint   `json:"uuid"`
	UserName string `gorm:"type:varchar(32);comment:用户名" json:"user_name"`
	NickName string `gorm:"type:varchar(32);comment:昵称" json:"nick_name"`
	Password string `gorm:"type:varchar(256);comment:密码" json:"password"`
	Email    string `gorm:"type:varchar(256);comment:邮箱" json:"email"`
	IsBan    int    `gorm:"type:smallint;comment:是否被封禁" json:"is_ban"`
}
