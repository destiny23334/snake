package sys

import (
	"github.com/gin-gonic/gin"
	v1 "snake/api/v1"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	userApi := v1.UserApi{}
	{
		//userRouter.POST("register", s.Register)
		userRouter.POST("login", userApi.Login)
		//userRouter.GET("info", s.Info)
		//userRouter.POST("logout", s.Logout)
	}
}
