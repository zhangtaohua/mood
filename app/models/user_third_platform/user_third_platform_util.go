package user_third_platform

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (userThirdPlatform UserThirdPlatform) {
	database.DB.Where("id", idstr).First(&userThirdPlatform)
	return
}

func GetBy(field, value string) (userThirdPlatform UserThirdPlatform) {
	database.DB.Where("? = ?", field, value).First(&userThirdPlatform)
	return
}

func All() (userThirdPlatforms []UserThirdPlatform) {
	database.DB.Find(&userThirdPlatforms)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(UserThirdPlatform{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (userThirdPlatforms []UserThirdPlatform, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(UserThirdPlatform{}),
		whereFields,
		&userThirdPlatforms,
		app.V1URL(database.TableName(&UserThirdPlatform{})),
		perPage,
	)
	return
}
