package webdriver

import (
  "bytes"
  "encoding/json"
  // "fmt"
  // "net/http"
)

type (

  // the standard Json returned from a server
  WireResponse struct {
    Name                  string `json:"name"`
    Session                       *Session
    SessionID             string `json:"sessionId"`
    Status                   int `json:"status"`
    Value        json.RawMessage `json:"value"`

  }

)

// Convenience method to extract a WireResponse.Value as a string.
func (s *WireResponse) StringValue() (value string) {

  if s.Value != nil {
    value = string(bytes.Trim(s.Value, "{}\""))
  }

  return value
}

// Convenience method to unmarshal the json.RawMessage Value to a string.
func (s *WireResponse) UnmarshalValue() (value string, err error) {

  if s.Value != nil {
    err = json.Unmarshal(s.Value, &value)
  }

  return value, err
}

// Convenience method to unmarshal the json.RawMessage Value to a string.
func (s *WireResponse) WebElement() (value *WebElement, err error) {

  if s.Value != nil {
    err = json.Unmarshal(s.Value, &value)
    value.Session = s.Session
  }

  return value, err
}

// Convenience method to unmarshal the json.RawMessage Value to a string.
func (s *WireResponse) WebElements() (value []*WebElement, err error) {

  if s.Value != nil {
    err = json.Unmarshal(s.Value, &value)
    for _, v := range value {
      v.Session = s.Session
    }
  }

  return value, err
}























