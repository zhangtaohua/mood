package image

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (image Image) {
	database.DB.Where("id", idstr).First(&image)
	return
}

func GetBy(field, value string) (image Image) {
	database.DB.Where("? = ?", field, value).First(&image)
	return
}

func All() (images []Image) {
	database.DB.Find(&images)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Image{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (images []Image, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Image{}),
		whereFields,
		&images,
		app.V1URL(database.TableName(&Image{})),
		perPage,
	)
	return
}
