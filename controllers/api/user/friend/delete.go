package friend

import (
  "net/http"
  . "MuShare/manager/user/friend"
  "github.com/jinzhu/gorm"
  "MuShare/datatype/request/user"
)

func Delete(db *gorm.DB, body *user.Friend, rw http.ResponseWriter){
  friend := Friend{DB:db}
  res := friend.Delete(body)
  Response(res, rw)
}