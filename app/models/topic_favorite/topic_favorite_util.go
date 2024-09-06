package topic_favorite

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (topicFavorite TopicFavorite) {
	database.DB.Where("id", idstr).First(&topicFavorite)
	return
}

func GetBy(field, value string) (topicFavorite TopicFavorite) {
	database.DB.Where("? = ?", field, value).First(&topicFavorite)
	return
}

func All() (topicFavorites []TopicFavorite) {
	database.DB.Find(&topicFavorites)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(TopicFavorite{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (topicFavorites []TopicFavorite, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(TopicFavorite{}),
		whereFields,
		&topicFavorites,
		app.V1URL(database.TableName(&TopicFavorite{})),
		perPage,
	)
	return
}
