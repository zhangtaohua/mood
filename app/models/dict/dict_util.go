package dict

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (dict Dict) {
	database.DB.Where("id", idstr).First(&dict)
	return
}

func GetBy(field, value string) (dict Dict) {
	database.DB.Where("? = ?", field, value).First(&dict)
	return
}

func All() (dicts []Dict) {
	database.DB.Find(&dicts)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Dict{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (dicts []Dict, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Dict{}),
		whereFields,
		&dicts,
		app.V1URL(database.TableName(&Dict{})),
		perPage,
	)
	return
}
