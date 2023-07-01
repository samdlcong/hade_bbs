package user

import "hade_bbs/app/provider/user"

// ConvertUserToDTO 将 user 转换为 UserDTO
func ConvertUserToDTO(user *user.User) *UserDTO {
	if user == nil {
		return nil
	}
	return &UserDTO{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt,
	}
}
