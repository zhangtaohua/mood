// Package topic_view 模型
package topic_view

import (
	"mood/app/models"
	"mood/app/models/topic"
	"mood/app/models/user"

	"mood/pkg/database"
)

type TopicView struct {
	models.BaseModel

	// Put fields in here
	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	TopicID string      `json:"topic_id"`
	Topic   topic.Topic `json:"topic"`

	models.CommonTimestampsField
}

func (topicView *TopicView) Create() {
	database.DB.Create(&topicView)
}

func (topicView *TopicView) Save() (rowsAffected int64) {
	result := database.DB.Save(&topicView)
	return result.RowsAffected
}

func (topicView *TopicView) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topicView)
	return result.RowsAffected
}
