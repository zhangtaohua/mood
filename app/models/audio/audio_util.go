package audio

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (audio Audio) {
	database.DB.Where("id", idstr).First(&audio)
	return
}

func GetBy(field, value string) (audio Audio) {
	database.DB.Where("? = ?", field, value).First(&audio)
	return
}

func All() (audio []Audio) {
	database.DB.Find(&audio)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Audio{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (audio []Audio, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Audio{}),
		whereFields,
		&audio,
		app.V1URL(database.TableName(&Audio{})),
		perPage,
	)
	return
}
