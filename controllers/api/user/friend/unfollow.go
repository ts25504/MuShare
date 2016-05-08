package friend

import (
  "net/http"
  . "MuShare/manager/user/friend"
  "github.com/jinzhu/gorm"
  "MuShare/datatype/request/user"
)

func UnFollow(db *gorm.DB, body *user.Friend, rw http.ResponseWriter){
  friend := Friend{DB:db}
  res := friend.UnFollow(body)
  Response(res, rw)
}