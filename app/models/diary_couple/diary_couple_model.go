// Package diary_couple 模型
package diary_couple

import (
	"mood/app/models"
	"mood/pkg/database"
)

type DiaryCouple struct {
	models.BaseModel

	// Put fields in here

	models.CommonTimestampsField
}

func (diaryCouple *DiaryCouple) Create() {
	database.DB.Create(&diaryCouple)
}

func (diaryCouple *DiaryCouple) Save() (rowsAffected int64) {
	result := database.DB.Save(&diaryCouple)
	return result.RowsAffected
}

func (diaryCouple *DiaryCouple) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&diaryCouple)
	return result.RowsAffected
}
