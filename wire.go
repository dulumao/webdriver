package webdriver

import (
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"
  "strings"
)

type (

  // Represents a cookie.
  Cookie struct {
    Domain string `json:"domain"`
    Expiry uint   `json:"expiry"`
    Name   string `json:"name"`
    Path   string `json:"path"`
    Secure bool   `json:"secure"`
    Value  string `json:"value"`
  }

  // Geo location.
  Location struct {
    Altitude        int `json:"altitude"`
    Latitude        int `json:"latitude"`
    Longitude       int `json:"longitude"`
  }

  // used primarilty as a convenience to construct maps being passed
  // to the http get, post, methods.
  Params map[string]interface{}

  // Represents an X,Y coordinate.
  Point struct {
    X     int `json:"x"`
    Y     int `json:"y"`
  }

  Session struct {

    ActualCapabilities *ActualCapabilities

    *Wire

  }

  // Represents an X,Y coordinate.
  Size struct {
    Height      int `json:"height"`
    Width       int `json:"width"`
  }

  // Represents all of the data and methods for the JsonWireProtocol API.
  // Include this in your client and make API calls.
  //
  // All JsonWireProtocol commands are attached to this struct.
  Wire struct {

    // a url pointing to a running server supporting the JsonWireProtocol.
    // typically, http://localhost:7055 for firefox.
    BaseUrl string

    // represents the most recent error
    Error error

    Response *WireResponse

    // represents a JsonWireProtocol Session ID.
    // The Session struct includes *WireHTTP, so, SessionID is available
    // to individual sessions.
    //
    // Most of the JsonWireProtocol API calls require a session id.
    // Only a couple do not.  GetFullUrl() will search for :sessionid
    // and replace it with SessionID during API calls.
    //
    // By default, SessionID is "", so, there should be no impact
    // for API calls that do not require a :sessionid
    SessionID string

    Sessions []*Session

  }

)

// Builds a complete url for a request including host and port.
// Relies on the current value of BaseUrl and SessionID.
//
//   // given:
//     BaseUrl = "http://localhost:7055"
//     SessionID = "my-session-id"
//
//   // the following call
//   BuildFullUrl("/session/:sessionid/forward")
//
//   // would produce
//   http://localhost:7055/session/my-session-id/forward
//
func (s *Wire) BuildFullUrl(url string) string {
  return fmt.Sprintf("%v%v", s.BaseUrl, strings.Replace(url, ":sessionid", s.SessionID, -1))
}

