package friend

import (
  "MuShare/datatype/request/user"
  "net/http"
  "MuShare/datatype"
  "MuShare/db/models"
)

func (this *Friend) UnFollow(body *user.Friend) datatype.Response {
  var res datatype.Response
  friend := models.Friends{}
  tx := this.DB.Begin()

  if body.FromID == 0 || body.ToID == 0 {
    goto BadRequest
  }

  tx.Where("from_id=? AND to_id=?", body.FromID, body.ToID).First(&friend)

  if tx.NewRecord(&friend) {
    goto Forbidden
  }

  if !tx.NewRecord(&friend) && friend.State == stateRequest {
    goto Forbidden
  }

  tx.Delete(&friend)
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
