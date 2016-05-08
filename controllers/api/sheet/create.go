package sheet

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"MuShare/manager/sheet"
	"MuShare/controllers/api/user/friend"
	"MuShare/datatype/request"
)

func Create(db *gorm.DB, body *request.Sheet, rw http.ResponseWriter) {
  sheet := sheet.Sheet{DB:db}
  res := sheet.Create(body)
  friend.Response(res, rw)
}