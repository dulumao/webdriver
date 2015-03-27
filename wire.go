package webdriver

import (
  "bytes"
  "encoding/json"
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

    // a url pointing to a running server supporting the JsonWireProtocol.
    // typically, http://localhost:7055 for firefox.
    BaseUrl string

    // represents a JsonWireProtocol Session ID.
    // The Session struct includes *Wire, so, SessionID is available
    // to individual sessions.
    //
    // Most of the JsonWireProtocol API calls require a session id.
    // Only a couple do not.  GetFullUrl() will search for :sessionid
    // and replace it with SessionID during API calls.
    //
    // By default, SessionID is "", so, there should be no impact
    // for API calls that do not require a :sessionid
    SessionID string
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

// // GET /session/:sessionId/source
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#GET_/session/:sessionId/source
// //
// // Get and return the browser's current page source as HTML.
// // wireResponse.StringValue() will contain the entire source as HTML.
// func (s *Wire) Source() (wireResponse *WireResponse, err error) {

//   if req, err := s.GetRequest("/session/:sessionid/source", nil); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // GET /status
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#GET_/status
// //
// // Query the status of the webdriver server.
// func (s *Wire) Status() (wireResponse *WireResponse, err error) {

//   if req, err := s.GetRequest("/status", nil); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // GET /session/:sessionId/title
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#GET_/session/:sessionId/title
// //
// // Get the current page title.
// func (s *Wire) Title() (wireResponse *WireResponse, err error) {

//   if req, err := s.GetRequest("/session/:sessionid/title", nil); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // GET /sessions
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/sessions
// //
// // Returns a list of the currently active sessions. Each session will be returned as a list of JSON objects with the following keys:
// //
// //    Key              Type      Description
// //    id               string    The session ID.
// //    capabilities     object    An object describing the session's capabilities.
// //
// //    Returns:
// //    {Array.<Object>} A list of the currently active sessions.
// func (s *Wire) Sessions() (wireResponse *WireResponse, err error) {

//   if req, err := s.GetRequest("/sessions", nil); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// GET /session/:sessionid
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId
//
// Retrieve the capabilities of the specified session.
//
//    Returns:
//    {object} An object describing the session's capabilities.
func (s *Wire) GetCapabilities() (wireResponse *WireResponse, err error) {

  if req, err := s.GetRequest("/session/:sessionid", nil); err == nil {

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

// // POST /session/:sessionid/timeouts
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/timeouts
// //
// // Configure the amount of time that a particular type of operation can execute for before they are aborted and a |Timeout| error is returned to the client.
// //
// //     type - {string} The type of operation to set the timeout for. Valid values are: "script" for script timeouts, "implicit" for modifying the implicit wait timeout and "page load" for setting a page load timeout.
// //     ms - {number} The amount of time, in milliseconds, that time-limited commands are permitted to run.
// func (s *Wire) Timeouts(type_value string, ms float64) (wireResponse *WireResponse, err error) {

//   if req, err := s.PostRequest("/session/:sessionid/timeouts",
//                               &Params{"type": type_value, "ms": ms}); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // POST /session/:sessionid/timeouts/async_script
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/timeouts/async_script
// //
// // Set the amount of time, in milliseconds, that asynchronous scripts executed by /session/:sessionId/execute_async are permitted to run before they are aborted and a |Timeout| error is returned to the client.
// //
// //     ms - {number} The amount of time, in milliseconds, that time-limited commands are permitted to run.
// func (s *Wire) TimeoutsAsyncScript(ms float64) (wireResponse *WireResponse, err error) {

//   if req, err := s.PostRequest("/session/:sessionid/timeouts/async_script", &Params{"ms": ms}); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // POST /session/:sessionid/timeouts/implicit_wait
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/timeouts/implicit_wait
// //
// // Set the amount of time the driver should wait when searching for elements. When searching for a single element, the driver should poll the page until an element is found or the timeout expires, whichever occurs first. When searching for multiple elements, the driver should poll the page until at least one element is found or the timeout expires, at which point it should return an empty list.
// //
// // If this command is never sent, the driver should default to an implicit wait of 0ms.
// //
// //     ms - {number} The amount of time to wait, in milliseconds. This value has a lower bound of 0.
// func (s *Wire) TimeoutsImplicitWait(ms float64) (wireResponse *WireResponse, err error) {

//   if req, err := s.PostRequest("/session/:sessionid/timeouts/implicit_wait", &Params{"ms": ms}); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // GET /session/:sessionId/window_handle
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window_handle
// //
// // Retrieve the current window handle.
// //
// //      Returns:
// //        {string} The current window handle.
// func (s *Wire) WindowHandle() (wireResponse *WireResponse, err error) {

//   var req *http.Request
//   if req, err = s.GetRequest("/session/:sessionid/window_handle", nil); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // GET /session/:sessionId/window_handles
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window_handles
// //
// // Retrieve the list of all window handles available to the session.
// //
// //      Returns:
// //        {Array.<string>} A list of window handles.
// func (s *Wire) WindowHandles() (wireResponse *WireResponse, err error) {

//   var req *http.Request
//   if req, err = s.GetRequest("/session/:sessionid/window_handles", nil); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // POST  /session/:sessionId/url
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/url
// //
// // Navigate to a new URL.
// //
// // Browser should navigate to the given url.  url is any valid http url
// // that you would normally enter in a browser.
// // 
// //      url - {string} The URL to navigate to.
// func (s *Wire) Url(url string) (wireResponse *WireResponse, err error) {

//   var req *http.Request
//   if req, err = s.PostRequest("/session/:sessionid/url",
//                               &Params{"url": url}); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // GET  /session/:sessionId/url
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/url
// //
// // Retrieve the URL of the current page.
// func (s *Wire) GetUrl() (wireResponse *WireResponse, err error) {

//   var req *http.Request
//   if req, err = s.GetRequest("/session/:sessionid/url", nil); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // POST  /session/:sessionId/forward
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/forward
// //
// // Navigate forwards in the browser history, if possible.
// //
// func (s *Wire) Forward() (wireResponse *WireResponse, err error) {

//   var req *http.Request
//   if req, err = s.PostRequest("/session/:sessionid/forward", nil); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // POST  /session/:sessionId/back
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/back
// //
// // Navigate forwards in the browser history, if possible.
// //
// func (s *Wire) Back() (wireResponse *WireResponse, err error) {

//   var req *http.Request
//   if req, err = s.PostRequest("/session/:sessionid/back", nil); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // POST  /session/:sessionId/refresh
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/refresh
// //
// // Refresh the current page.
// //
// func (s *Wire) Refresh() (wireResponse *WireResponse, err error) {

//   var req *http.Request
//   if req, err = s.PostRequest("/session/:sessionid/refresh", nil); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // POST  /session/:sessionId/execute
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/execute
// //
// // Inject a snippet of JavaScript into the page for execution in the context of the currently selected
// // frame. The executed script is assumed to be synchronous and the result of evaluating the script is
// // returned to the client.
// //
// // The script argument defines the script to execute in the form of a function body. The value returned
// // by that function will be returned to the client. The function will be invoked with the provided args
// // array and the values may be accessed via the arguments object in the order specified.
// //
// // Arguments may be any JSON-primitive, array, or JSON object. JSON objects that define a WebElement
// // reference will be converted to the corresponding DOM element. Likewise, any WebElements in the
// // script result will be returned to the client as WebElement JSON objects.
// //
// //      JSON Parameters:
// //        script - {string} The script to execute.
// //        args - {Array.<*>} The script arguments.
// //
// //      Returns:
// //        {*} The script result.
// func (s *Wire) Execute(script string, args string) (wireResponse *WireResponse, err error) {

//   var req *http.Request
//   if req, err = s.PostRequest("/session/:sessionid/execute", &Params{"script": script, "args": args}); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }



























