package main

import (
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/sessions"
  config "MuShare/conf"
  "MuShare/router"
  "MuShare/middlewares"
)

func main() {
  conf := config.LoadConf(martini.Env)
  m := martini.Classic()
  //create new session middleware
  store := sessions.NewCookieStore([]byte("MushareSecret"))
  store.Options(sessions.Options{
    Path: "/",
    Domain: conf.App.Host,
    MaxAge: 60 * 60 * 60 * 24,
    HttpOnly: true,
  })

  //middleware
  m.Handlers(
    middlewares.SetLogOutput,
    martini.Recovery(),
    martini.Logger(),
    sessions.Sessions("_session", store),
    martini.Static("static", martini.StaticOptions{}),
    middlewares.InjectRedis(conf.Redis),
    middlewares.InjectDB(conf.Mysql),
  )

  //routers
  router.Include(m)
  //start server
  m.RunOnAddr(conf.App.Host + ":" + conf.App.Port)
}
