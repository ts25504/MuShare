package sheet

import (
	"MuShare/datatype/request/music"
	"MuShare/datatype"
	"MuShare/db/models"
	"net/http"
	"time"
	"strconv"
)

func (this *Sheet) Create(body *music.Sheet) datatype.Response{
	var res datatype.Response
	sheet := models.Sheet{}
	tx := this.DB.Begin()
	u := models.User{}

	if body.RequestFromID == 0 {
    goto BadRequest
  }

	tx.Where("id = ?", strconv.Itoa(body.RequestFromID)).First(&u)

	if u.ID == 0{
		goto Forbidden
	}
	//One user cannot create sheet with same name
	tx.Where("user_id = ? AND name = ?",strconv.Itoa(u.ID),body.Name).First(&sheet)
	if sheet.ID != 0 {
		goto Forbidden
	}

	CreateSheet(body, &sheet)
	tx.Create(&sheet)
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
		ResponseText: "not available user or same sheet name for the user",
  }
	return res
}

func CreateSheet(body *music.Sheet, sheet *models.Sheet) {
	sheet.Name = body.Name
	sheet.Privilege = body.Privilege
	sheet.UserID = body.RequestFromID
	sheet.CreatedAt = time.Now().Unix()
	sheet.UpdatedAt = time.Now().Unix()
}