// Package topic_like 模型
package topic_like

import (
	"mood/app/models"
	"mood/app/models/topic"
	"mood/app/models/user"

	"mood/pkg/database"
)

type TopicLike struct {
	models.BaseModel

	// Put fields in here
	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	TopicID string      `json:"topic_id"`
	Topic   topic.Topic `json:"topic"`

	models.CommonTimestampsField
}

func (topicLike *TopicLike) Create() {
	database.DB.Create(&topicLike)
}

func (topicLike *TopicLike) Save() (rowsAffected int64) {
	result := database.DB.Save(&topicLike)
	return result.RowsAffected
}

func (topicLike *TopicLike) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topicLike)
	return result.RowsAffected
}
