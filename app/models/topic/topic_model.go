// Package topic 模型
package topic

import (
	"mood/app/models"
	"mood/app/models/category"
	"mood/app/models/user"
	"mood/pkg/database"
)

type Topic struct {
	models.BaseModel

	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`

	DebateMode string `json:"debate_mode"`

	// 通过 user_id 关联用户
	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	// 通过 category_id 关联分类
	CategoryID string            `json:"category_id"`
	Category   category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
