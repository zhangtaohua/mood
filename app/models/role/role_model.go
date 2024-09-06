// Package role 模型
package role

import (
	"mood/app/models"
	"mood/pkg/database"
)

type Role struct {
	models.BaseModel

	// Put fields in here
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`

	models.CommonTimestampsField
}

func (role *Role) Create() {
	database.DB.Create(&role)
}

func (role *Role) Save() (rowsAffected int64) {
	result := database.DB.Save(&role)
	return result.RowsAffected
}

func (role *Role) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&role)
	return result.RowsAffected
}
