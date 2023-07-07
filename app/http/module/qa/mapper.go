package qa

import (
	"github.com/PuerkitoBio/goquery"
	"hade_bbs/app/http/module/user"
	"hade_bbs/app/provider/qa"
	"strings"
)

// ConvertQuestionToDTO 将 question转换为 DTO
func ConvertQuestionToDTO(question *qa.Question, flags map[string]string) *QuestionDTO {
	if question == nil {
		return nil
	}
	author := user.ConvertUserToDTO(question.Author)
	if author == nil {
		author = &user.UserDTO{
			ID: question.AuthorID,
		}
	}

	context := question.Context
	if flags != nil {
		if isShortContext, ok := flags["is_short_context"]; ok && isShortContext == "true" {
			context = getShortContext(context)
		}
	}

	return &QuestionDTO{
		ID:        question.ID,
		Title:     question.Title,
		Context:   context,
		CrearedAt: question.CreatedAt,
		UpdateAt:  question.UpdatedAt,
		Author:    author,
		Answers:   ConvertAnswersToDTO(question.Answers),
	}
}

func getShortContext(context string) string {
	p := strings.NewReader(context)
	doc, _ := goquery.NewDocumentFromReader(p)

	doc.Find("script").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})

	text := doc.Text()
	if len(text) > 20 {
		text = text[:20] + "..."
	}
	return text
}

// ConvertQuestionsToDTO 将questions转换为 DTO
func ConvertQuestionsToDTO(questions []*qa.Question) []*QuestionDTO {
	if questions == nil {
		return nil
	}
	ret := make([]*QuestionDTO, 0, len(questions))
	for _, question := range questions {
		ret = append(ret, ConvertQuestionToDTO(question, map[string]string{"is_short_context": "true"}))
	}
	return ret
}

// ConvertAnswerToDTO 将answer转换为 DTO
func ConvertAnswerToDTO(answer *qa.Answer) *AnswerDTO {
	if answer == nil {
		return nil
	}
	author := user.ConvertUserToDTO(answer.Author)
	if author == nil {
		author = &user.UserDTO{
			ID: answer.AuthorID,
		}
	}
	return &AnswerDTO{
		ID:        answer.ID,
		Content:   answer.Context,
		CreatedAt: answer.CreatedAt,
		UpdateAt:  answer.UpdatedAt,
		Author:    author,
	}
}

// ConvertAnswersToDTO
func ConvertAnswersToDTO(answers []*qa.Answer) []*AnswerDTO {
	if answers == nil {
		return nil
	}

	ret := make([]*AnswerDTO, 0, len(answers))
	for _, answer := range answers {
		ret = append(ret, ConvertAnswerToDTO(answer))
	}
	return ret
}
