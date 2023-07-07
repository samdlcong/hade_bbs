package qa

import (
	"github.com/gohade/hade/framework/gin"
	provide "hade_bbs/app/provider/qa"
)

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
	qaService := c.MustMake(provide.QaKey).(provide.Service)
	id, exist := c.DefaultQueryInt64("id", 0)
	if !exist {
		c.ISetStatus(400).IText("参数错误")
		return
	}

	question, err := qaService.GetQuestion(c, id)
	if err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}

	if err := qaService.QuestionLoadAuthor(c, question); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	if err := qaService.QuestionLoadAnswers(c, question); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	if err := qaService.AnswersLoadAuthor(c, &(question.Answers)); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}

	questionDTO := ConvertQuestionToDTO(question, nil)

	c.ISetOkStatus().IJsonp(questionDTO)
}
