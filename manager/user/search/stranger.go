package search

import (
  "MuShare/datatype/request/user"
  "MuShare/datatype"
  "regexp"
  "MuShare/db/models"
  "strings"
)

func (this *Search)Search(body *user.Search) datatype.Response {
  tx := this.DB.Begin()
  result := []models.User{}

  mailReg := regexp.MustCompile(`^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`)

  if strings.TrimSpace(body.Keyword) == "" {
    return badRequest("Keyword Can't Be Empty")
  }

  if mailReg.FindAllString(body.Keyword, -1) != nil {
    if err := tx.Where("mail like ?", "%" + body.Keyword + "%").Find(&result).Error; err != nil {
      panic(err.Error())
    }
  } else {
    if err := tx.Where("name like ?", "%" + body.Keyword + "%").Find(&result).Error; err != nil {
      panic(err.Error())
    }
  }

  tx.Commit()
  return ok("", result)
}
