package friend

import (
  "MuShare/datatype/request/user"
  "MuShare/datatype"
  "MuShare/db/models"
  "strconv"
)

func (this *Friend) List(body *user.Friend) datatype.Response{
  friends := []models.Friends{}
  tx := this.DB.Begin()

  if body.UserID == 0 {
    return badRequest("")
  }

  tx.Where("user_id = ? AND state = ?",
    strconv.Itoa(body.UserID), stateAgree).Find(&friends)

  for i, _ := range friends {
    tx.Model(&friends[i]).Related(&friends[i].Friend, "Friend")
  }

  tx.Commit()


  res := make([]models.User, 0)

  for _, item := range friends {
    res = append(res, item.Friend)
  }

  return ok("", res)
}