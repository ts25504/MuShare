package middlewares

import (
	config "MuShare/conf"
	"github.com/go-martini/martini"
	"MuShare/db/models"
)

func InjectDB(mysql config.Mysql) martini.Handler {
	return func(c martini.Context) {
		db := models.New(mysql)
		c.Map(db)
		c.Next()
		db.Close()
	}
}
