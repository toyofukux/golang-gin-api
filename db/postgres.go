package db

import (
	_ "github.com/lib/pq" // TODO: for sql.Open
	"github.com/go-gorp/gorp"
	"database/sql"
	"golang-gin-api/model"
)

func InitDb() *gorp.DbMap {
	db, err := sql.Open("postgres", "host=db user=takasing dbname=toyo password=takapass sslmode=disable")
	if err != nil {
		// handle error
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(model.Article{}, "articles").SetKeys(true, "Id")
	dbmap.AddTableWithName(model.User{}, "users").SetKeys(true, "Id")
	return dbmap
}
