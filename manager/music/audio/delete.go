package audio
import (
	"MuShare/datatype/request/music"
	"MuShare/datatype"
	"MuShare/db/models"
	"net/http"
)

func (this *Audio) DeleteAudio(body *music.Audio) datatype.Response {
	var res datatype.Response
	audio := models.Audio{}
	tx := this.DB.Begin()

	tx.Where("audio_url = ?", body.AudioUrl).First(&audio)

	if audio.ID == 0 {
		goto BadRequest
	}

	tx.Delete(&audio)
	tx.Commit()

	res = datatype.Response{
		Status:http.StatusOK,
	}
	return res

	BadRequest:
	res = datatype.Response{
		Status:http.StatusBadRequest,
	}
	return res
}