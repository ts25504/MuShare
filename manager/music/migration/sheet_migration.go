package audio

import (
	"net/http"
	"MuShare/datatype"
	"github.com/jinzhu/gorm"
	"MuShare/datatype/request/music"
	"MuShare/db/models"
	"fmt"
	"strconv"
)

type SheetMigration struct {
  DB *gorm.DB
}

func ok(responseText string, body interface{}) datatype.Response{
  res := datatype.Response{
    Status: http.StatusOK,
    ResponseText: responseText,
    Body: body,
  }

  return res
}

func forbidden(responseText string) datatype.Response{
  res := datatype.Response{
    Status:http.StatusForbidden,
    ResponseText: responseText,
  }
  return res
}

func (this *SheetMigration) Migration (body *music.SheetMigration) datatype.Response{
	sheet := models.Sheet{}
	tx := this.DB.Begin()
	for _, id := range body.IdList{
		audio := models.Audio{}
		tx.Where("id = ?", id).First(&audio)
		tx.Where("id = ?",strconv.Itoa(body.ToSheetId)).First(&sheet)
		if sheet.ID == 0 || sheet.UserID != body.UserID{
			return forbidden("terminal sheet doesn't exist or not belong to user!")

		}
		// if the second select does not find , it will stay in origin, so need to reset sheetid to 1
		sheet.ID = 0
		tx.Where("id = ?",strconv.Itoa(audio.SheetID)).First(&sheet)
		fmt.Println(sheet.ID)
		if sheet.ID == 0 || sheet.UserID != body.UserID {
			return forbidden("the audio not belong to request user!")
		}
		audio.SheetID = body.ToSheetId
		tx.Save(&audio)
	}
	tx.Commit()
	return ok("success",  "")
}