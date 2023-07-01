package user

import "github.com/gohade/hade/framework/gin"

type registerParam struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
	Email    string `json:"email" binding:"required,gte=6"`
}

// Register 注册
// @Summary 用户注册
// @Description 用户注册接口
// @Accept json
// @Produce json
// @Tags user
// @Param registerParam body registerParam true "注册参数"
// @Success 200 string Message "注册成功"
// @Router /user/register [post]
func (api *UserApi) Register(c *gin.Context) {

}
