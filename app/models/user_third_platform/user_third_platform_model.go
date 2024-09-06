// Package user_third_platform 模型
package user_third_platform

import (
	"mood/app/models"
	"mood/pkg/database"

	"gorm.io/datatypes"
)

type UserThirdPlatform struct {
	models.BaseModel

	UserID      string            `json:"user_id"`
	Platform    string            `json:"platform"`
	Information datatypes.JSONMap `json:"information"`

	// Put fields in here
	models.CommonTimestampsField
}

func (userThirdPlatform *UserThirdPlatform) Create() {
	database.DB.Create(&userThirdPlatform)
}

func (userThirdPlatform *UserThirdPlatform) Save() (rowsAffected int64) {
	result := database.DB.Save(&userThirdPlatform)
	return result.RowsAffected
}

func (userThirdPlatform *UserThirdPlatform) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&userThirdPlatform)
	return result.RowsAffected
}
