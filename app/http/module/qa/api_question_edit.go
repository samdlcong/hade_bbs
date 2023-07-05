package qa

import "github.com/gohade/hade/framework/gin"

type questionEditParam struct {
	ID      int64  `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// QuestionEdit 编辑问题
// @Summary 编辑问题
// @Description 编辑问题
// @Accept json
// @Produce json
// @Tags qa
// @Param questionEditParam body questionEditParam true "编辑问题参数"
// @Success 200 string Msg "操作成功"
// @Router /question/edit [post]
func (api *QAApi) QuestionEdit(c *gin.Context) {

}
