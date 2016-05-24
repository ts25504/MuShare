package account

import (
  "net/http"
  "github.com/jinzhu/gorm"
  . "MuShare/manager/user/account"
  "github.com/go-martini/martini"
  "MuShare/datatype/request/user"
)

func Register(db *gorm.DB, c martini.Context, body *user.Account, rw http.ResponseWriter) {
	if (db == nil) {
    panic("db is not exist")
  }
  account := Account{DB:db}
  res := account.Register(body)
  if res.Status == http.StatusOK {
    c.Map(res.Body)
    c.Next()
  }

  Response(res, rw)
}


