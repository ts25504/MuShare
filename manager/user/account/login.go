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
  var decodeSolt string
  user := models.User{}
  salt := models.Salts{}

  tx := this.DB.Begin()
  if body.Password == "" {
    return badRequest("")
  }

  if body.Mail == "" {
    return badRequest("")
  }

  tx.Where("mail=?", body.Mail).First(&user)
  tx.Where("user_id=?", user.ID).First(&salt)
  decodeSolt, _ = utils.TokenDecode(salt.Salt)
  checkPsd, _ := utils.PsdHandler(body.Password, []byte(decodeSolt))

  if !checkPassword(user, checkPsd) {
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
