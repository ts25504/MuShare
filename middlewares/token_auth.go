package middlewares

import (
  "net/http"
  "MuShare/datatype"
  "encoding/json"
  "gopkg.in/redis.v3"
  "regexp"
  "MuShare/utils"
  "strings"
  "reflect"
  "github.com/go-martini/martini"
  "strconv"
)

const UserIdField = "UserID"

func TokenAuth(redis *redis.Client, c martini.Context, typ reflect.Type,
rw http.ResponseWriter, req *http.Request) {

  var err error
  var decodeToken string
  var encodeToken string
  var expectToken string
  r := regexp.MustCompile(`\s*(?P<token>.{10,})\s*`)
  group := make(map[string]string)
  match := r.FindStringSubmatch(req.Header.Get("Authorization"))
  if len(match) < 2 {
    Unauthorized("User Auth Failed", rw)
    return
  }
  for i, name := range r.SubexpNames() {
    if i != 0 {
      group[name] = match[i]
    }
  }
  encodeToken = group["token"]
  decodeToken, err = utils.TokenDecode(group["token"])
  if err != nil {
    panic("Token Decode Error")
  }

  userId := strings.Split(decodeToken, ":")[0]

  result := redis.HGet("login", userId)
  expectToken, _ = result.Result()
  if expectToken != encodeToken {
    Unauthorized("User Auth Failed", rw)
    return
  }

  setUserId(c, typ, userId)
}

func setUserId(c martini.Context, typ reflect.Type, userId string) {
  body := c.Get(typ)
  e := body.Elem()
  if e.Kind() == reflect.Struct {
    value := e.FieldByName(UserIdField)
    if value.IsValid() && value.CanSet() {
      id, _ := strconv.ParseInt(userId, 10, 64)
      value.SetInt(id)
    }
  }
}

func Unauthorized(responseText string, rw http.ResponseWriter) {
  rw.Header().Set("Content-Type", "application/json;charset=utf-8")
  rw.WriteHeader(http.StatusUnauthorized)
  res, _ := json.Marshal(datatype.Response{
    Status: http.StatusUnauthorized,
    ResponseText: "Token Auth Failed",
  })
  rw.Write(res)
}
