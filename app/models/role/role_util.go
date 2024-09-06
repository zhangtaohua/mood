package role

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (role Role) {
	database.DB.Where("id", idstr).First(&role)
	return
}

func GetBy(field, value string) (role Role) {
	database.DB.Where("? = ?", field, value).First(&role)
	return
}

func All() (roles []Role) {
	database.DB.Find(&roles)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Role{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, preloadFields []paginator.BasePreloadField, whereFields []paginator.BaseWhereField, perPage int) (roles []Role, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Role{}),
		preloadFields,
		whereFields,
		&roles,
		app.V1URL(database.TableName(&Role{})),
		perPage,
	)
	return
}
