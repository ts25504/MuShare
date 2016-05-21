package sheet

import (
	"MuShare/datatype/request/music"
	"MuShare/datatype"
	"MuShare/db/models"
	"strconv"
)

const priFriend  = "friend"
const priPublic = "public"
const stateAgree = "agree"

func (this *Sheet) ListSheet(body *music.Sheet) datatype.Response{
	sheets := []models.Sheet{}

	user := models.User{}
	friend := models.Friends{}
	tx := this.DB.Begin()

	if body.RequestFromID == 0 || body.RequestToID == 0{
		return badRequest("")
	}

	tx.Where("id = ?",
		strconv.Itoa(body.RequestFromID)).Find(&user)
	if user.ID == 0{
		return forbidden("no such request user")
	}
	tx.Where("id = ?",
		strconv.Itoa(body.RequestToID)).Find(&user)
	if user.ID == 0{
		return forbidden("no such required user")
	}

	if body.RequestFromID == body.RequestToID{
		tx.Where("user_id = ?",
				strconv.Itoa(body.RequestFromID)).Find(&sheets)

	}else {
		//detect whether friend or not
		tx.Where("from_id = ? AND to_id = ? AND state = ?",
			strconv.Itoa(body.RequestFromID), strconv.Itoa(body.RequestToID), stateAgree).First(&friend)

		if friend.ID == 0 {
			tx.Where("user_id = ? AND privilege = ?",
				strconv.Itoa(body.RequestToID), priPublic).Find(&sheets)

		}else {
			tx.Where("user_id = ? AND privilege in (?)",
				strconv.Itoa(body.RequestToID), []string{priPublic, priFriend}).Find(&sheets)
		}
	}
	tx.Commit()

	return ok("", sheets)

}

