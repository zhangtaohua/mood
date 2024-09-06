// Package topic_comment 模型
package topic_comment

import (
	"mood/app/models"
	"mood/app/models/topic"
	"mood/app/models/user"

	"mood/pkg/database"
)

type TopicComment struct {
	models.BaseModel

	// Put fields in here
	Comment    string `json:"comment,omitempty"`
	DebateSide string `json:"debate_side"`

	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	TopicID string      `json:"topic_id"`
	Topic   topic.Topic `json:"topic"`

	models.CommonTimestampsField
}

func (topicComment *TopicComment) Create() {
	database.DB.Create(&topicComment)
}

func (topicComment *TopicComment) Save() (rowsAffected int64) {
	result := database.DB.Save(&topicComment)
	return result.RowsAffected
}

func (topicComment *TopicComment) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topicComment)
	return result.RowsAffected
}
