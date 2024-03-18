package router

import (
	"github.com/gin-gonic/gin"
	"snake/router/sys"
)

// GroupApp 新路由往这里加
type GroupApp struct {
	UserRouter sys.UserRouter
}

var GlobalRouter = new(GroupApp)

// Routers 注册路由
func Routers() *gin.Engine {
	router := gin.Default()

	userRouter := GlobalRouter.UserRouter

	curRouter := router.Group("/api/v1")
	{
		userRouter.InitUserRouter(curRouter)

	}

	return router
}
