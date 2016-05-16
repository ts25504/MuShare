package friend

import (
  "MuShare/datatype/request/user"
  "MuShare/datatype"
  "MuShare/db/models"
)

func (this *Friend) List(body *user.Friend) datatype.Response {
  friends := []models.Friends{}
  tx := this.DB.Begin()

  if body.UserID == 0 {
    return badRequest("")
  }

  tx.Preload("Friend").Where("user_id=? AND state=?",
    body.UserID, stateAgree).Find(&friends)

  tx.Commit()

  res := make([]models.User, 0)

  for _, item := range friends {
    res = append(res, item.Friend)
  }

  return ok("", res)
}