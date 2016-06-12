package middlewares

import (
  "reflect"
  "strconv"
  "net/http"
  "MuShare/datatype"
  "encoding/json"
  "regexp"
)

func setInt(rv reflect.Value, value string) bool {

  intValue, err := strconv.ParseInt(value, 10, 64)

  if err != nil {
    return false
  }

  rv.SetInt(intValue)

  return true

}

func setString(rv reflect.Value, value string) {
  rv.SetString(value)
}

func setInterface(rv reflect.Value, value string) bool{

  reg := regexp.MustCompile(`(?i).*id.*`)

  if reg.FindAllString(rv.Type().Name(), -1) != nil {
    intValue, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
      return false
    }
    rv.Set(reflect.ValueOf(intValue))
  } else {
    rv.Set(reflect.ValueOf(value))
  }

  return true
}

func unauthorized(responseText string, rw http.ResponseWriter) {
  rw.Header().Set("Content-Type", "application/json;charset=utf-8")
  rw.WriteHeader(http.StatusUnauthorized)
  res, _ := json.Marshal(datatype.Response{
    Status: http.StatusUnauthorized,
    ResponseText: responseText,
  })
  rw.Write(res)
}

func badRequest(responseText string, rw http.ResponseWriter) {
  rw.Header().Set("Content-Type", "application/json;charset=utf-8")
  rw.WriteHeader(http.StatusBadRequest)
  res, _ := json.Marshal(datatype.Response{
    Status: http.StatusUnauthorized,
    ResponseText: responseText,
  })
  rw.Write(res)
}
