package webdriver

import (
  "encoding/json"
  "fmt"
  "net/http"
)

// DELETE /session/:sessionId/window
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window
//
// Close the current window.
func (s *Wire) DeleteWindow() *Wire {

  var req *http.Request
  if req, s.Error = s.DeleteRequest("/session/:sessionid/window", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST /session/:sessionId/frame
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/frame
//
// Change focus to another frame on the page. If the frame id is null, the server should switch to the page's default content.
func (s *Wire) Frame(id string) *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest("/session/:sessionid/frame", &Params{"id": id}); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST /session/:sessionId/frame/parent
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/frame/parent
//
// Change focus to the parent context. If the current context is the top level browsing context, the context remains unchanged.
func (s *Wire) FrameParent() *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest("/session/:sessionid/frame/parent", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST /session/:sessionId/window/:windowHandle/maximize
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window/:windowHandle/maximize
//
// Maximize the specified window if not already maximized. If the :windowHandle URL parameter is "current", the currently active window will be maximized.
func (s *Wire) Maximize(windowHandle string) *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest(fmt.Sprintf("/session/:sessionid/window/%v/maximize", windowHandle), nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// GET /session/:sessionId/window/:windowHandle/position
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window/:windowHandle/position
//
// Get the position of the specified window. If the :windowHandle URL parameter is "current", the position of the currently active window will be returned.
//
//     Returns:
//       {x: number, y: number} The X and Y coordinates for the window, relative to the upper left corner of the screen.
func (s *Wire) Position(windowHandle string) *Wire {

  var req *http.Request
  if req, s.Error = s.GetRequest(fmt.Sprintf("/session/:sessionid/window/%v/position", windowHandle), nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST /session/:sessionId/window/:windowHandle/position
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window/:windowHandle/position
//
// Change the position of the specified window. If the :windowHandle URL parameter is "current", the currently active window will be moved.
func (s *Wire) SetPosition(windowHandle string, point *Point) *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest(fmt.Sprintf("/session/:sessionid/window/%v/position", windowHandle),
                                  &Params{"x": point.X, "y": point.Y}); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST /session/:sessionId/window/:windowHandle/size
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window/:windowHandle/size
//
// Change the size of the specified window. If the :windowHandle URL parameter is "current", the currently active window will be resized.
func (s *Wire) SetSize(windowHandle string, size *Size) *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest(fmt.Sprintf("/session/:sessionid/window/%v/size", windowHandle),
                                  &Params{"height": size.Height, "width": size.Width}); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// GET /session/:sessionId/window/:windowHandle/size
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window/:windowHandle/size
//
// Get the size of the specified window. If the :windowHandle URL parameter is "current", the size of the currently active window will be returned.
//
//     Returns:
//       {width: number, height: number} The size of the window.
func (s *Wire) Size(windowHandle string) *Wire {

  var req *http.Request
  if req, s.Error = s.GetRequest(fmt.Sprintf("/session/:sessionid/window/%v/size", windowHandle), nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// POST /session/:sessionId/window
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window
//
// Change focus to another window. The window to change focus to may be specified by its server assigned window handle, or by the value of its name attribute.
//
//     JSON Parameters:
//       name - {string} The window to change focus to.
func (s *Wire) Window(name string) *Wire {

  var req *http.Request
  if req, s.Error = s.PostRequest("/session/:sessionid/window", &Params{"name": name}); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// GET /session/:sessionId/window_handle
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window_handle
//
// Retrieve the current window handle.
//
//      Returns:
//        {string} The current window handle.
func (s *Wire) WindowHandle() *Wire {

  var req *http.Request
  if req, s.Error = s.GetRequest("/session/:sessionid/window_handle", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// GET /session/:sessionId/window_handles
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/window_handles
//
// Retrieve the list of all window handles available to the session.
//
//      Returns:
//        {Array.<string>} A list of window handles.
func (s *Wire) WindowHandles() *Wire {

  var req *http.Request
  if req, s.Error = s.GetRequest("/session/:sessionid/window_handles", nil); s.Error == nil {

    s.Response, s.Error = s.Do(req)

  }

  return s
}

// Convenience method to unmarshal the json.RawMessage Value to a Point.
func (s *Wire) GetPoint() (value *Point, err error) {

  value = &Point{}

  if s.Success() && s.Response.Value != nil {
    s.Error = json.Unmarshal(s.Response.Value, value)
  }

  return value, s.Error
}

// Convenience method to unmarshal the json.RawMessage Value to a Size.
func (s *Wire) GetSize() (value *Size, err error) {

  value = &Size{}

  if s.Success() && s.Response.Value != nil {
    s.Error = json.Unmarshal(s.Response.Value, value)
  }

  return value, s.Error
}













