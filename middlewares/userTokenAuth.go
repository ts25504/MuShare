package middlewares

import (
  "gopkg.in/redis.v3"
  "MuShare/db/models"
  "strconv"
  "net/http"
  "encoding/json"
  "MuShare/datatype"
)

func UserTokenAuth(redis redis.Client, user models.User, rw http.ResponseWriter) {
  token, err := redis.HGet("login", strconv.Itoa(user.ID)).Result()
  if err != nil {
    panic("Token Check Failed")
  }

  if(token != user.Token) {
    rw.Header().Set("Content-Type", "application/json;charset=utf-8")
    rw.WriteHeader(http.StatusForbidden)
    res, _ :=json.Marshal(datatype.Response{
      Status: http.StatusForbidden,
      ResponseText: "Token Auth Failed",
    })
    rw.Write(res)
  }
}
