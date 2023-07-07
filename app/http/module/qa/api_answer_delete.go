package qa

import (
	"github.com/gohade/hade/framework/gin"
	"hade_bbs/app/http/middleware/auth"
	provider "hade_bbs/app/provider/qa"
)

// AnswerDelete 代表删除回答
// @Summary 创建回答
// @Description 创建回答
// @Accept json
// @Procude json
// @Tags qa
// @Param id query int true "删除id"
// @Success 200 string Msg "操作成功"
// @Router /answer/delete [get]
func (api *QAApi) AnswerDelete(c *gin.Context) {
	qaService := c.MustMake(provider.QaKey).(provider.Service)
	id, exist := c.DefaultQueryInt64("id", 0)
	if !exist {
		c.ISetStatus(400).IText("参数错误")
		return
	}
	user := auth.GetAuthUser(c)

	answer, err := qaService.GetAnswer(c, id)
	if err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	if answer.AuthorID != user.ID {
		c.ISetStatus(500).IText("没有权限做此操作")
		return
	}
	c.ISetOkStatus().IText("操作成功")
}
