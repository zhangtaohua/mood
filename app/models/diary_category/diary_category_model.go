// Package diary_category 模型
package diary_category

import (
	"mood/app/models"
	"mood/app/models/user"

	"mood/pkg/database"
)

type DiaryCategory struct {
	models.BaseModel

	// Put fields in here
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	models.CommonTimestampsField
}

func (diaryCategory *DiaryCategory) Create() {
	database.DB.Create(&diaryCategory)
}

func (diaryCategory *DiaryCategory) Save() (rowsAffected int64) {
	result := database.DB.Save(&diaryCategory)
	return result.RowsAffected
}

func (diaryCategory *DiaryCategory) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&diaryCategory)
	return result.RowsAffected
}
