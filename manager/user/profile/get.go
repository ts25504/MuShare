package profile

import (
  "MuShare/datatype"
  "MuShare/datatype/request/user"
  "MuShare/db/models"
)

func (this *Profile) GetProfile(body *user.Profile) datatype.Response {
  tx := this.DB.Begin()
  user := models.User{}
  tx.Find(user, body.UserID)
  return ok("", user)
}
