package diary_follow

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (diaryFollow DiaryFollow) {
	database.DB.Where("id", idstr).First(&diaryFollow)
	return
}

func GetBy(field, value string) (diaryFollow DiaryFollow) {
	database.DB.Where("? = ?", field, value).First(&diaryFollow)
	return
}

func All() (diaryFollows []DiaryFollow) {
	database.DB.Find(&diaryFollows)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(DiaryFollow{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (diaryFollows []DiaryFollow, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(DiaryFollow{}),
		whereFields,
		&diaryFollows,
		app.V1URL(database.TableName(&DiaryFollow{})),
		perPage,
	)
	return
}
