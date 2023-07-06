package qa

import "github.com/gohade/hade/framework/gin"

type answerCreateParam struct {
	QuestionID int64  `json:"question_id" binding:"required"`
	Context    string `json:"context" binding:"required"`
}

// AnswerCreate 代表创建回答
// @Summary 创建回答
// @Description 创建回答
// @Accpet json
// @Produce json
// @Tags qa
// @Param answerCreateParam body answerCreateParam true "创建回答参数"
// @Success 200 string Msg "操作成功"
// @Router /answer/create [post]
func (api *QAApi) AnswerCreate(c *gin.Context) {

}
