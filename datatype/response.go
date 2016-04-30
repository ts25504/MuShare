package datatype

type Response struct {
  Status       int         `json:"status"`
  ResponseText string      `json:"responseText"`
  Body         interface{} `json:"body,omitempty"`
}
