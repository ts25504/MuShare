package test

import (
  "fmt"
  "testing"
  "encoding/json"
)

func TestJson(t *testing.T) {
  jsonStr := `{"name": [{"first": "li", "second": "yifan"}, {"first": "li", "second": "yifan"}]}`
  obj := make(map[string]interface{})
  json.Unmarshal([]byte(jsonStr), &obj)
  fmt.Println(obj)
}

