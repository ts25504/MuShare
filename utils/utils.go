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

func TokenEncode(userId string) string{
  b := make([]byte, 10)
  rand.Read(b)
  src := append([]byte(userId + ":"),b...)
  dst := base64.URLEncoding.EncodeToString(src)
  return dst
}

func TokenDecode(src string) (string, error){
  dst, err := base64.URLEncoding.DecodeString(src)
  return string(dst), err
}
