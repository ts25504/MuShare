package friend

import (
  "MuShare/datatype/request/user"
  "net/http"
  "MuShare/datatype"
  "MuShare/db/models"
)

func (this *Friend) Delete(body *user.Friend) datatype.Response {
  var res datatype.Response
  friend1 := models.Friends{}
  friend2 := models.Friends{}
  tx := this.DB.Begin()

  if body.UserID == 0 || body.FriendID == 0 {
    goto BadRequest
  }

  tx.Where("user_id=? AND friend_id=?",
    body.UserID, body.FriendID).First(&friend1)

  if tx.NewRecord(&friend1) {
    goto Forbidden
  }

  if !tx.NewRecord(&friend1) && friend1.State == stateRequest {
    goto Forbidden
  }

  tx.Delete(&friend1)

  tx.Where("user_id=? AND friend_id=?",
    body.FriendID, body.UserID).First(&friend2)

  if tx.NewRecord(&friend2) {
    goto Forbidden
  }

  if !tx.NewRecord(&friend2) && friend2.State == stateRequest {
    goto Forbidden
  }

  tx.Delete(&friend2)

  tx.Commit()

  res.Status = http.StatusOK
  return res

  BadRequest:
  res = datatype.Response{
    Status: http.StatusBadRequest,
  }
  return res

  Forbidden:
  res = datatype.Response{
    Status:http.StatusForbidden,
    ResponseText: "Operation Not Accepted",
  }
  return res
}
