package diary

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (diary Diary) {
	database.DB.Where("id", idstr).First(&diary)
	return
}

func GetBy(field, value string) (diary Diary) {
	database.DB.Where("? = ?", field, value).First(&diary)
	return
}

func All() (diaries []Diary) {
	database.DB.Find(&diaries)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Diary{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (diaries []Diary, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Diary{}),
		whereFields,
		&diaries,
		app.V1URL(database.TableName(&Diary{})),
		perPage,
	)
	return
}
