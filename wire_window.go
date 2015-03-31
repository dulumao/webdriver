package webdriver

import (
  "net/http"
)

// POST /session/:sessionId/window
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window
//
// Change focus to another window. The window to change focus to may be specified by its server assigned window handle, or by the value of its name attribute.
//
//     JSON Parameters:
//       name - {string} The window to change focus to.
func (s *Wire) Window(name string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/window", &Params{"name": name}); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// DELETE /session/:sessionId/window
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window
//
// Close the current window.
func (s *Wire) DeleteWindow() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.DeleteRequest("/session/:sessionid/window", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /session/:sessionId/window_handle
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window_handle
//
// Retrieve the current window handle.
//
//      Returns:
//        {string} The current window handle.
func (s *Wire) WindowHandle() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.GetRequest("/session/:sessionid/window_handle", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /session/:sessionId/window_handles
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window_handles
//
// Retrieve the list of all window handles available to the session.
//
//      Returns:
//        {Array.<string>} A list of window handles.
func (s *Wire) WindowHandles() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.GetRequest("/session/:sessionid/window_handles", nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}
