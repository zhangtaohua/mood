package topic_category

import (
	"mood/pkg/app"
	"mood/pkg/database"
	"mood/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (topicCategory TopicCategory) {
	database.DB.Where("id", idstr).First(&topicCategory)
	return
}

func GetBy(field, value string) (topicCategory TopicCategory) {
	database.DB.Where("? = ?", field, value).First(&topicCategory)
	return
}

func All() (topicCategories []TopicCategory) {
	database.DB.Find(&topicCategories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(TopicCategory{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, whereFields []interface{}, perPage int) (topicCategories []TopicCategory, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(TopicCategory{}),
		whereFields,
		&topicCategories,
		app.V1URL(database.TableName(&TopicCategory{})),
		perPage,
	)
	return
}
