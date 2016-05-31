package middlewares

import (
  "net/http"
  "github.com/go-martini/martini"
  "MuShare/utils"
  "reflect"
  "log"
  "strings"
  "strconv"
)

func RetrieveBody(typ reflect.Type) martini.Handler {
  return func (req *http.Request, c martini.Context, logger *log.Logger) {
    body := reflect.New(typ).Interface()
    c.Map(body)
    c.Map(reflect.TypeOf(body))

    if req.Method == http.MethodGet {
      rv := reflect.ValueOf(body).Elem()
      for key, _ := range req.URL.Query()  {
        value := req.URL.Query().Get(key)
        field := strings.Title(key)
        if rv.FieldByName(field).IsValid() && rv.FieldByName(field).CanSet() {
          logger.Println(value)
          if strings.Contains(field, "ID"){
            intValue, _ := strconv.Atoi(value)
            rv.FieldByName(field).SetInt(int64(intValue))
          }else {
            rv.FieldByName(field).SetString(value)
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
