package router

import (
  "github.com/go-martini/martini"
  "MuShare/controllers/pages"
  . "MuShare/middlewares"
  "MuShare/controllers/api"
  "MuShare/datatype"
)

func Include(m *martini.ClassicMartini) {
  includePages(m)
  includeApi(m)
}

func includePages(m *martini.ClassicMartini) {
  m.Get("/", pages.Index)
  m.Get("/test",pages.TestPage)
}

func includeApi(m *martini.ClassicMartini) {
  m.Post("/user/login",
    RetrieveBody(&datatype.LoginBody{}),
    api.Login, api.LoginSetToken)

  m.Post("/user/register",
    RetrieveBody(&datatype.RegisterBody{}),
    api.Register)
}

