package webdriver

import (
//   "bytes"
  "encoding/json"
//   "errors"
//   // "fmt"
//   // "net/http"
)

type (

//   // Represents a cookie.
//   Cookie struct {
//     Domain string `json:"domain"`
//     Expiry uint   `json:"expiry"`
//     Name   string `json:"name"`
//     Path   string `json:"path"`
//     Secure bool   `json:"secure"`
//     Value  string `json:"value"`
//   }

//   // Geo location.
//   Location struct {
//     Altitude        int `json:"altitude"`
//     Latitude        int `json:"latitude"`
//     Longitude       int `json:"longitude"`
//   }

//   // Represents an X,Y coordinate.
//   Point struct {
//     X     int `json:"x"`
//     Y     int `json:"y"`
//   }

//   // Represents an X,Y coordinate.
//   Size struct {
//     Height      int `json:"height"`
//     Width       int `json:"width"`
//   }

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

// // Convenience method to unmarshal the json.RawMessage Value to a Location.
// func (s *WireResponse) Location() (value *Location, err error) {

//   value = &Location{}

//   if s.Value != nil {
//     err = json.Unmarshal(s.Value, value)
//   } else {
//     err = errors.New("WireResponse.Value is nil")
//   }

//   return value, err
// }

// // Convenience method to unmarshal the json.RawMessage Value to a Point.
// func (s *WireResponse) Point() (value *Point, err error) {

//   value = &Point{}

//   if s.Value != nil {
//     err = json.Unmarshal(s.Value, value)
//   } else {
//     err = errors.New("WireResponse.Value is nil")
//   }

//   return value, err
// }

// // Convenience method to unmarshal the json.RawMessage Value to a Size.
// func (s *WireResponse) Size() (value *Size, err error) {

//   value = &Size{}

//   if s.Value != nil {
//     err = json.Unmarshal(s.Value, value)
//   } else {
//     err = errors.New("WireResponse.Value is nil")
//   }

//   return value, err
// }

// // Convenience method to unmarshal the json.RawMessage Value to a Cookie.
// func (s *WireResponse) Cookies() (value []*Cookie, err error) {

//   // value = &Cookie{}

//   if s.Value != nil {
//     err = json.Unmarshal(s.Value, &value)
//   } else {
//     err = errors.New("WireResponse.Value is nil")
//   }

//   return value, err
// }




















