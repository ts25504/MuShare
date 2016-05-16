package account

import (
  "MuShare/db/models"
  "MuShare/datatype"
  "MuShare/utils"
  "MuShare/datatype/request/user"
  "time"
  "strconv"
)

func (this *Account) Login(body *user.Account) datatype.Response{
  user := models.User{}
  tx := this.DB.Begin()
  if body.Password == "" {
    return badRequest("")
  }

  if body.Mail == "" {
    return badRequest("")
  }

  tx.Where("mail=?", body.Mail).First(&user)

  if !checkPassword(user, body.Password) {
    return forbidden("Login Failed")
  }

  tx.Model(&user).UpdateColumns(models.User{LastLoginAt:time.Now().Unix()})
  tx.Commit()

  user.Token = utils.TokenEncode(strconv.Itoa(user.ID))
  return ok("", user)
}

func checkPassword(user models.User, password string) bool {
  if user.ID == 0 || password != user.Password {
    return false
  }
  return true
}
