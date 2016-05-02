package friend

import (
  "net/http"
  "github.com/go-martini/martini"
  "github.com/jinzhu/gorm"
  "MuShare/datatype/request/user"
)

func GetFriendsList(db *gorm.DB, c martini.Context, body *user.Friend,
  rw http.ResponseWriter){

}
