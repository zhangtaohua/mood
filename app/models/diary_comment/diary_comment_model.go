// Package diary_comment 模型
package diary_comment

import (
	"mood/app/models"
	"mood/app/models/diary"
	"mood/app/models/user"

	"mood/pkg/database"
)

type DiaryComment struct {
	models.BaseModel

	// Put fields in here
	Comment string `json:"comment,omitempty"`

	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	DiaryID string      `json:"diary_id"`
	Diary   diary.Diary `json:"diary"`

	models.CommonTimestampsField
}

func (diaryComment *DiaryComment) Create() {
	database.DB.Create(&diaryComment)
}

func (diaryComment *DiaryComment) Save() (rowsAffected int64) {
	result := database.DB.Save(&diaryComment)
	return result.RowsAffected
}

func (diaryComment *DiaryComment) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&diaryComment)
	return result.RowsAffected
}
