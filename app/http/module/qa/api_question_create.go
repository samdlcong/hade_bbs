package qa

import "github.com/gohade/hade/framework/gin"

type questionCreateParam struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// QuestionCreate 代表创建问题
// @Summary 创建问题
// @Description 创建问题
// @Accept json
// @Produce json
// @Tags qa
// @Param questionCreateParam body questionCreateParam true "创建问题参数"
// @Success 200 string Msg "操作成功"
// @Router /question/create [post]
func (api *QAApi) QuestionCreate(c *gin.Context) {

}
