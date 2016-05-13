package audio

import (
	"MuShare/datatype/request/music"
	"MuShare/datatype"
	"MuShare/db/models"
	"net/http"
	"time"
)

func (this *Audio) AddAudio(body *music.Audio) datatype.Response{
	var res datatype.Response
	audio := models.Audio{}
	tx := this.DB.Begin()
	sheet := models.Sheet{}

	if body.SheetID == 0{
		goto BadRequest
	}
	if body.AudioUrl == ""{
		goto BadRequest
	}

	tx.Where("audio_url = ?",body.AudioUrl).First(&audio)
	if audio.ID != 0{
		goto Forbidden
	}

	tx.Where("id = ?", body.SheetID).First(&sheet)
	if sheet.ID == 0{
		goto Forbidden
	}

	if sheet.UserID != body.RequestFromID{
		goto Forbidden
	}

	CreateMusic(body, &audio)
	tx.Create(&audio)
	tx.Commit()
	res = datatype.Response{
    Status: http.StatusOK,
		ResponseText: "success",
  }
	return res

	BadRequest:
	res = datatype.Response{
    Status: http.StatusBadRequest,
  }
  return res

	Forbidden:
	res = datatype.Response{
    Status: http.StatusForbidden,
		ResponseText: "not available sheet or existed audio url",
  }
	return res

}

func CreateMusic(body *music.Audio, audio *models.Audio){
	audio.Name = body.Name
	audio.SheetID = body.SheetID
	audio.AudioUrl = body.AudioUrl
	audio.ImageUrl = body.ImageUrl
	audio.CreatedAt = time.Now().Unix()
	audio.UpdatedAt = time.Now().Unix()

}