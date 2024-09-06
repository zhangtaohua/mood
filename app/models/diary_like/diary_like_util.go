package diary_like

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (diaryLike DiaryLike) {
	database.DB.Where("id", idstr).First(&diaryLike)
	return
}

func GetBy(field, value string) (diaryLike DiaryLike) {
	database.DB.Where("? = ?", field, value).First(&diaryLike)
	return
}

func All() (diaryLikes []DiaryLike) {
	database.DB.Find(&diaryLikes)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(DiaryLike{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (diaryLikes []DiaryLike, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(DiaryLike{}),
		whereFields,
		&diaryLikes,
		app.V1URL(database.TableName(&DiaryLike{})),
		perPage,
	)
	return
}
