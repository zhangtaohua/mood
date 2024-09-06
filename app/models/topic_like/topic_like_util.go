package topic_like

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (topicLike TopicLike) {
	database.DB.Where("id", idstr).First(&topicLike)
	return
}

func GetBy(field, value string) (topicLike TopicLike) {
	database.DB.Where("? = ?", field, value).First(&topicLike)
	return
}

func All() (topicLikes []TopicLike) {
	database.DB.Find(&topicLikes)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(TopicLike{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (topicLikes []TopicLike, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(TopicLike{}),
		whereFields,
		&topicLikes,
		app.V1URL(database.TableName(&TopicLike{})),
		perPage,
	)
	return
}
