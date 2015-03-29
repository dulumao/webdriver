package webdriver

import (
  "bytes"
  "encoding/json"
  // "fmt"
  // "net/http"
)

type (

  // WebElement - An object in the WebDriver API that represents a DOM element on the page.
  //
  // WebElement JSON Object - The JSON representation of a WebElement for transmission over the wire.
  //     Key Type  Description
  //     ELEMENT string  The opaque ID assigned to the element by the server. This ID should be used in all subsequent commands issued against the element.
  WebElement struct {
    Value               string `json:"element"`
  }

  // the standard Json returned from a server
  WireResponse struct {
    Name                  string `json:"name"`
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
  }

  return value, err
}

// Convenience method to unmarshal the json.RawMessage Value to a string.
func (s *WireResponse) WebElements() (value []*WebElement, err error) {

  if s.Value != nil {
    err = json.Unmarshal(s.Value, &value)
  }

  return value, err
}























