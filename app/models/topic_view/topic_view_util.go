package topic_view

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (topicView TopicView) {
	database.DB.Where("id", idstr).First(&topicView)
	return
}

func GetBy(field, value string) (topicView TopicView) {
	database.DB.Where("? = ?", field, value).First(&topicView)
	return
}

func All() (topicViews []TopicView) {
	database.DB.Find(&topicViews)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(TopicView{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (topicViews []TopicView, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(TopicView{}),
		whereFields,
		&topicViews,
		app.V1URL(database.TableName(&TopicView{})),
		perPage,
	)
	return
}
