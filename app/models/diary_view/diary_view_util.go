package diary_view

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (diaryView DiaryView) {
	database.DB.Where("id", idstr).First(&diaryView)
	return
}

func GetBy(field, value string) (diaryView DiaryView) {
	database.DB.Where("? = ?", field, value).First(&diaryView)
	return
}

func All() (diaryViews []DiaryView) {
	database.DB.Find(&diaryViews)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(DiaryView{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (diaryViews []DiaryView, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(DiaryView{}),
		whereFields,
		&diaryViews,
		app.V1URL(database.TableName(&DiaryView{})),
		perPage,
	)
	return
}
