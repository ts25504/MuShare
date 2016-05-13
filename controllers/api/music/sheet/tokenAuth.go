package sheet

import (
  "strconv"
  "net/http"
  "MuShare/datatype"
  "gopkg.in/redis.v3"
  "encoding/json"
	"MuShare/datatype/request/music"
)

func TokenAuth(body *music.Sheet, redis *redis.Client, rw http.ResponseWriter,
  req *http.Request) {
  var token string
  if(req.Method == http.MethodGet){
		body.RequestToID, _ = strconv.Atoi(req.URL.Query().Get("to_id"))
    body.RequestFromID, _ = strconv.Atoi(req.URL.Query().Get("from_id"))
    body.Token = req.URL.Query().Get("token")
  }

  result := redis.HGet("login", strconv.Itoa(body.RequestFromID))
  if result == nil {
    goto Forbidden
  }

  token, _ = result.Result()

  if(body.Token != token) {
    goto Forbidden
  }

  return

  Forbidden:
  rw.Header().Set("Content-Type", "application/json;charset=utf-8")
  rw.WriteHeader(http.StatusForbidden)
  res, _ :=json.Marshal(datatype.Response{
    Status: http.StatusForbidden,
    ResponseText: "Token Auth Failed",
  })
  rw.Write(res)
}