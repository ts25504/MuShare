package audio

import (
	"MuShare/datatype/request/music"
	"github.com/jinzhu/gorm"
	"net/http"
	"MuShare/manager/music/migration"
	"MuShare/controllers/api/user/friend"
)

func SheetMigration (db *gorm.DB, body *music.SheetMigration,rw http.ResponseWriter){
	sheetMigra := audio.SheetMigration{DB:db}
	res := sheetMigra.Migration(body)
	friend.Response(res, rw)
}