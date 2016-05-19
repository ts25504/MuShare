package middlewares

import (
	"github.com/martini-contrib/sessions"
)

func UserSessionAuth(session sessions.Session) {
	login := session.Get("login")
	if login != nil && login == true {

	} else {

	}
}
