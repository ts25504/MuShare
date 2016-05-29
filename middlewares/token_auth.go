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
  "MuShare/conf"
)

const UserIdField = "UserID"

func TokenAuth(redis *redis.Client, c martini.Context, typ reflect.Type,
rw http.ResponseWriter, req *http.Request, config *conf.Conf) {

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
    panic(err.Error())
  }

  userId := strings.Split(decodeToken, ":")[0]

  hSetKey := config.Redis.Prefix + "_token"
  mapKey := "user_" + userId
  result := redis.HGet(hSetKey, mapKey)
  expectToken, _ = result.Result()
  if expectToken != encodeToken {
    Unauthorized("User Auth Failed", rw)
    return
  }

  if !setUserId(c, typ, userId) {
    Unauthorized("User Auth Failed", rw)
  }
}

func setUserId(c martini.Context, typ reflect.Type, userId string) bool {
  body := c.Get(typ)
  e := body.Elem()
  if e.Kind() == reflect.Struct {
    value := e.FieldByName(UserIdField)
    if value.IsValid() && value.CanSet() {
      id, err := strconv.ParseInt(userId, 10, 64)
      if err != nil {
        return false
      }
      value.SetInt(id)
    }
  }
  return true
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
