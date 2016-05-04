package friend

import (
  "MuShare/datatype/request/user"
  "MuShare/datatype"
  "net/http"
  "MuShare/db/models"
  "strconv"
)

func (this *Friend) List(body *user.Friend) datatype.Response{
  var res datatype.Response
  friends := []models.Friends{}
  tx := this.DB.Begin()

  if body.FromID == 0{
    goto BadRequest
  }

  tx.Where("from_id = ? AND state = ?",
    strconv.Itoa(body.FromID), stateAgree).Find(&friends)

  for i, _ := range friends {
    tx.Model(&friends[i]).Related(&friends[i].User, "User")
  }

  tx.Commit()
  res.Status = http.StatusOK
  res.Body = friends
  return res

  BadRequest:
  res = datatype.Response{
    Status: http.StatusBadRequest,
  }
  return res
}