package test

import (
  "github.com/bitly/go-simplejson"
  "MuShare/db/models"
  "net/http"
  "fmt"
  "testing"
)

func TestJson(t *testing.T) {
  user := models.User{}
  user.ID = 1
  user.Name = "liyifan"
  user.Password = "w19920610"
  user.Friends = []models.Friends{{ID:1, FriendID:2}}
  user.Friends[0].User = models.User{Name:"shuaihua"}
  user.Friends[0].User.ID = 2
  json := simplejson.New()
  json.Set("status", http.StatusOK)
  json.Set("responseText", "hahahahah")
  friends := append([]models.User{}, user.Friends[0].User)
  json.Set("body", friends)
  res, _ := json.MarshalJSON()
  fmt.Println(string(res))
}

