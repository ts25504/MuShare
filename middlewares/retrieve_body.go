package middlewares

import (
  "net/http"
  "github.com/go-martini/martini"
  "MuShare/utils"
  "reflect"
  "log"
  "strings"
  "net/url"
)

func RetrieveBody(typ reflect.Type) martini.Handler {
  return func(req *http.Request, c martini.Context, logger *log.Logger,
  rw http.ResponseWriter) {
    body := reflect.New(typ).Interface()
    c.Map(body)
    c.Map(reflect.TypeOf(body))

    if req.Method == http.MethodGet {
      rv := reflect.ValueOf(body).Elem()
      parseGetQuery(rv, req.URL.Query(), rw)
    } else {
      utils.JsonDecoder(req.Body, body)
    }
    c.Next()
    logger.Println(body)
  }
}

func parseGetQuery(rv reflect.Value, originQuery url.Values, rw http.ResponseWriter) {
  if len(originQuery) > 10 {
    badRequest("Too Many Query", rw)
  }

  solvedQuery := make(map[string]string)

  for key, _ := range originQuery {
    solvedQuery[strings.ToLower(key)] = originQuery.Get(key)
  }

  for i := 0; i < rv.NumField(); i++ {
    value := rv.Field(i)
    key := rv.Type().Field(i).Name
    vs := solvedQuery[strings.ToLower(key)]
    if value.IsValid() && value.CanSet() && vs != "" {
      switch value.Kind() {
      case reflect.String:
        setString(value, vs)
        break
      case reflect.Int:
        if !setInt(value, vs) {
          badRequest(key + " field parse failed", rw)
        }
        break
      case reflect.Interface:
        if !setInterface(value, vs) {
          badRequest(key + " field parse failed", rw)
        }
        break
      }
    }
  }
}
