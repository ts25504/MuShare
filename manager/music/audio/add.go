package audio

import (
	"MuShare/datatype/request/music"
	"MuShare/datatype"
	"MuShare/db/models"
	"time"
)

func (this *Audio) AddAudio(body *music.Audio) datatype.Response{
	audio := models.Audio{}
	tx := this.DB.Begin()
	sheet := models.Sheet{}

	if body.SheetID == 0{
		return badRequest("")
	}
	if body.AudioUrl == ""{
		return badRequest("No such music")
	}

	tx.Where("audio_url = ?",body.AudioUrl).First(&audio)
	if audio.ID != 0{
		return forbidden("url already existed")
	}

	tx.Where("id = ?", body.SheetID).First(&sheet)
	if sheet.ID == 0{
		return forbidden("No sheet")
	}

	if sheet.UserID != body.RequestFromID{
		return forbidden("auth fail")
	}

	createMusic(body, &audio)
	tx.Create(&audio)
	tx.Commit()
  return ok("success", audio)
}

func createMusic(body *music.Audio, audio *models.Audio){
	audio.Name = body.Name
	audio.SheetID = body.SheetID
	audio.AudioUrl = body.AudioUrl
	audio.ImageUrl = body.ImageUrl
	audio.CreatedAt = time.Now().Unix()
	audio.UpdatedAt = time.Now().Unix()

}