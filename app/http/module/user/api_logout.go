package user

import (
	"github.com/gohade/hade/framework/gin"
	"hade_bbs/app/http/middleware/auth"
)

// Logout 代表登出
// @Summary 用户登出
// @Description 调用表示用户登出
// @Accpet json
// @Produce json
// @Tags user
// @Success 200 string Message "用户登出成功"
// @Router /user/logout [get]
func (api *UserApi) Logout(c *gin.Context) {
	authUser := auth.GetAuthUser(c)
	if authUser == nil {
		c.ISetStatus(500).IText("用户未登录")
		return
	}
	c.IRedirect("/#/login")
	return

}
