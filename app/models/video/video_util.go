package video

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (video Video) {
	database.DB.Where("id", idstr).First(&video)
	return
}

func GetBy(field, value string) (video Video) {
	database.DB.Where("? = ?", field, value).First(&video)
	return
}

func All() (videos []Video) {
	database.DB.Find(&videos)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Video{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (videos []Video, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Video{}),
		whereFields,
		&videos,
		app.V1URL(database.TableName(&Video{})),
		perPage,
	)
	return
}
