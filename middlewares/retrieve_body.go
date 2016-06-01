package middlewares

import (
  "net/http"
  "github.com/go-martini/martini"
  "MuShare/utils"
  "reflect"
  "log"
  "strings"
)

func RetrieveBody(typ reflect.Type) martini.Handler {
  return func(req *http.Request, c martini.Context, logger *log.Logger,
  rw http.ResponseWriter) {
    body := reflect.New(typ).Interface()
    c.Map(body)
    c.Map(reflect.TypeOf(body))

    if req.Method == http.MethodGet {
      rv := reflect.ValueOf(body).Elem()
      for key, _ := range req.URL.Query() {
        v := req.URL.Query().Get(key)
        field := strings.Title(key)
        value := rv.FieldByName(field)
        if value.IsValid() && value.CanSet() {
          switch value.Kind() {
          case reflect.String:
            setString(value, v)
            break
          case reflect.Int:
            if !setInt(value, v) {
              badRequest(key + " field parse failed", rw)
            }
            break
          }
        }
      }
    } else {
      utils.JsonDecoder(req.Body, body)
    }
    c.Next()
    logger.Println(body)
  }
}
