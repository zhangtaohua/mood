package diary_favorite

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (diaryFavorite DiaryFavorite) {
	database.DB.Where("id", idstr).First(&diaryFavorite)
	return
}

func GetBy(field, value string) (diaryFavorite DiaryFavorite) {
	database.DB.Where("? = ?", field, value).First(&diaryFavorite)
	return
}

func All() (diaryFavorites []DiaryFavorite) {
	database.DB.Find(&diaryFavorites)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(DiaryFavorite{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (diaryFavorites []DiaryFavorite, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(DiaryFavorite{}),
		whereFields,
		&diaryFavorites,
		app.V1URL(database.TableName(&DiaryFavorite{})),
		perPage,
	)
	return
}
