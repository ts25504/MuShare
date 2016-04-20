package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "MuShare/models/migration"
	config "MuShare/conf"
	"MuShare/router"
	"MuShare/middlewares"
)

func main() {
	conf := config.LoadConf()
	m := martini.Classic()
	var host string
	var port string
	if martini.Env == "Prod" {
		host = conf.Prod.Host
		port = conf.Prod.Port
	}else {
		host = conf.Dev.Host
		port = conf.Dev.Port
	}

	//create new session middleware
	store := sessions.NewCookieStore([]byte("secret123"))
	store.Options(sessions.Options{
		Path: "/",
		Domain: host,
		MaxAge: 60 * 60 * 60 * 24,
		HttpOnly: true,
	})
	//middleware
	m.Handlers(
		sessions.Sessions("_session", store),
		martini.Static("static", martini.StaticOptions{}),
		middlewares.InjectDB(conf),
	)

	//routers
	router.Include(m)
	//start server
	m.RunOnAddr(host + ":" + port)
}
