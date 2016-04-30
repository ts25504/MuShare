package api

import (
  "net/http"
  "github.com/jinzhu/gorm"
  . "MuShare/manager/post"
  "encoding/json"
  "github.com/go-martini/martini"
	"MuShare/datatype"
)

func Register(db *gorm.DB, c martini.Context, body *datatype.RegisterBody, rw http.ResponseWriter) {
	if (db == nil) {
    panic("db is not exist")
  }
  post := Post{DB:db}
  res := post.Register(body)
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


