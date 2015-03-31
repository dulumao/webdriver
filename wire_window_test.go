package webdriver

import (
  "testing"
)

////////////////////////////////////////////////////////////////
func TestWindow(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  var tempSession *Session

  var sessions *Sessions

  for _, v := range clients {

    sleepForSeconds(1)

    sessions = v.GetSessions()
    if len(sessions.List) > 0 {

      currentSession := sessions.List[0]

      if wireResponse, err = currentSession.WindowHandle(); err == nil && wireResponse.Success() {
        if wireResponse.StringValue() != "" {

          windowHandle := wireResponse.StringValue()

          if tempSession, err = v.NewSession(); err == nil {

            defer tempSession.Delete()

            sleepForSeconds(1)

            if wireResponse, err = currentSession.Window(windowHandle); err != nil || !wireResponse.Success() {
              // TODO: find out why firefox server 500 error
              // t.Error(err, wireResponse.HttpStatusCode, wireResponse.Status, wireResponse.StringValue())
            }

            sleepForSeconds(1)

            if wireResponse, err = tempSession.DeleteWindow(); err != nil || !wireResponse.Success() {
              t.Error(err)
            }

            sleepForSeconds(1)

          } else {
            t.Error("unable to create new temporary session: ", err)
          }


        } else {
          t.Error(err, wireResponse, "window handle seems to be empty")
        }
      } else {
        t.Error(err, wireResponse)
      }

    } else {
      t.Error("No sessions found: client does not have any open sessions")
    }

  }

}

////////////////////////////////////////////////////////////////
func TestWindowHandle(t *testing.T) {

  var err error
  var wireResponse *WireResponse

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.WindowHandle(); err == nil && wireResponse.Success() {
        if wireResponse.StringValue() == "" {
          t.Error(err, wireResponse, "window handle seems to be empty")
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestWindowHandles(t *testing.T) {

  var err error
  var wireResponse *WireResponse

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.WindowHandles(); err == nil && wireResponse.Success() {

        if wireResponse.StringValue() == "" {
          t.Error(err, wireResponse, "window handle seems to be empty")
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestWindowSize(t *testing.T) {

  var err error
  var size *Size
  var windowHandle string
  var wireResponse *WireResponse

  // TODO: find out why firefox server 500
  t.Skip()

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.WindowHandle(); err == nil && wireResponse.Success() {
        if wireResponse.StringValue() != "" {

          windowHandle = wireResponse.StringValue()

          if wireResponse, err = v.Size(windowHandle); err == nil && wireResponse.Success() {
            if size, err = wireResponse.Size(); err == nil {

              size.Height = 500
              size.Width = 500

              if wireResponse, err = v.SetSize(windowHandle, size); err == nil && wireResponse.Success() {

                sleepForSeconds(1)

                if wireResponse, err = v.Maximize(windowHandle); err == nil && wireResponse.Success() {
                  sleepForSeconds(1)
                } else {
                  t.Error(err, wireResponse)
                }

              } else {
                t.Error(err, wireResponse)
              }

            } else {
              t.Error(err, wireResponse)
            }

          } else {
            t.Error(err, wireResponse)
          }

        } else {
          t.Error(err, wireResponse, "window handle seems to be empty")
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestWindowPosition(t *testing.T) {

  var err error
  var point *Point
  var windowHandle string
  var wireResponse *WireResponse

  // TODO: find out why firefox server 500
  t.Skip()

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.WindowHandle(); err == nil && wireResponse.Success() {
        if wireResponse.StringValue() != "" {

          windowHandle = wireResponse.StringValue()

          if wireResponse, err = v.Position(windowHandle); err == nil && wireResponse.Success() {
            if point, err = wireResponse.Point(); err == nil {

              point.X = 250
              point.Y = 250

              if wireResponse, err = v.SetPosition(windowHandle, point); err == nil && wireResponse.Success() {

              } else {
                t.Error(err, wireResponse)
              }

            } else {
              t.Error(err, wireResponse)
            }

          } else {
            t.Error(err, wireResponse)
          }

        } else {
          t.Error(err, wireResponse, "window handle seems to be empty")
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}














