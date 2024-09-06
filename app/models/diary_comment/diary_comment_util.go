package diary_comment

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (diaryComment DiaryComment) {
	database.DB.Where("id", idstr).First(&diaryComment)
	return
}

func GetBy(field, value string) (diaryComment DiaryComment) {
	database.DB.Where("? = ?", field, value).First(&diaryComment)
	return
}

func All() (diaryComments []DiaryComment) {
	database.DB.Find(&diaryComments)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(DiaryComment{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (diaryComments []DiaryComment, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(DiaryComment{}),
		whereFields,
		&diaryComments,
		app.V1URL(database.TableName(&DiaryComment{})),
		perPage,
	)
	return
}
