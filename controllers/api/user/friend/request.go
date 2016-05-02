package friend

import (
  "net/http"
  "github.com/go-martini/martini"
  "github.com/jinzhu/gorm"
  "MuShare/datatype/request/user"
)

func GetRequests(db *gorm.DB, c martini.Context, body *user.Friend,
  rw http.ResponseWriter) {

}

func NewRequest(db *gorm.DB, c martini.Context, body *user.Friend,
  rw http.ResponseWriter) {

}

func AcceptRequest(db *gorm.DB, c martini.Context, body *user.Friend,
  rw http.ResponseWriter) {

}

