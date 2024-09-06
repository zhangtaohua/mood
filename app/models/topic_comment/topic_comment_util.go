package topic_comment

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (topicComment TopicComment) {
	database.DB.Where("id", idstr).First(&topicComment)
	return
}

func GetBy(field, value string) (topicComment TopicComment) {
	database.DB.Where("? = ?", field, value).First(&topicComment)
	return
}

func All() (topicComments []TopicComment) {
	database.DB.Find(&topicComments)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(TopicComment{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (topicComments []TopicComment, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(TopicComment{}),
		whereFields,
		&topicComments,
		app.V1URL(database.TableName(&TopicComment{})),
		perPage,
	)
	return
}
