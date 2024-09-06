// Package topic_follow 模型
package topic_follow

import (
	"mood/app/models"
	"mood/app/models/topic"
	"mood/app/models/user"

	"mood/pkg/database"
)

type TopicFollow struct {
	models.BaseModel

	// Put fields in here
	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	TopicID string      `json:"topic_id"`
	Topic   topic.Topic `json:"topic"`

	models.CommonTimestampsField
}

func (topicFollow *TopicFollow) Create() {
	database.DB.Create(&topicFollow)
}

func (topicFollow *TopicFollow) Save() (rowsAffected int64) {
	result := database.DB.Save(&topicFollow)
	return result.RowsAffected
}

func (topicFollow *TopicFollow) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topicFollow)
	return result.RowsAffected
}
