package middlewares

import (
  "net/http"
  "github.com/go-martini/martini"
  "MuShare/utils"
)

func RetrieveBody(body interface{}) martini.Handler {
  return func (req *http.Request, c martini.Context) {
    c.Map(body)
    err := utils.JsonDecoder(req.Body, body)
    if err != nil {

    }
  }
}
