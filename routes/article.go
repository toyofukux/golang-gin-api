package routes
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/go-gorp/gorp"
	"golang-gin-api/model"
	"log"
	"golang-gin-api/form"
)

func GetArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	v, err := c.Get("dbMap")
	if err != nil {
		log.Printf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	dbMap, ok := v.(*gorp.DbMap)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	art, err := dbMap.Get(model.Article{}, id)
	if err != nil {
		log.Printf("error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "db error"})
		return
	}
	if art == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "article": art})
}

func PostArticle(c *gin.Context) {
	var json form.ArticleJSON
	c.Bind(&json)
	article := model.NewArticle(json.Title, json.Content)
	if !model.ValidArticle(&article) {
		log.Printf("validate error.")
		c.JSON(http.StatusBadRequest, gin.H{"status": "validate error"})
		return
	}
	v, err := c.Get("dbMap")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
	}
	dbMap, ok := v.(*gorp.DbMap)
	if ok {
		err := dbMap.Insert(&article)
		if err != nil {
			log.Printf("error: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": "db error"})
			return
		}
	} else {
		// error
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
