package sheet

import (
	"MuShare/datatype/request/music"
	"MuShare/datatype"
	"MuShare/db/models"
	"net/http"
	"strconv"
)

func (this *Sheet) Update(body *music.Sheet) datatype.Response{
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

	tx.Where("user_id = ? AND name = ?",strconv.Itoa(u.ID),body.UpdateName).First(&sheet)
	if sheet.ID != 0 {
		goto Forbidden
	}
	//One user cannot create sheet with same name
	tx.Where("user_id = ? AND name = ?",strconv.Itoa(u.ID),body.Name).First(&sheet)

	if sheet.ID == 0 {
		goto Forbidden
	}
	if body.UpdateName != ""{
		sheet.Name = body.UpdateName
	}
	if body.Privilege != ""{
		sheet.Privilege = body.Privilege
	}
	tx.Save(&sheet)
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
		ResponseText: "no such sheet or same sheet name for the user",
  }
	return res
}
