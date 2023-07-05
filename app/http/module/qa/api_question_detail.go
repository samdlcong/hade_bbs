package qa

import "github.com/gohade/hade/framework/gin"

// QuestionDetail 获取问题详情
// @Summary 获取问题详情
// @Description 获取问题详情，包括问题的所有回答
// @Accept json
// @Produce json
// @Tags qa
// @Param id query int true "问题id"
// @Success 200 QuestionDTO question "问题详情，带回答和作者"
// @Router /question/detail [get]
func (api *QAApi) QuestionDetail(c *gin.Context) {

}
