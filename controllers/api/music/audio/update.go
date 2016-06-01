package audio


import (
  "net/http"
  "github.com/jinzhu/gorm"
  "MuShare/manager/music/audio"
	"MuShare/datatype/request/music"
	"MuShare/controllers/api/user/friend"
)

func Update(db *gorm.DB, body *music.Audio, rw http.ResponseWriter){
  audio := audio.Audio{DB:db}
  res := audio.Update(body)
	friend.Response(res, rw)
}

