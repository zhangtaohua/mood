// Package user 存放用户 Model 相关逻辑
package user

import (
	"mood/app/models"
	"mood/app/models/role"
	"mood/app/models/user_third_platform"
	"mood/pkg/database"
	"mood/pkg/hash"
)

// User 用户模型
// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name" default:""`
	NickName string `json:"nick_name"`

	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
	// omitempty

	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	RoleID string    `json:"role_id"`
	Role   role.Role `json:"role"`

	ThirdPlatformID string                                `json:"third_platform_id"`
	ThirdPlatform   user_third_platform.UserThirdPlatform `json:"third_platform"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
