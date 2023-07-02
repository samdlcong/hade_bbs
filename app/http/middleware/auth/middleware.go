package auth

import (
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/gin"
	"hade_bbs/app/provider/user"
)

// AuthMiddleware 登录中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		envService := c.MustMake(contract.EnvKey).(contract.Env)
		userService := c.MustMake(user.UserKey).(user.Service)
		// 如果在调试模式下，根据参数的 user_id 获取信息
		if envService.AppEnv() == contract.EnvDevelopment {
			userID, exist := c.DefaultQueryInt64("user_id", 0)
			if exist {
				authUser, _ := userService.GetUser()
			}
		}
	}
}
