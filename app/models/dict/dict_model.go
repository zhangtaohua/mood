// Package dict 模型
package dict

import (
	"mood/app/models"
	"mood/pkg/database"
)

type Dict struct {
	models.BaseModel

	// Put fields in here
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	Key   string `json:"key"`
	Value string `json:"value"`

	models.CommonTimestampsField
}

func (dict *Dict) Create() {
	database.DB.Create(&dict)
}

func (dict *Dict) Save() (rowsAffected int64) {
	result := database.DB.Save(&dict)
	return result.RowsAffected
}

func (dict *Dict) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&dict)
	return result.RowsAffected
}
