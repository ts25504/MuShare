package middlewares

import (
	config "MuShare/conf"
	"github.com/go-martini/martini"
	"MuShare/db/models"
)

func InjectDB(mysql config.Mysql) martini.Handler {
	return func(c martini.Context) {
		db := models.New(mysql)
    db.Callback().Create().Remove("gorm:update_time_stamp")
    db.Callback().Update().Remove("gorm:update_time_stamp")
		c.Map(db)
		c.Next()
		db.Close()
	}
}
