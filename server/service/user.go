package service

type UserService struct {
}

func (u *UserService) Login(userName string, password string) {
	//global.DB.Where("user_name = ? AND password = ?", userName, password).First(&SysUser{}

}
