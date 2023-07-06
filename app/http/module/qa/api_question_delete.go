package qa

import "github.com/gohade/hade/framework/gin"

// QuestionDelete 删除问卷
// @Summary 删除问题
// @Description 删除问题，同时删除问题中的所有答案
// @Accept json
// @Produce json
// @Tags qa
// @Param id query int true "删除id"
// @Success 200 string Msg "操作成功"
// @Router /question/delete [get]
func (api *QAApi) QuestionDelete(c *gin.Context) {

}
