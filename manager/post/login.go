package post

import (
  "MuShare/db/models"
  "MuShare/datatype"
  "MuShare/utils"
  "net/http"
)

func (this *Post) Login(body *datatype.LoginBody) datatype.Response{
  var res datatype.Response
  u := models.User{}
  if (body.Password == "") {
    goto BadRequest
  }

  if body.Name == "" && body.Mail == "" && body.Phone == "" {
    goto BadRequest
  }

  if body.Mail != "" {
    this.DB.Where("mail=?", body.Mail).First(&u)
  } else if body.Phone != "" {
    this.DB.Where("phone=?", body.Phone).First(&u)
  }else {
    this.DB.Where("name=?", body.Name).First(&u)
  }
  if checkPassword(u, body.Password) {
    goto Ok
  } else {
    goto Forbidden
  }

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

  Ok:
  u.Token = utils.RandomTaken()
  res = datatype.Response{
    Status: http.StatusOK,
    Body: u,
  }
  return res
}

func checkPassword(user models.User, password string) bool {
  if user.ID == 0 || password != user.Password {
    return false
  }
  return true
}
