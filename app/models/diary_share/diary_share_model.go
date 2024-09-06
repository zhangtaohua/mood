// Package diary_share 模型
package diary_share

import (
	"mood/app/models"
	"mood/app/models/diary"
	"mood/app/models/user"
	"mood/pkg/database"
)

type DiaryShare struct {
	models.BaseModel

	// Put fields in here
	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	DiaryID string      `json:"diary_id"`
	Diary   diary.Diary `json:"diary"`

	models.CommonTimestampsField
}

func (diaryShare *DiaryShare) Create() {
	database.DB.Create(&diaryShare)
}

func (diaryShare *DiaryShare) Save() (rowsAffected int64) {
	result := database.DB.Save(&diaryShare)
	return result.RowsAffected
}

func (diaryShare *DiaryShare) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&diaryShare)
	return result.RowsAffected
}
