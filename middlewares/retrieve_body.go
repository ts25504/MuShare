package middlewares

import (
  "net/http"
  "github.com/go-martini/martini"
  "MuShare/utils"
  "reflect"
  "log"
)

func RetrieveBody(typ reflect.Type) martini.Handler {
  return func (req *http.Request, c martini.Context, logger *log.Logger) {
    body := reflect.New(typ).Interface()
    c.Map(body)
    c.Map(reflect.TypeOf(body))
    err := utils.JsonDecoder(req.Body, body)
    if err != nil {

    }
    logger.Println(body)
  }
}
