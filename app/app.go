package app
import (
	"github.com/go-gorp/gorp"
	"log"
	"github.com/gin-gonic/gin"
	"golang-gin-api/db"
)

type App struct {
	DbMap *gorp.DbMap
}

func (app *App) Init() {
	app.DbMap = db.InitDb()
}

func (app *App) AppendEnv(c *gin.Context) {
	log.Print("AppendEnv is executed each request.")
	c.Set("dbMap", app.DbMap)
}
