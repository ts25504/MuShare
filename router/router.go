package router

import (
	"github.com/go-martini/martini"
	"MuShare/controllers/pages"
	"MuShare/middlewares"
)

func Include(m *martini.ClassicMartini) {
	includePages(m);
}

func includePages(m *martini.ClassicMartini) {
	m.Get("/", middlewares.LoginAuthentication, pages.Index);
}

