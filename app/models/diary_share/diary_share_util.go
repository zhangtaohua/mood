package diary_share

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (diaryShare DiaryShare) {
	database.DB.Where("id", idstr).First(&diaryShare)
	return
}

func GetBy(field, value string) (diaryShare DiaryShare) {
	database.DB.Where("? = ?", field, value).First(&diaryShare)
	return
}

func All() (diaryShares []DiaryShare) {
	database.DB.Find(&diaryShares)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(DiaryShare{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (diaryShares []DiaryShare, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(DiaryShare{}),
		whereFields,
		&diaryShares,
		app.V1URL(database.TableName(&DiaryShare{})),
		perPage,
	)
	return
}
