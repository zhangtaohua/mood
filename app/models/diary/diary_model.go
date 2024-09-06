// Package diary 模型
package diary

import (
	"mood/app/models"
	"mood/app/models/diary_category"
	"mood/app/models/user"

	"mood/pkg/database"

	"gorm.io/datatypes"
)

type Diary struct {
	models.BaseModel

	// Put fields in here
	Weather  string            `json:"weather"`
	Location datatypes.JSONMap `json:"location"`
	Title    string            `json:"title,omitempty" `
	Body     string            `json:"body"`
	Audio    []string          `json:"audio"`
	Image    []string          `json:"image"`
	Video    []string          `json:"video"`

	EmotionalScore   int    `json:"emotional_score"`
	EmotionalColor   string `json:"emotional_color"`
	EmotionalTrigger string `json:"emotional_trigger"`
	EmotionalRitual  string `json:"emotional_Ritual"`

	IsShared string `json:"is_shared"`

	UserID string    `json:"user_id"`
	User   user.User `json:"user"`

	CategoryID string                    `json:"category_id,omitempty"`
	Category   []diary_category.Category `json:"category"`

	models.CommonTimestampsField
}

func (diary *Diary) Create() {
	database.DB.Create(&diary)
}

func (diary *Diary) Save() (rowsAffected int64) {
	result := database.DB.Save(&diary)
	return result.RowsAffected
}

func (diary *Diary) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&diary)
	return result.RowsAffected
}
