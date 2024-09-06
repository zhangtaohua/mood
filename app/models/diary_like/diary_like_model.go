// Package diary_like 模型
package diary_like

import (
	"mood/app/models"
	"mood/app/models/diary"
	"mood/app/models/user"
	"mood/pkg/database"
)

type DiaryLike struct {
	models.BaseModel

	// Put fields in here
	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	DiaryID string      `json:"diary_id"`
	Diary   diary.Diary `json:"diary"`

	models.CommonTimestampsField
}

func (diaryLike *DiaryLike) Create() {
	database.DB.Create(&diaryLike)
}

func (diaryLike *DiaryLike) Save() (rowsAffected int64) {
	result := database.DB.Save(&diaryLike)
	return result.RowsAffected
}

func (diaryLike *DiaryLike) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&diaryLike)
	return result.RowsAffected
}
