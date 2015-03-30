package webdriver

import (
  "bytes"
  "encoding/json"
  "errors"
  // "fmt"
  // "net/http"
)

type (

  // the standard Json returned from a server
  WireResponse struct {

    // non-json stuff
    HttpStatusCode                int
    Session                       *Session

    // json stuff
    Name                  string `json:"name"`
    SessionID             string `json:"sessionId"`
    Status                   int `json:"status"`
    Value        json.RawMessage `json:"value"`

  }

)

// Checks the values of the response from a webdriver server.  If the
// http response code is 200 and the Status from the webdriver is 0, then,
// the request is considered successful.
func (s *WireResponse) Success() bool {
  return s.HttpStatusCode == 200 && s.Status == 0
}

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
  } else {
    err = errors.New("WebElement.Value is nil")
  }

  return value, err
}

// Convenience method to unmarshal the json.RawMessage Value to a string.
func (s *WireResponse) WebElement() (value *WebElement, err error) {

  value = &WebElement{}

  if s.Value != nil {
    err = json.Unmarshal(s.Value, value)
    value.Session = s.Session
  } else {
    err = errors.New("WireResponse.Value is nil")
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
  } else {
    err = errors.New("WireResponse.Value is nil")
  }

  return value, err
}























