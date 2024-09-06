// Package diary_favorite 模型
package diary_favorite

import (
	"mood/app/models"
	"mood/app/models/diary"
	"mood/app/models/user"
	"mood/pkg/database"
)

type DiaryFavorite struct {
	models.BaseModel

	// Put fields in here
	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	DiaryID string      `json:"diary_id"`
	Diary   diary.Diary `json:"diary"`

	models.CommonTimestampsField
}

func (diaryFavorite *DiaryFavorite) Create() {
	database.DB.Create(&diaryFavorite)
}

func (diaryFavorite *DiaryFavorite) Save() (rowsAffected int64) {
	result := database.DB.Save(&diaryFavorite)
	return result.RowsAffected
}

func (diaryFavorite *DiaryFavorite) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&diaryFavorite)
	return result.RowsAffected
}
