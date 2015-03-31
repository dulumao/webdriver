package webdriver

import (
  "fmt"
  "net/http"
)

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

// POST /session/:sessionId/window/:windowHandle/maximize
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window/:windowHandle/maximize
//
// Maximize the specified window if not already maximized. If the :windowHandle URL parameter is "current", the currently active window will be maximized.
func (s *Wire) Maximize(windowHandle string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest(fmt.Sprintf("/session/:sessionid/window/%v/maximize", windowHandle), nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /session/:sessionId/window/:windowHandle/position
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window/:windowHandle/position
//
// Get the position of the specified window. If the :windowHandle URL parameter is "current", the position of the currently active window will be returned.
//
//     Returns:
//       {x: number, y: number} The X and Y coordinates for the window, relative to the upper left corner of the screen.
func (s *Wire) Position(windowHandle string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.GetRequest(fmt.Sprintf("/session/:sessionid/window/%v/position", windowHandle), nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// POST /session/:sessionId/window/:windowHandle/position
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window/:windowHandle/position
//
// Change the position of the specified window. If the :windowHandle URL parameter is "current", the currently active window will be moved.
func (s *Wire) SetPosition(windowHandle string, point *Point) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest(fmt.Sprintf("/session/:sessionid/window/%v/position", windowHandle),
                                  &Params{"x": point.X, "y": point.Y}); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// POST /session/:sessionId/window/:windowHandle/size
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window/:windowHandle/size
//
// Change the size of the specified window. If the :windowHandle URL parameter is "current", the currently active window will be resized.
func (s *Wire) SetSize(windowHandle string, size *Size) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest(fmt.Sprintf("/session/:sessionid/window/%v/size", windowHandle),
                                  &Params{"height": size.Height, "width": size.Width}); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

// GET /session/:sessionId/window/:windowHandle/size
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window/:windowHandle/size
//
// Get the size of the specified window. If the :windowHandle URL parameter is "current", the size of the currently active window will be returned.
//
//     Returns:
//       {width: number, height: number} The size of the window.
func (s *Wire) Size(windowHandle string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.GetRequest(fmt.Sprintf("/session/:sessionid/window/%v/size", windowHandle), nil); err == nil {

    wireResponse, err = s.Do(req)

  }

  return wireResponse, err
}

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















