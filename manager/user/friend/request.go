package friend

import (
  "MuShare/datatype/request/user"
  "MuShare/datatype"
  "net/http"
  "MuShare/db/models"
  "time"
  "strconv"
)

const stateRequest = "request"
const stateAgree = "agree"

func (this *Friend) Get(body *user.Friend) datatype.Response {
  var res datatype.Response
  friends := []models.Friends{}
  tx := this.DB.Begin()

  if body.FromID == 0 {
    goto BadRequest
  }

  tx.Where("to_id = ? AND state = ?",
    strconv.Itoa(body.ToID), stateRequest).Find(&friends)

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

func (this *Friend) NewRequest(body *user.Friend) datatype.Response {
  var res datatype.Response
  toUser := models.User{}
  friend := models.Friends{}
  tx := this.DB.Begin()

  if body.FromID == 0 || body.ToID == 0 {
    goto BadRequest
  }

  toUser.ID = body.ToID

  tx.First(&toUser)

  if toUser.Mail == "" {
    goto Forbidden
  }


  tx.Where("from_id=? AND to_id=?", strconv.Itoa(body.FromID),
    strconv.Itoa(body.ToID)).First(&friend)

  if !tx.NewRecord(&friend) {
    goto Forbidden
  }

  friend.FromID = body.FromID
  friend.ToID = body.ToID
  friend.State = stateRequest
  friend.CreatedAt = time.Now().Unix()
  friend.UpdatedAt = time.Now().Unix()

  tx.Create(&friend)

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
    ResponseText: "Operation not accepted",
  }
  return res
}

func (this *Friend) AcceptRequest(body *user.Friend) datatype.Response {
  var res datatype.Response
  friend := models.Friends{}
  tx := this.DB.Begin()

  if (body.FromID == 0 || body.ToID == 0) {
    goto BadRequest
  }

  tx.Where("from_id=? AND to_id=?", strconv.Itoa(body.FromID),
    strconv.Itoa(body.ToID)).First(&friend)

  if tx.NewRecord(&friend) {
    goto Forbidden
  }

  if !tx.NewRecord(&friend) && friend.State != stateRequest {
    goto Forbidden
  }

  friend.State = stateAgree;
  friend.UpdatedAt = time.Now().Unix()
  tx.Save(&friend)

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
    ResponseText: "Operation not accepted",
  }
  return res
}
