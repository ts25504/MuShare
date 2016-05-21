package sheet

import (
	"MuShare/datatype/request/music"
	"MuShare/datatype"
	"MuShare/db/models"
	"strconv"
	"net/http"
)

func (this *Sheet) DeleteSheet(body *music.Sheet) datatype.Response{
	var res datatype.Response
	sheet := models.Sheet{}
	tx := this.DB.Begin()

	if body.RequestFromID== 0{
		return badRequest("")
	}

	//get id of sheet
	tx.Where("name = ? AND user_id = ?",
		body.Name, strconv.Itoa(body.RequestFromID)).First(&sheet)
	//find the releted audios
	tx.Model(&sheet).Related(&sheet.Audios)
	for _, audio := range sheet.Audios{
		tx.Delete(&audio)
	}

	tx.Delete(&sheet)
	tx.Commit()
  return ok("deleted",sheet)
}