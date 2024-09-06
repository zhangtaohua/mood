// Package audio 模型
package audio

import (
	"mood/app/models"
	"mood/app/models/user"

	"mood/pkg/database"
)

type Audio struct {
	models.BaseModel

	// Put fields in here
	Url string `json:"url"`

	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	models.CommonTimestampsField
}

func (audio *Audio) Create() {
	database.DB.Create(&audio)
}

func (audio *Audio) Save() (rowsAffected int64) {
	result := database.DB.Save(&audio)
	return result.RowsAffected
}

func (audio *Audio) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&audio)
	return result.RowsAffected
}
