package account

import (
  "MuShare/db/models"
  "MuShare/datatype"
  "MuShare/utils"
  "net/http"
  "MuShare/datatype/request/user"
  "time"
)

func (this *Account) Login(body *user.Account) datatype.Response{
  var res datatype.Response
  user := models.User{}
  tx := this.DB.Begin()
  if body.Password == "" {
    goto BadRequest
  }

  if body.Name == "" && body.Mail == "" && body.Phone == "" {
    goto BadRequest
  }

  if body.Mail != "" {
    tx.Where("mail=?", body.Mail).First(&user)
  } else if body.Phone != "" {
    tx.Where("phone=?", body.Phone).First(&user)
  }else {
    tx.Where("name=?", body.Name).First(&user)
  }

  if !checkPassword(user, body.Password) {
    goto Forbidden
  }

  tx.Model(&user).UpdateColumns(models.User{LastLoginAt:time.Now().Unix()})
  tx.Commit()

  user.Token = utils.RandomTaken()
  res = datatype.Response{
    Status: http.StatusOK,
    Body: user,
  }

  return res


  BadRequest:
  res = datatype.Response{
    Status: http.StatusBadRequest,
  }
  return res

  Forbidden:
  res = datatype.Response{
    Status:http.StatusForbidden,
    ResponseText: "Login Failed",
  }
  return res
}

func checkPassword(user models.User, password string) bool {
  if user.ID == 0 || password != user.Password {
    return false
  }
  return true
}
