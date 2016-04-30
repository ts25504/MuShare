package utils

import (
  "io"
  "encoding/json"
  "crypto/rand"
  "encoding/base64"
)

func JsonDecoder(r io.Reader, v interface{}) error {
  decoder := json.NewDecoder(r)
  err := decoder.Decode(&v)
  if err != nil {
    return err
  }
  return nil
}

func RandomTaken() string {
  b := make([]byte, 10)
  rand.Read(b)
  return base64.URLEncoding.EncodeToString(b)
}
