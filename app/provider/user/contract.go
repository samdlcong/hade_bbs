package user

import (
	"context"
	"encoding/json"
	"time"
)

const UserKey = "user"

// Service 用户相关服务
type Service interface {
	// Register 注册用户
	// 参数：user必填，username, password, email
	// 返回值：user 带上token
	Register(ctx context.Context, user *User) (*User, error)
	// SendRegisterMail 发送注册邮件
	// 参数：user必填 username, password, email, token
	SendRegisterMail(ctx context.Context, user *User) error
}

// User 代表一个用户，注意这里的用户信息字段在不同接口和参数可能为空
type User struct {
	ID        int64     `gorm:"colunm:id;primaryKey"` // 代表用户唯一ID
	UserName  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email"`
	CreatedAt time.Time `gorm:"column:created_at"`

	Token string `gorm:"-"` // token
}

// MarshalBinary 实现 BinaryMarshaler 接口
func (b *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

// UnmarshalBinary 实现 BinaryUnMarshaler 接口
func (b *User) UnmarshalBinary(bt []byte) error {
	return json.Unmarshal(bt, b)
}