// POST  /session/:sessionId/back
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/back
//
// Navigate forwards in the browser history, if possible.
//
func (s *Wire) Back() *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest("/session/:sessionid/back", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// Closes all of the active sessions.
func (s *Wire) CloseSessions() *Wire {

  for _, v := range s.Sessions {
    v.DeleteSession()
  }

  return s
}

// GET /session/:sessionId/cookie
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/cookie
//
// Retrieve all cookies visible to the current page.
//
func (s *Wire) Cookie() *Wire {

  // TODO: ALL of the cookie related code needs to be fully tested.
  // will revisit later.
  var req *http.Request
  if req, s.Error = s.GetRequest("/session/:sessionid/cookie", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// DELETE /session/:sessionId/cookie/:name
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/cookie/:name
//
// Delete the cookie with the given name. This command should be a no-op if there is no such cookie visible to the current page.
//
func (s *Wire) DeleteCookie(name string) *Wire {

  var req *http.Request
  if req, s.Error = s.DeleteRequest(fmt.Sprintf("/session/:sessionid/cookie/%v", name), nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// DELETE /session/:sessionId/cookie
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/cookie
//
// Delete all cookies visible to the current page.
//
func (s *Wire) DeleteCookies() *Wire {

  var req *http.Request
  if req, s.Error = s.DeleteRequest("/session/:sessionid/cookie", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}


// POST /session/:sessionId/cookie
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/cookie
//
// Set a cookie. If the cookie path is not specified, it should be set to "/". Likewise, if the domain is omitted, it should default to the current page's domain.
//
func (s *Wire) SetCookie(value *Cookie) *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest("/session/:sessionid/cookie", &Params{"cookie": value}); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// DELETE /session/:sessionid
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId
//
// Delete the session.
//
func (s *Wire) DeleteSession() *Wire {

  var req *http.Request
  if req, s.Error = s.DeleteRequest("/session/:sessionid", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// GET /session/:sessionid
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId
//
// Retrieve the capabilities of the specified session.
//
//    Returns:
//    {object} An object describing the session's capabilities.
func (s *Wire) GetSession() *Wire {

  var req *http.Request
  if req, s.Error = s.GetRequest("/session/:sessionid", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST  /session/:sessionId/forward
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/forward
//
// Navigate forwards in the browser history, if possible.
//
func (s *Wire) Forward() *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest("/session/:sessionid/forward", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST /session/:sessionId/keys
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/keys
//
// Send a sequence of key strokes to the active element. This command is similar to the send keys command in every aspect except the implicit termination: The modifiers are not released at the end of the call. Rather, the state of the modifier keys is kept between calls, so mouse interactions can be performed while modifier keys are depressed.
//
func (s *Wire) Keys(value []string) *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest("/session/:sessionid/keys", &Params{"value": value}); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// GET /session/:sessionId/location
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/location
//
// Get the current geo location.
//
func (s *Wire) Location() *Wire {

  var req *http.Request
  if req, s.Error = s.GetRequest("/session/:sessionid/location", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST /session/:sessionId/location
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/location
//
// Set the current geo location.
//
func (s *Wire) SetLocation(value *Location) *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest("/session/:sessionid/location", &Params{"location": value}); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST  /session/:sessionId/refresh
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/refresh
//
// Refresh the current page.
//
func (s *Wire) Refresh() *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest("/session/:sessionid/refresh", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST /session
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#POST_/session
//
// See webdriver.NewSession() for more detail.
//
// Creates a new session for a Client.  An JsonWireProtocol call is made
// to establish a session with a server.  The new session is added to the
// list of active sessions and returned to the caller.
//
// Capabilities are optional, however, it you define them, then, you must
// pass them in a specific order to this method.  Desired first, then, Required.
// Capabilities are currently implemented as a simple map and quite frankly
// I high doubt there will be much need to even support passing capabilities.
// However, it is in the spec, so, there is minimal support for it.
//
//      session, err := client.NewSession(
//               &webdriver.Capabilities{"Platform": "Linux"}, // desired
//               &webdriver.Capabilities{})                    // required
//
// When a new session is created, the server will return the actual capabilities
// currently supported.  An ActualCapabilities struct is created and attached
// to the returned session.
func (s *Wire) Session(values ...*Capabilities) (session *Session, err error) {

  // capabilities are optioinal to the newSession method, but,
  // not optional to the server.
  // desired needs to be first in line, then, required.
  // know a better way?  I'm all ears...
  capabilities := map[string]*Capabilities{
    "desiredCapabilities": &Capabilities{},
    "requiredCapabilities": &Capabilities{},
  }

  if len(values) > 0 && values[0] != nil {
    capabilities["desiredCapabilities"] = values[0]
  }

  if len(values) > 1 && values[1] != nil {
    capabilities["requiredCapabilities"] = values[1]
  }

  var req *http.Request
  if req, s.Error = s.PostRequest("/session", capabilities); s.Error == nil {

    if s.Response, s.Error = s.Do(req); s.Success() {

      // seems like everything went as planned.
      // create a new session and initialize it.
      session = &Session{}
      session.Wire = &Wire{}

      // setting the BaseUrl on the session is critical for http requests
      session.BaseUrl = s.BaseUrl

      // the Session ID returned during the API call.
      session.SessionID = s.Response.SessionID

      // extract the actual capabilities from the response and attach
      // them to the session
      actualCapabilities := &ActualCapabilities{}
      if s.Error = json.Unmarshal(s.Response.Value, actualCapabilities); s.Error == nil {

        session.ActualCapabilities = actualCapabilities

        // add the newly created session to the list of sessions
        s.Sessions = append(s.Sessions, session)

      }

    }

  }

  return session, s.Error
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
func (s *Wire) WireSessions() *Wire {

  var req *http.Request
  if req, s.Error = s.GetRequest("/sessions", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
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
func (s *Wire) Source() *Wire {

  var req *http.Request
  if req, s.Error = s.GetRequest("/session/:sessionid/source", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// GET /status
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#GET_/status
//
// Query the status of the webdriver server.
// func (s *Wire) Status() (wireResponse *WireResponse, err error) {
func (s *Wire) Status() *Wire {

  var req *http.Request
  if req, s.Error = s.GetRequest("/status", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// Extracts a WireResponse.Value as a string.
func (s *Wire) StringValue() (value string, err error) {

  if s.Success() && s.Response.Value != nil {
    value = string(bytes.Trim(s.Response.Value, "{}\""))
  }

  return value, s.Error
}

// Convenience method to unmarshal the json.RawMessage Value to a Cookie.
func (s *Wire) Cookies() (value []*Cookie, err error) {

  if s.Success() && s.Response.Value != nil {
    s.Error = json.Unmarshal(s.Response.Value, &value)
  }

  return value, s.Error
}

// Convenience method to unmarshal the json.RawMessage Value to a string.
func (s *Wire) UnmarshalValue() (value string, err error) {

  if s.Success() && s.Response.Value != nil {
    s.Error = json.Unmarshal(s.Response.Value, &value)
  }

  return value, s.Error
}

// Checks the values of the response from a webdriver server.  If the
// http response code is 200 and the Status from the webdriver is 0, then,
// the request is considered successful.
func (s *Wire) Success() bool {
  return s.Error == nil && s.Response.HttpStatusCode == 200 && s.Response.Status == 0
}

// GET /session/:sessionId/title
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#GET_/session/:sessionId/title
//
// Get the current page title.
func (s *Wire) Title() *Wire {

  var req *http.Request
  if req, s.Error = s.GetRequest("/session/:sessionid/title", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
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
func (s *Wire) Url(url string) *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest("/session/:sessionid/url", &Params{"url": url}); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

