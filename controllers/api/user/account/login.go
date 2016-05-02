package account

import (
  "net/http"
  "github.com/jinzhu/gorm"
  "encoding/json"
  "github.com/go-martini/martini"
  "gopkg.in/redis.v3"
  "MuShare/db/models"
  "strconv"
  . "MuShare/manager/user/account"
  "MuShare/datatype/request/user"
)

func Login(db *gorm.DB, c martini.Context, body *user.Account, rw http.ResponseWriter) {
  if (db == nil) {
    panic("db is not exist")
  }

  account := Account{DB:db}
  res := account.Login(body)

  if res.Status == http.StatusOK {
    c.Map(res.Body)
    c.Next()
  }

  resJson, err := json.Marshal(res)
  if err != nil {
    panic(err.Error())
  }
  //send json response
  rw.Header().Set("content-Type", "application/json; charset=utf-8")
  rw.WriteHeader(res.Status)
  rw.Write(resJson)
}

func LoginSetToken(redis *redis.Client, user models.User) {
  redis.HSet("login", strconv.Itoa(user.ID), string(user.Token)).Result()
}
