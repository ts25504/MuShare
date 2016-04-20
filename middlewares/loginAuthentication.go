package middlewares

import (
	"github.com/martini-contrib/sessions"
)

func LoginAuthentication(session sessions.Session) {
	login := session.Get("login")
	if login != nil && login == true {

	} else {

	}
}