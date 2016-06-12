package profile

import (
  "net/http"
  "github.com/jinzhu/gorm"
  "MuShare/datatype/request/user"
  . "MuShare/manager/user/profile"
  "fmt"
)

func GetProfile(db *gorm.DB, body *user.Profile, rw http.ResponseWriter){
  profile := Profile{DB:db}
  res := profile.GetProfile(body)
  fmt.Print(res)
  Response(res, rw)
}
