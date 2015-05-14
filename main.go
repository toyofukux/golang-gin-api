package main
import (
	"golang-gin-api/app"
	"golang-gin-api/routes"
)

func main() {
	app := &app.App{}
	app.Init()

	router := routes.Init(app)

	router.Run(":8080")
}
