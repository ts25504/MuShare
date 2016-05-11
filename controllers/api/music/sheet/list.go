package sheet
import (
	"github.com/jinzhu/gorm"
	"net/http"
	"MuShare/manager/music/sheet"
	"MuShare/controllers/api/user/friend"
	"MuShare/datatype/request/music"

)

func ListSheet(db *gorm.DB, body *music.Sheet, rw http.ResponseWriter) {
  sheet := sheet.Sheet{DB:db}
  res := sheet.ListSheet(body)
  friend.Response(res, rw)
}
