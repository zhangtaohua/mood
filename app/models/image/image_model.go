// Package image 模型
package image

import (
	"mood/app/models"
	"mood/app/models/user"

	"mood/pkg/database"
)

type Image struct {
	models.BaseModel

	// Put fields in here
	Url string `json:"url"`

	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	models.CommonTimestampsField
}

func (image *Image) Create() {
	database.DB.Create(&image)
}

func (image *Image) Save() (rowsAffected int64) {
	result := database.DB.Save(&image)
	return result.RowsAffected
}

func (image *Image) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&image)
	return result.RowsAffected
}
