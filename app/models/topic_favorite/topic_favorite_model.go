// Package topic_favorite 模型
package topic_favorite

import (
	"mood/app/models"
	"mood/app/models/topic"
	"mood/app/models/user"

	"mood/pkg/database"
)

type TopicFavorite struct {
	models.BaseModel

	// Put fields in here
	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	TopicID string      `json:"topic_id"`
	Topic   topic.Topic `json:"topic"`

	models.CommonTimestampsField
}

func (topicFavorite *TopicFavorite) Create() {
	database.DB.Create(&topicFavorite)
}

func (topicFavorite *TopicFavorite) Save() (rowsAffected int64) {
	result := database.DB.Save(&topicFavorite)
	return result.RowsAffected
}

func (topicFavorite *TopicFavorite) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topicFavorite)
	return result.RowsAffected
}
