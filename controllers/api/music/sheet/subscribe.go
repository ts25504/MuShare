package sheet

import (
  "github.com/jinzhu/gorm"
  "MuShare/datatype/request/music"
  "net/http"
  "MuShare/manager/music/sheet"
  "MuShare/controllers/api/user/friend"
)

func Subscribe(db *gorm.DB, body *music.Sheet, rw http.ResponseWriter) {
  sheet := sheet.Sheet{DB:db}
  res := sheet.Subscribe(body)
  friend.Response(res, rw)
}
