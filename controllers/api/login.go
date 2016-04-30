package api

import (
  "net/http"
  "github.com/jinzhu/gorm"
  . "MuShare/manager/post"
  "encoding/json"
  "github.com/go-martini/martini"
  "gopkg.in/redis.v3"
  "MuShare/db/models"
  "strconv"
  "MuShare/datatype"
)

func Login(db *gorm.DB, c martini.Context, body *datatype.LoginBody, rw http.ResponseWriter) {
  if (db == nil) {
    panic("db is not exist")
  }
  post := Post{DB:db}
  res := post.Login(body)
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

func LoginSetToken(redis *redis.Client, user models.User) {
  redis.HSet("login", strconv.Itoa(user.ID), string(user.Token)).Result()
}
