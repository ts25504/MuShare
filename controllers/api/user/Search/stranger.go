package Search

import (
  "encoding/json"
  "net/http"
  "MuShare/datatype"
  "github.com/jinzhu/gorm"
  "MuShare/datatype/request/user"
  . "MuShare/manager/user/search"
)

func Stranger(db *gorm.DB, body *user.Search, rw http.ResponseWriter){
  search := Search{DB:db}
  res := search.Search(body)
  Response(res, rw)
}

func Response(res datatype.Response, rw http.ResponseWriter){
  resJson, err := json.Marshal(res)
  if err != nil {
    panic(err.Error())
  }
  //send json response
  rw.Header().Set("content-Type", "application/json; charset=utf-8")
  rw.WriteHeader(res.Status)
  rw.Write(resJson)
}
