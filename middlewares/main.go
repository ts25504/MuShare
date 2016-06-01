package middlewares

import (
  "reflect"
  "strconv"
  "net/http"
  "MuShare/datatype"
  "encoding/json"
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
