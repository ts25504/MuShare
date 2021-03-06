package profile

import (
  "MuShare/datatype"
  "MuShare/datatype/request/user"
  "MuShare/db/models"
)

func (this *Profile) UpdateProfile(body *user.Profile) datatype.Response {
  tx := this.DB.Begin()
  user := models.User{}
  user.ID = body.UserID
  update := make(map[string]interface{})


  if body.Name != nil {
    update["name"] = body.Name
  }

  if body.Avatar != nil {
    if gender[body.Avatar.(string)] == "" {
      update["avatar"] = gender[body.Avatar.(string)]
    } else {
      update["avatar"] = gender["Male"]
    }
  }

  if body.Birth != nil {
    update["birth"] = body.Birth
  }

  if body.Gender != nil {
    update["gender"] = body.Gender
  }

  if body.Phone != nil {
    update["phone"] = body.Phone
  }

  if body.Description != nil {
    update["description"] = body.Description
  }

  if body.Name == "" {
    badRequest("Name Can't Be Null")
  }

  err := tx.Model(&user).Updates(update).Error

  if err != nil {
    panic(err.Error())
  }

  tx.Commit()

  return ok("", nil)
}
