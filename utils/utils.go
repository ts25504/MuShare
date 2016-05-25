package utils

import (
  "io"
  "encoding/json"
  "crypto/rand"
  "encoding/base64"
  "golang.org/x/crypto/scrypt"
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
  id := []string{userId}
  return EncryptRandSequence(10, id...)
}

func TokenDecode(src string) (string, error){
  dst, err := base64.URLEncoding.DecodeString(src)
  return string(dst), err
}

func PsdHandler(psd string, salts []byte) (string, error){
  dk, err := scrypt.Key([]byte(psd), salts, 16384, 8, 1, 32)
  dst := base64.URLEncoding.EncodeToString(dk)
  return dst, err
}


func EncryptRandSequence(n int, userId ...string) string{
  var src []byte
  b := make([]byte, n)
  rand.Read(b)
  if userId != nil{
    src = append([]byte(userId[0] + ":"),b...)
  }else{
    src = b
  }
  dst := base64.URLEncoding.EncodeToString(src)
  return dst
}


