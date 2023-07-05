package qa

import (
	"hade_bbs/app/http/module/user"
	"time"
)

// QuestionDTO 问题列表返回结构
type QuestionDTO struct {
	ID        int64         `json:"id,omitempty"`
	Title     string        `json:"title,omitempty"`
	Context   string        `json:"context,omitempty"`
	AnswerNum int           `json:"answer_num"`
	CrearedAt time.Time     `json:"creared_at"`
	UpdateAt  time.Time     `json:"update_at"`
	Author    *user.UserDTO `json:"author,omitempty"`
	Answers   []*AnswerDTO  `json:"answers,omitempty"`
}

// AnswerDTO 回答返回结构
type AnswerDTO struct {
	ID        int64     `json:"id,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`

	Author *user.UserDTO `json:"author,omitempty"`
}
