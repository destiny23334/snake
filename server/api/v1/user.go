package v1

import (
	"github.com/gin-gonic/gin"
	"snake/model/request"
)

type UserApi struct {
}

func (u *UserApi) Login(c *gin.Context) {
	var loginRequest request.Login // 从请求中获取参数
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "参数错误",
		})
		return
	}
	// 检查数据库是否匹配
	//userService.Login(user)
	// TODO: 获取token和过期时间
	// 返回结果
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "登录成功",
	})

}
