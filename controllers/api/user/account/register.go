package account

import (
  "net/http"
  "github.com/jinzhu/gorm"
  . "MuShare/manager/user/account"
  "encoding/json"
  "github.com/go-martini/martini"
  "MuShare/datatype/request/user"
)

func Register(db *gorm.DB, c martini.Context, body *user.Account, rw http.ResponseWriter) {
	if (db == nil) {
    panic("db is not exist")
  }
  account := Account{DB:db}
  res := account.Register(body)
  resJson, err := json.Marshal(res)
  if err != nil {
    panic(err.Error())
  }

  if res.Status == http.StatusOK {
    c.Map(res.Body)
    c.Next()
  }
  //send json response
  rw.Header().Set("content-Type", "application/json; charset=utf-8")
  rw.WriteHeader(res.Status)
  rw.Write(resJson)
}


