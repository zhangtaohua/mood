// Package video 模型
package video

import (
	"mood/app/models"
	"mood/app/models/user"

	"mood/pkg/database"
)

type Video struct {
	models.BaseModel

	// Put fields in here
	Url string `json:"url"`

	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	models.CommonTimestampsField
}

func (video *Video) Create() {
	database.DB.Create(&video)
}

func (video *Video) Save() (rowsAffected int64) {
	result := database.DB.Save(&video)
	return result.RowsAffected
}

func (video *Video) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&video)
	return result.RowsAffected
}
