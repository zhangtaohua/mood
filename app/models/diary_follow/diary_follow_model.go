// Package diary_follow 模型
package diary_follow

import (
	"mood/app/models"
	"mood/app/models/diary"
	"mood/app/models/user"
	"mood/pkg/database"
)

type DiaryFollow struct {
	models.BaseModel

	// Put fields in here
	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	DiaryID string      `json:"diary_id"`
	Diary   diary.Diary `json:"diary"`

	models.CommonTimestampsField
}

func (diaryFollow *DiaryFollow) Create() {
	database.DB.Create(&diaryFollow)
}

func (diaryFollow *DiaryFollow) Save() (rowsAffected int64) {
	result := database.DB.Save(&diaryFollow)
	return result.RowsAffected
}

func (diaryFollow *DiaryFollow) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&diaryFollow)
	return result.RowsAffected
}
