package filter
import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"log"
	"net/http"
	"golang-gin-api/model"
)

func Auth(c *gin.Context) {
	cookie, e := c.Request.Cookie("uid")
	if e != nil {
		log.Printf("error: %s", e.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
		c.Abort()
		return
	}
	id := cookie.Value
	v, err := c.Get("dbMap")
	if err != nil {
		log.Printf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		c.Abort()
		return
	}
	dbMap, ok := v.(*gorp.DbMap)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		c.Abort()
		return
	}
	u, err := dbMap.Get(model.User{}, id)
	if err != nil {
		log.Printf("error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "db error"})
		c.Abort()
		return
	}
	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		c.Abort()
		return
	}
	user, ok := u.(*model.User)
	if !ok {
		log.Print("bad interface.")
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		c.Abort()
		return
	}
	log.Printf("authorized: %s", user.Name)
}
