package sheet

import (
	"MuShare/datatype/request"
	"MuShare/datatype"
	"MuShare/db/models"
	"net/http"
	"time"
)

func (this *Sheet) Create(body *request.Sheet) datatype.Response{
	var res datatype.Response
	sheet := models.Sheet{}
	tx := this.DB.Begin()
	u := models.User{}

	if body.UserID == 0 {
    goto BadRequest
  }

	tx.Where("id=?", body.UserID).First(&u)

	if u.ID == 0{
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
		ResponseText: "not available user",
  }
	return res
}

func CreateSheet(body *request.Sheet, sheet *models.Sheet) {
	sheet.Name = body.Name
	sheet.Privilege = body.Privilege
	sheet.UserID = body.UserID
	sheet.CreatedAt = time.Now().Unix()
	sheet.UpdatedAt = time.Now().Unix()
}