package middlewares

import (
	config "MuShare/conf"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
)

func initDB(conf *config.Conf) *gorm.DB {
	db, err := gorm.Open("mysql", conf.Mysql.User + ":" +
	conf.Mysql.Password + "@/" + conf.Mysql.Database)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func InjectDB(conf *config.Conf) martini.Handler {
	return func(c martini.Context) {
		db := initDB(conf)
		c.Map(db)
	}
}
