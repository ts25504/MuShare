package sheet

import (
	"MuShare/datatype/request/music"
	"MuShare/datatype"
	"MuShare/db/models"
	"strconv"
)

func (this *Sheet) Update(body *music.Sheet) datatype.Response{
	sheet := models.Sheet{}
	tx := this.DB.Begin()
	u := models.User{}

	if body.RequestFromID == 0 {
    return badRequest("")
  }

	tx.Where("id = ?", strconv.Itoa(body.RequestFromID)).First(&u)

	if u.ID == 0{
		return forbidden("no such user")
	}

	tx.Where("user_id = ? AND name = ?",strconv.Itoa(u.ID),body.UpdateName).First(&sheet)
	if sheet.ID != 0 {
		return forbidden("no such sheet")
	}
	//One user cannot create sheet with same name
	tx.Where("user_id = ? AND name = ?",strconv.Itoa(u.ID),body.Name).First(&sheet)

	if sheet.ID == 0 {
		return forbidden("not this user's sheet")
	}
	if body.UpdateName != ""{
		sheet.Name = body.UpdateName
	}
	if body.Privilege != ""{
		sheet.Privilege = body.Privilege
	}
	tx.Save(&sheet)
	tx.Commit()
	return ok("updated", sheet)
}
