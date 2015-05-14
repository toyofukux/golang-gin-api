package routes
import (
	"github.com/gin-gonic/gin"
	"golang-gin-api/app"
	"golang-gin-api/filter"
)

func Init(app *app.App) *gin.Engine {

	//router := gin.Default()
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	router.Use(app.AppendEnv)

	v1 := router.Group("/v1")
	{
		v1.Use(filter.Auth)
		v1.GET("/article/:id", GetArticle)
		v1.POST("/article", PostArticle)
	}

	return router
}
