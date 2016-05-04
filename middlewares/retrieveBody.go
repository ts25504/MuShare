package middlewares

import (
  "net/http"
  "github.com/go-martini/martini"
  "MuShare/utils"
  "reflect"
)

func RetrieveBody(typ reflect.Type) martini.Handler {
  return func (req *http.Request, c martini.Context) {
    body := reflect.New(typ).Interface()
    c.Map(body)
    err := utils.JsonDecoder(req.Body, body)
    if err != nil {

    }
  }
}
