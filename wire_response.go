package webdriver

import (
//   "bytes"
  "encoding/json"
//   "errors"
//   // "fmt"
//   // "net/http"
)

type (

  // the standard Json returned from a server
  WireResponse struct {

    // non-json stuff
    HttpStatusCode                int
    // Session                       *Session

    // json stuff
    Name                  string `json:"name"`
    SessionID             string `json:"sessionId"`
    Status                   int `json:"status"`
    Value        json.RawMessage `json:"value"`

  }

)

// // Convenience method to unmarshal the json.RawMessage Value to a string.
// func (s *WireResponse) WebElement() (value *WebElement, err error) {

//   value = &WebElement{}

//   if s.Value != nil {
//     err = json.Unmarshal(s.Value, value)
//     value.Session = s.Session
//   } else {
//     err = errors.New("WireResponse.Value is nil")
//   }

//   return value, err
// }

// // Convenience method to unmarshal the json.RawMessage Value to a string.
// func (s *WireResponse) WebElements() (value []*WebElement, err error) {

//   if s.Value != nil {
//     err = json.Unmarshal(s.Value, &value)
//     for _, v := range value {
//       v.Session = s.Session
//     }
//   } else {
//     err = errors.New("WireResponse.Value is nil")
//   }

//   return value, err
// }























