package sheet

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"MuShare/manager/music/sheet"
	"MuShare/controllers/api/user/friend"
	"MuShare/datatype/request/music"

)

func Create(db *gorm.DB, body *music.Sheet, rw http.ResponseWriter) {
  sheet := sheet.Sheet{DB:db}
  res := sheet.Create(body)
  friend.Response(res, rw)
}