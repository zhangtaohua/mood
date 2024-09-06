package diary_category

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (diaryCategory DiaryCategory) {
	database.DB.Where("id", idstr).First(&diaryCategory)
	return
}

func GetBy(field, value string) (diaryCategory DiaryCategory) {
	database.DB.Where("? = ?", field, value).First(&diaryCategory)
	return
}

func All() (diaryCategories []DiaryCategory) {
	database.DB.Find(&diaryCategories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(DiaryCategory{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (diaryCategories []DiaryCategory, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(DiaryCategory{}),
		whereFields,
		&diaryCategories,
		app.V1URL(database.TableName(&DiaryCategory{})),
		perPage,
	)
	return
}
