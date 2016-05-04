package router

import (
  "github.com/go-martini/martini"
  "MuShare/controllers/pages"
  . "MuShare/middlewares"
  "MuShare/controllers/api/user/account"
  "MuShare/controllers/api/user/friend"
  "MuShare/datatype/request/user"
  "reflect"
)

func Include(m *martini.ClassicMartini) {
  includePages(m)
  includeUserApi(m)
}

func includePages(m *martini.ClassicMartini) {
  m.Get("/", pages.Index)
  m.Get("/test", pages.TestPage)
}

func includeUserApi(m *martini.ClassicMartini) {
  m.Group("/user/account", func(r martini.Router) {
    r.Post("/login", account.Login, account.LoginSetToken)
    r.Post("/register", account.Register)
  }, RetrieveBody(reflect.TypeOf(user.Account{})))

  m.Group("/user/friend", func(r martini.Router) {
    r.Get("/list", friend.GetFriendsList)
    r.Get("/request", friend.GetRequests)
    r.Post("/request", friend.NewRequest)
    r.Put("/request", friend.AcceptRequest)
    r.Delete("/delete", friend.UnFollow)
  }, RetrieveBody(reflect.TypeOf(user.Friend{})), friend.TokenAuth)

  m.Group("/user/profile", func(r martini.Router) {
    r.Get("/:id")
    r.Put("update")
  })
}

