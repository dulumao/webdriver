package webdriver

import (
  // "bytes"
  // "encoding/json"
  // "fmt"
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

    // // a url pointing to a running server supporting the JsonWireProtocol.
    // // typically, http://localhost:7055 for firefox.
    // BaseUrl string

    // // represents a JsonWireProtocol Session ID.
    // // The Session struct includes *Wire, so, SessionID is available
    // // to individual sessions.
    // //
    // // Most of the JsonWireProtocol API calls require a session id.
    // // Only a couple do not.  GetFullUrl() will search for :sessionid
    // // and replace it with SessionID during API calls.
    // //
    // // By default, SessionID is "", so, there should be no impact
    // // for API calls that do not require a :sessionid
    // SessionID string

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





























