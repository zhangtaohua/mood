package topic_follow

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (topicFollow TopicFollow) {
	database.DB.Where("id", idstr).First(&topicFollow)
	return
}

func GetBy(field, value string) (topicFollow TopicFollow) {
	database.DB.Where("? = ?", field, value).First(&topicFollow)
	return
}

func All() (topicFollows []TopicFollow) {
	database.DB.Find(&topicFollows)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(TopicFollow{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (topicFollows []TopicFollow, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(TopicFollow{}),
		whereFields,
		&topicFollows,
		app.V1URL(database.TableName(&TopicFollow{})),
		perPage,
	)
	return
}
