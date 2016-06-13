package sheet

import (
  "MuShare/datatype"
  "MuShare/datatype/request/music"
  "MuShare/db/models"
)

func (this *Sheet) Subscribe(body *music.Sheet) datatype.Response {

  sheet := models.Sheet{}
  subscribe := models.Subscribe{}
  friend := models.Friends{}
  tx := this.DB.Begin()

  if body.ToID == 0 {
    return badRequest("")
  }

  if err := tx.Where("id=?", body.ToID).Find(&sheet).Error; err != nil {
    panic(err.Error())
  }

  if tx.NewRecord(&sheet) {
    return forbidden("Sheet Doesn't Exist")
  }

  if sheet.UserID == body.UserID {
    return forbidden("Can't Subscribe Self's Sheet")
  }

  if sheet.Privilege == priPrivate {
    return forbidden("Can't Subscribe Private Sheet")
  }

  if sheet.Privilege == priFriend {
    if err := tx.Where("user_id=? AND friend_id=?", body.UserID,
      sheet.UserID).Find(&friend).Error ; err != nil{
      panic(err.Error())
    }

    if tx.NewRecord(&friend) || friend.State != stateAgree {
      return forbidden("Not Friend")
    }
  }

  tx.Where("user_id=? AND sheet_id=?", body.UserID, body.ToID).Find(&subscribe)

  if !tx.NewRecord(&subscribe) {
    return forbidden("Already Subscribe")
  }

  subscribe.SheetID = body.ToID
  subscribe.UserID = body.UserID

  if err := tx.Create(&subscribe).Error; err != nil {
    panic(err.Error())
  }

  tx.Model(&sheet).Related(&sheet.User, "User")

  tx.Commit()

  return ok("", sheet)
}
