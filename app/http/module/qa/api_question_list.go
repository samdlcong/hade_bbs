package qa

import "github.com/gohade/hade/framework/gin"

// QuestionList 获取问题列表
// @Summary 获取问题列表
// @Description 获取问题列表，包含作者信息，不包含回答
// @Accept json
// @Produce json
// @Tags qa
// @Param page query int false "列表页页数"
// @Param size query int false "列表页单页个数"
// @Success 200 array QuestionDTO questions "问题列表"
// @Router /question/list [get]
func (api *QAApi) QuestionList(c *gin.Context) {

}
