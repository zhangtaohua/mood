// Package topic_category 模型
package topic_category

import (
	"mood/app/models"
	"mood/app/models/user"

	"mood/pkg/database"
)

type TopicCategory struct {
	models.BaseModel

	// Put fields in here
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	models.CommonTimestampsField
}

func (topicCategory *TopicCategory) Create() {
	database.DB.Create(&topicCategory)
}

func (topicCategory *TopicCategory) Save() (rowsAffected int64) {
	result := database.DB.Save(&topicCategory)
	return result.RowsAffected
}

func (topicCategory *TopicCategory) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topicCategory)
	return result.RowsAffected
}
