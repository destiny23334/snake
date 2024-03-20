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
			"data":   gin.H{"msg": "参数错误"},
		})
		return
	}
	// 检查数据库是否匹配
	result := userService.Login(loginRequest.UserName, loginRequest.Password)
	// TODO: 获取token和过期时间
	// 返回结果
	if result {
		c.JSON(200, gin.H{
			"status": 0,
			"data":   gin.H{"msg": "登录成功"},
		})
	} else {
		c.JSON(200, gin.H{
			"status": 0,
			"data":   gin.H{"msg": "用户名或密码错误"},
		})
	}

}
