//Package diary_view 模型
package diary_view

import (
  "mood/app/models"
	"mood/app/models/diary"
	"mood/app/models/user"
	"mood/pkg/database"
)

type DiaryView struct {
    models.BaseModel

    // Put fields in here
    UserID string    `json:"user_id"`
    User   user.User `json:"user"`
  
    DiaryID string      `json:"diary_id"`
    Diary   diary.Diary `json:"diary"`

    models.CommonTimestampsField
}

func (diaryView *DiaryView) Create() {
    database.DB.Create(&diaryView)
}

func (diaryView *DiaryView) Save() (rowsAffected int64) {
    result := database.DB.Save(&diaryView)
    return result.RowsAffected
}

func (diaryView *DiaryView) Delete() (rowsAffected int64) {
    result := database.DB.Delete(&diaryView)
    return result.RowsAffected
}
