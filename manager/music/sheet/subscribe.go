package sheet

import (
  "MuShare/datatype"
  "MuShare/datatype/request/music"
  "MuShare/db/models"
)

func (this *Sheet) Subscribe(body *music.Sheet) datatype.Response {

  sheet := models.Sheet{}
  subscribe := models.Subscribe{}
  tx := this.DB.Begin()

  if body.ToID == 0 {
    badRequest("")
  }

  if err := tx.Where("id=?", body.ToID).Find(&sheet).Error; err != nil {
    panic(err.Error())
  }

  if tx.NewRecord(&sheet) {
    forbidden("Sheet Doesn't Exist")
  }

  if sheet.UserID == body.UserID {
    forbidden("Can't Subscribe Self's Sheet")
  }

  if err := tx.Where("user_id=? AND sheet_id=?", body.UserID,
    body.ToID).Find(&subscribe).Error; err != nil {
    panic(err.Error())
  }

  if !tx.NewRecord(&subscribe) {
    forbidden("Already Subscribe")
  }

  subscribe.SheetID = body.ToID
  subscribe.UserID = body.UserID

  if err := tx.Create(&subscribe).Error; err != nil {
    panic(err.Error())
  }

  return ok("", nil)
}
