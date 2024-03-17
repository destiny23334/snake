package model

type SysUser struct {
	Basic
	UserName string
	NickName string
	Password string
	Email    string
	IsBan    int
}
