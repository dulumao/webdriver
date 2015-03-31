package webdriver

import (
  // "bytes"
  // "encoding/json"
  "fmt"
  "net/http"
)

type (

  // used primarilty as a convenience to construct maps being passed
  // to the http get, post, methods.
  Params map[string]interface{}

  // Represents all of the data and methods for the JsonWireProtocol API.
  // Include this in your client and make API calls.
  //
  // All JsonWireProtocol commands are attached to this struct.
  Wire struct {

    *WireHTTP

  }

)

// Sets the default values for a *WireHTTP.
func (s *Wire) SetDefaults() (err error) {

  s.WireHTTP = &WireHTTP{}

  return err
}

// POST  /session/:sessionId/back
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/back
//
// Navigate forwards in the browser history, if possible.
//
func (s *Wire) Back() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/back", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /session/:sessionId/cookie
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/cookie
//
// Retrieve all cookies visible to the current page.
//
func (s *Wire) Cookie() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.GetRequest("/session/:sessionid/cookie", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// DELETE /session/:sessionId/cookie/:name
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/cookie/:name
//
// Delete the cookie with the given name. This command should be a no-op if there is no such cookie visible to the current page.
//
func (s *Wire) DeleteCookie(name string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.DeleteRequest(fmt.Sprintf("/session/:sessionid/cookie/%v", name), nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// DELETE /session/:sessionId/cookie
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/cookie
//
// Delete all cookies visible to the current page.
//
func (s *Wire) DeleteCookies() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.DeleteRequest("/session/:sessionid/cookie", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// POST /session/:sessionId/cookie
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/cookie
//
// Set a cookie. If the cookie path is not specified, it should be set to "/". Likewise, if the domain is omitted, it should default to the current page's domain.
//
func (s *Wire) SetCookie(value *Cookie) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/cookie", &Params{"cookie": value}); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// DELETE /session/:sessionid
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId
//
// Delete the session.
//
func (s *Wire) Delete() (wireResponse *WireResponse, err error) {

  if req, err := s.DeleteRequest("/session/:sessionid", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /session/:sessionid
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId
//
// Retrieve the capabilities of the specified session.
//
//    Returns:
//    {object} An object describing the session's capabilities.
func (s *Wire) GetSession() (wireResponse *WireResponse, err error) {

  if req, err := s.GetRequest("/session/:sessionid", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// POST  /session/:sessionId/forward
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/forward
//
// Navigate forwards in the browser history, if possible.
//
func (s *Wire) Forward() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/forward", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// POST /session/:sessionId/keys
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/keys
//
// Send a sequence of key strokes to the active element. This command is similar to the send keys command in every aspect except the implicit termination: The modifiers are not released at the end of the call. Rather, the state of the modifier keys is kept between calls, so mouse interactions can be performed while modifier keys are depressed.
//
func (s *Wire) Keys(value []string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/keys", &Params{"value": value}); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /session/:sessionId/location
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/location
//
// Get the current geo location.
//
func (s *Wire) Location() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.GetRequest("/session/:sessionid/location", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// POST /session/:sessionId/location
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/location
//
// Set the current geo location.
//
func (s *Wire) SetLocation(value *Location) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/location", &Params{"location": value}); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// POST  /session/:sessionId/refresh
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/refresh
//
// Refresh the current page.
//
func (s *Wire) Refresh() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/refresh", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// POST /session
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#POST_/session
//
// See webdriver.NewSession() for more detail.
func (s *Wire) Session(capabilities ...*Capabilities) (wireResponse *WireResponse, err error) {

  // capabilities are optioinal to the newSession method, but,
  // not optional to the server.
  // desired needs to be first in line, then, required.
  // know a better way?  I'm all ears...
  x := map[string]*Capabilities{}
  if len(capabilities) > 0 && capabilities[0] != nil {
    x["desiredCapabilities"] = capabilities[0]
  } else {
    x["desiredCapabilities"] = &Capabilities{}
  }

  if len(capabilities) > 1 && capabilities[1] != nil {
    x["requiredCapabilities"] = capabilities[1]
  } else {
    x["requiredCapabilities"] = &Capabilities{}
  }

  // physically make the API call.
  var req *http.Request
  if req, err = s.PostRequest("/session", x); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /sessions
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/sessions
//
// Returns a list of the currently active sessions. Each session will be returned as a list of JSON objects with the following keys:
//
//    Key              Type      Description
//    id               string    The session ID.
//    capabilities     object    An object describing the session's capabilities.
//
//    Returns:
//    {Array.<Object>} A list of the currently active sessions.
func (s *Wire) Sessions() (wireResponse *WireResponse, err error) {

  if req, err := s.GetRequest("/sessions", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /session/:sessionId/source
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#GET_/session/:sessionId/source
//
// Get and return the browser's current page source as HTML.
// wireResponse.StringValue() will contain the entire source as HTML.
//
// Source will return a wireResponse struct.  Value will contain a json.RawMessage value
// returned from the server.  Firefox and chrome return different encodings, so, the raw
// bytes are left "as is" from the server.  You can use wireResponse.UnmarshalValue() to attempt
// to decode the value into a normal string.
func (s *Wire) Source() (wireResponse *WireResponse, err error) {

  if req, err := s.GetRequest("/session/:sessionid/source", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /status
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#GET_/status
//
// Query the status of the webdriver server.
func (s *Wire) Status() (wireResponse *WireResponse, err error) {

  if req, err := s.GetRequest("/status", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /session/:sessionId/title
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#GET_/session/:sessionId/title
//
// Get the current page title.
func (s *Wire) Title() (wireResponse *WireResponse, err error) {

  if req, err := s.GetRequest("/session/:sessionid/title", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// POST  /session/:sessionId/url
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/url
//
// Navigate to a new URL.
//
// Browser should navigate to the given url.  url is any valid http url
// that you would normally enter in a browser.
// 
//      url - {string} The URL to navigate to.
func (s *Wire) Url(url string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/url",
                              &Params{"url": url}); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}





























