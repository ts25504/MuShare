package middlewares

import (
	"MuShare/conf"
	"github.com/go-martini/martini"
	"MuShare/db/models"
)

func InjectDB() martini.Handler {
	return func(c martini.Context, config *conf.Conf) {
		db := models.New(config.Mysql)
		db.Callback().Update().Remove("gorm:update_time_stamp")
		db.Callback().Create().Remove("gorm:update_time_stamp")
		c.Map(db)
		c.Next()
		db.Close()
	}
}
