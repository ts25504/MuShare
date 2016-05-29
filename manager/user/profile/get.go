package profile

import (
  "MuShare/datatype"
  "MuShare/datatype/request/user"
  "MuShare/db/models"
)

func (this *Profile) GetProfile(body *user.Profile) datatype.Response {
  tx := this.DB.Begin()
  user := models.User{}

  if body.FriendID == nil || body.FriendID == 0 {
     badRequest("")
  }

  if err := tx.Find(&user, body.FriendID); err != nil {
    panic(err.Error)
  }

  if tx.NewRecord(&user) {
    forbidden("User Doesn't Exsit")
  }

  tx.Commit()

  return ok("", user)
}
