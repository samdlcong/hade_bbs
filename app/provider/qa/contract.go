package qa

import (
	"context"
	"gorm.io/gorm"
	"hade_bbs/app/provider/user"
	"time"
)

const QaKey = "qa"

type Service interface {
	// GetQuestions 获取问题列表
	GetQuestions(ctx context.Context, pager *Pager) ([]*Question, error)
	// GetQuestion 获取问题
	GetQuestion(ctx context.Context, questionID int64) (*Question, error)
	// PostQuestion 提交问题
	PostQuestion(ctx context.Context, question *Question) error

	// QuestionLoadAuthor 问题加载Author字段
	QuestionLoadAuthor(ctx context.Context, question *Question) error
	// QuestionsLoadAuthor 批量加载Author字段
	QuestionsLoadAuthor(ctx context.Context, questions *[]*Question) error
	// QuestionLoadAnswers 获取单个问题答案
	QuestionLoadAnswers(ctx context.Context, question *Question) error
	// QuestionsLoadAnswers 批量获取问题答案
	QuestionsLoadAnswers(ctx context.Context, questions *[]*Question) error

	// PostAnswer 上传回答
	PostAnswer(ctx context.Context, answer *Answer) error
	// GetAnswer 获取回答
	GetAnswer(ctx context.Context, answerID int64) (*Answer, error)

	// AnswerLoadAuthor 回答加载Author字段
	AnswerLoadAuthor(ctx context.Context, question *Answer) error
	// AnswersLoadAuthor 批量加载 Author字段
	AnswersLoadAuthor(ctx context.Context, questions *[]*Answer) error

	// DeleteQuestion 删除问题，同时删除对应的回答
	DeleteQuestion(ctx context.Context, questionID int64) error
	// DeleteAnswer 删除某个回答
	DeleteAnswer(ctx context.Context, answerID int64) error
	// UpdateQuestion 代表更新问题
	updateQuestion(ctx context.Context, question *Question) error
}

// Question
type Question struct {
	ID        int64          `gorm:"column:id;primaryKey"`
	Title     string         `gorm:"column:title;comment:标题"`
	Context   string         `gorm:"column:context;comment:内容"`
	AuthorID  int64          `gorm:"column:author_id;comment:作者id;not null;default:0"`
	AnswerNum int            `gorm:"column:answer_num;comment:回答数;not null;default:0"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;autoCreateTime;<-:false;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Author    *user.User     `gorm:"foreignKey:AuthorID"`
	Answers   []*Answer      `gorm:"foreignKey:QuestionID"`
}

// Answer
type Answer struct {
	ID         int64          `gorm:"column:id;primaryKey"`
	QuestionID int64          `gorm:"column:question_id;index;comment:问题id;not null;default 0"`
	Context    string         `gorm:"column:context;comment:内容"`
	AuthorID   int64          `gorm:"column:author_id;comment:作者id;not null;default:0"`
	CreatedAt  time.Time      `gorm:"column:created_at;autoCreateTime;comment:创建时间"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;autoUpdateTime;autoCreateTime;<-:false;comment:更新时间"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Author     *user.User     `gorm:"foreignKey:AuthorID"`
	Question   *Question      `gorm:"foreignKey:QuestionID"`
}

// Pager
type Pager struct {
	Total int64 // 共有多少数据 只有返回值使用
	Start int   // 起始位置
	size  int   // 每个页面个数
}
