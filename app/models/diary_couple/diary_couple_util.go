package diary_couple

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (diaryCouple DiaryCouple) {
	database.DB.Where("id", idstr).First(&diaryCouple)
	return
}

func GetBy(field, value string) (diaryCouple DiaryCouple) {
	database.DB.Where("? = ?", field, value).First(&diaryCouple)
	return
}

func All() (diaryCouples []DiaryCouple) {
	database.DB.Find(&diaryCouples)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(DiaryCouple{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (diaryCouples []DiaryCouple, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(DiaryCouple{}),
		whereFields,
		&diaryCouples,
		app.V1URL(database.TableName(&DiaryCouple{})),
		perPage,
	)
	return
}
