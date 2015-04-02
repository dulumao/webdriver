package webdriver

import (
  "testing"
)

////////////////////////////////////////////////////////////////
func TestWindow(t *testing.T) {

  t.Skip()

  for _, client := range clients {

    sleepForSeconds(1)

    current_sessions := client.GetSessions()
    if len(current_sessions) > 0 {

      currentSession := current_sessions[0]
      if currentSession.WindowHandle(); currentSession.Success() {

        sleepForSeconds(1)

        if windowHandle, err := currentSession.StringValue(); err == nil && windowHandle != "" {

          if tempSession, err := client.Session(); err == nil {

            defer tempSession.DeleteSession()

            sleepForSeconds(1)

            // TODO: figure out why firefox 500 server error
            if currentSession.Window(windowHandle); !currentSession.Success() {
              t.Error("unable to select window", currentSession.Error)
            }

            sleepForSeconds(1)

            if tempSession.DeleteWindow(); !tempSession.Success() {
              t.Error("unable to delete window", tempSession.Error)
            }

            sleepForSeconds(1)

          } else {
            t.Error("unable to create new temporary session: ", err)
          }

        } else {
          t.Error("window handle seems to be empty", err, currentSession.Response.Value)
        }

      } else {
        t.Error("cannot get window handle", currentSession.Error)
      }

    } else {
      t.Error("No sessions found: client does not have any open sessions")
    }

  }

}

////////////////////////////////////////////////////////////////
func TestWindowHandle(t *testing.T) {

  for _, session := range getSessions() {
    if session.Url("http://localhost:8080/index.html"); session.Success() {

      sleepForSeconds(1)

      if session.WindowHandle(); session.Success() {
        if windowHandle, err := session.StringValue(); err == nil {
          if len(windowHandle) <= 0 {
            t.Error("window handle appears to be empty", windowHandle)
          }
        } else {
          t.Error("could not get window handle: ", session.Error)
        }
      } else {
        t.Error("could not get window handle: ", session.Error)
      }
    } else {
      t.Error("could not get url: ", session.Error)
    }

  }


}

////////////////////////////////////////////////////////////////
func TestWindowHandles(t *testing.T) {

  for _, session := range getSessions() {
    if session.Url("http://localhost:8080/index.html"); session.Success() {

      sleepForSeconds(1)

      if session.WindowHandles(); session.Success() {
        if windowHandle, err := session.StringValue(); err == nil {
          if len(windowHandle) <= 0 {
            t.Error("window handle appears to be empty", windowHandle)
          }
        } else {
          t.Error("could not get window handle: ", session.Error)
        }
      } else {
        t.Error("could not get window handle: ", session.Error)
      }
    } else {
      t.Error("could not get url: ", session.Error)
    }

  }


}

////////////////////////////////////////////////////////////////
func TestWindowSize(t *testing.T) {

//   // TODO: find out why firefox server 500
  t.Skip()

  for _, session := range getSessions() {

    if session.Url("http://localhost:8080/index.html"); session.Success() {

      sleepForSeconds(1)

      if session.WindowHandle(); session.Success() {

        if windowHandle, err := session.StringValue(); err == nil {
          if len(windowHandle) > 0 {
            if session.Size(windowHandle); session.Success() {
                sleepForSeconds(1)
                if size, err := session.GetSize(); err == nil {
                  size.Height = 500
                  size.Width = 500
                  if session.SetSize(windowHandle, size); session.Success() {
                    sleepForSeconds(1)
                    if session.Maximize(windowHandle); session.Success() {
                      sleepForSeconds(1)
                    } else {
                      t.Error("cannot set window size", session.Error, size)
                    }
                  } else {
                    t.Error("cannot set window size", session.Error, size)
                  }
                } else {
                  t.Error("cannot get window size", session.Error)
                }
              } else {
                t.Error("cannot get window size", session.Error)
              }
            } else {
              t.Error("cannot get window size", session.Error)
            }
          } else {
            t.Error("window handle appears to be empty", windowHandle)
          }

        } else {
          t.Error("could not get window handle: ", session.Error)
        }

    } else {
      t.Error("could not get url: ", session.Error)
    }

  }

}

////////////////////////////////////////////////////////////////
func TestWindowPosition(t *testing.T) {

//   // TODO: find out why firefox server 500
  t.Skip()

  for _, session := range getSessions() {

    if session.Url("http://localhost:8080/index.html"); session.Success() {

      sleepForSeconds(1)

      if session.WindowHandle(); session.Success() {

        if windowHandle, err := session.StringValue(); err == nil {
          if len(windowHandle) > 0 {
            if session.Position(windowHandle); session.Success() {
                sleepForSeconds(1)
                if point, err := session.GetPoint(); err == nil {
                  point.X = 250
                  point.Y = 200
                  if session.SetPosition(windowHandle, point); session.Success() {
                    sleepForSeconds(1)
                  } else {
                    t.Error("cannot set window position", session.Error, point)
                  }
                } else {
                  t.Error("cannot get window position", session.Error)
                }
              } else {
                t.Error("cannot get window position", session.Error)
              }
            } else {
              t.Error("cannot get window position", session.Error)
            }
          } else {
            t.Error("window handle appears to be empty", windowHandle)
          }

        } else {
          t.Error("could not get window handle: ", session.Error)
        }

    } else {
      t.Error("could not get url: ", session.Error)
    }

  }

}

////////////////////////////////////////////////////////////////
func TestFrame(t *testing.T) {

  // TODO: find out why firefox server 500
  t.Skip()

  for _, session := range getSessions() {

    if session.Url("http://localhost:8080/frame.html"); session.Success() {

      sleepForSeconds(1)

      if session.WindowHandle(); session.Success() {

        if windowHandle, err := session.StringValue(); err == nil {
          if len(windowHandle) > 0 {
            if session.Frame("main-frame"); session.Success() {
                sleepForSeconds(1)
                if session.FrameParent(); !session.Success() {
                  t.Error("cannot get frame parent", session.Error)
                }
              } else {
                t.Error("cannot get window frame", session.Error)
              }
            } else {
              t.Error("cannot get window frame", session.Error)
            }
          } else {
            t.Error("window handle appears to be empty", windowHandle)
          }

        } else {
          t.Error("could not get window handle: ", session.Error)
        }

    } else {
      t.Error("could not get url: ", session.Error)
    }

  }


}













