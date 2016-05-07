package middlewares

import (
  "net/http"
  "github.com/go-martini/martini"
  "MuShare/utils"
  "reflect"
  "fmt"
  "log"
)

func RetrieveBody(typ reflect.Type) martini.Handler {
  return func (req *http.Request, c martini.Context, logger log.Logger) {
    body := reflect.New(typ).Interface()
    logger.Println(body)
    c.Map(body)
    err := utils.JsonDecoder(req.Body, body)
    if err != nil {

    }
  }
}
