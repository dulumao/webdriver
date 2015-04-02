package webdriver

import (
  "encoding/json"
  "strings"
  "testing"
)

////////////////////////////////////////////////////////////////
func TestGetStatus(t *testing.T) {

  for _, client := range clients {
    if wire := client.Status(); wire.Error == nil {

      var params Params
      if err := json.Unmarshal(wire.Response.Value, &params); err != nil {
        t.Error("could not Unmarshal: ", err)
      } else {

        if _, ok := params["build"]; !ok {
          t.Log("WARNING: build not found in status, however, not considering this an error due to different browsers may not support it (but SHOULD).")
        }

      }

    } else {
      t.Error("could not get status: ", wire.Error)
    }
  }

}

////////////////////////////////////////////////////////////////
func TestGetTitle(t *testing.T) {

  for _, session := range getSessions() {
    if session.Url("http://localhost:8080/index.html").Title(); session.Error == nil {
      if title, err := session.StringValue(); err == nil {
        if title != "index" {
          t.Error("title should equal index: ", title)
        }
      } else {
        t.Error("could not get title: ", err)
      }
    } else {
      t.Error("could not get title: ", session.Error)
    }

  }

}

////////////////////////////////////////////////////////////////
func TestSessions(t *testing.T) {

  for _, client := range clients {

    sleepForSeconds(1)

    var params []Params

    if wire := client.WireSessions(); wire.Error == nil {

      if err := json.Unmarshal(wire.Response.Value, &params); err != nil {
        t.Log("WARNING: could not Unmarshal: ", err)
        t.Log("Should have been able to marshall.  however, some clients (firefox) don't seem to support it well.")
      } else {

        if len(params) <= 0 {
          t.Log("WARNING: params should have at least one item.  not an error.")
        }

      }

    }

  }

}

////////////////////////////////////////////////////////////////
func TestGetSession(t *testing.T) {

  for _, session := range getSessions() {
    if session.GetSession(); session.Error == nil {

      var params Params
      if err := json.Unmarshal(session.Response.Value, &params); err != nil {
        t.Error("could not Unmarshal: ", err)
      } else {

        if _, ok := params["platform"]; !ok {
          t.Log("WARNING: platform not found in params, however, not considering this an error due to different browsers may not support it (but SHOULD).")
        }

      }
    } else {
      t.Error("could not get session: ", session.Error)
    }

  }

}

////////////////////////////////////////////////////////////////
func TestNav01(t *testing.T) {

  for _, session := range getSessions() {

    t.Log("session id", session.SessionID)

    if session.Url("http://localhost:8080/step01.html").Title(); session.Error == nil {
      if title, err := session.StringValue(); err == nil {
        if title != "step 01" {
          t.Error("title should equal step 01: ", title)
        }
      } else {
        t.Error("could not get title: ", err)
      }
    } else {
      t.Error("could not get title: ", session.Error)
    }

    sleepForSeconds(1)

    if session.Url("http://localhost:8080/step02.html").Title(); session.Error == nil {
      if title, err := session.StringValue(); err == nil {
        if title != "step 02" {
          t.Error("title should equal step 02: ", title)
        }
      } else {
        t.Error("could not get title: ", err)
      }
    } else {
      t.Error("could not get title: ", session.Error)
    }

    sleepForSeconds(1)

    if session.Url("http://localhost:8080/step03.html").Title(); session.Error == nil {
      if title, err := session.StringValue(); err == nil {
        if title != "step 03" {
          t.Error("title should equal step 03: ", title)
        }
      } else {
        t.Error("could not get title: ", err)
      }
    } else {
      t.Error("could not get title: ", session.Error)
    }

    sleepForSeconds(1)

    if session.Back(); session.Error == nil {

      sleepForSeconds(1)

      if session.Title(); session.Error == nil {
        if title, err := session.StringValue(); err == nil {
          if title != "step 02" {
            t.Error("title should equal step 02: ", title)
          }
        } else {
          t.Error("could not get title: ", err)
        }
      } else {
        t.Error("could not get title: ", session.Error)
      }
    } else {
      t.Error("could not nav back: ", session.Error)
    }

    sleepForSeconds(1)

    if session.Back(); session.Error == nil {

      sleepForSeconds(1)

      if session.Title(); session.Error == nil {
        if title, err := session.StringValue(); err == nil {
          if title != "step 01" {
            t.Error("title should equal step 01: ", title)
          }
        } else {
          t.Error("could not get title: ", err)
        }
      } else {
        t.Error("could not get title: ", session.Error)
      }
    } else {
      t.Error("could not nav back: ", session.Error)
    }

    sleepForSeconds(1)

    if session.Forward(); session.Error == nil {

      sleepForSeconds(1)

      if session.Title(); session.Error == nil {
        if title, err := session.StringValue(); err == nil {
          if title != "step 02" {
            t.Error("title should equal step 02: ", title)
          }
        } else {
          t.Error("could not get title: ", err)
        }
      } else {
        t.Error("could not get title: ", session.Error)
      }
    } else {
      t.Error("could not nav back: ", session.Error)
    }

    sleepForSeconds(1)

    if session.Forward(); session.Error == nil {

      sleepForSeconds(1)

      if session.Title(); session.Error == nil {
        if title, err := session.StringValue(); err == nil {
          if title != "step 03" {
            t.Error("title should equal step 03: ", title)
          }
        } else {
          t.Error("could not get title: ", err)
        }
      } else {
        t.Error("could not get title: ", session.Error)
      }
    } else {
      t.Error("could not nav back: ", session.Error)
    }

    if session.Refresh(); session.Error == nil {

      sleepForSeconds(1)

      if session.Title(); session.Error == nil {
        if title, err := session.StringValue(); err == nil {
          if title != "step 03" {
            t.Error("title should equal step 03: ", title)
          }
        } else {
          t.Error("could not get title: ", err)
        }
      } else {
        t.Error("could not get title: ", session.Error)
      }
    } else {
      t.Error("could not nav back: ", session.Error)
    }

  }

}

////////////////////////////////////////////////////////////////
func TestGetSource(t *testing.T) {

  for _, session := range getSessions() {
    if session.Url("http://localhost:8080/index.html"); session.Error == nil {

    sleepForSeconds(1)

      if session.Source(); session.Error == nil {

        if source, err := session.UnmarshalValue(); err == nil {

          if !strings.Contains(source, "</body>") {
            t.Error("source should contain </body>: ", source)
          }

        } else {
          t.Error("could not get source: ", err)
        }

      } else {
        t.Error("could not get source: ", session.Error)
      }

    } else {
      t.Error("could not get source: ", session.Error)
    }

  }

}

////////////////////////////////////////////////////////////////
func TestGeoLocation(t *testing.T) {

  // firefox doesn't support this command very well (maybe not at all.  did not look)

  for _, session := range getSessions("chrome") {
    if session.Url("http://localhost:8080/index.html"); session.Success() {

      sleepForSeconds(1)

      if session.SetLocation(&Location{Altitude: 0, Latitude: 50, Longitude: 70}); session.Success() {

        sleepForSeconds(1)

        if session.Location(); session.Success() {

          var location *Location
          if err := json.Unmarshal(session.Response.Value, &location); err == nil {
            if location.Altitude != 0 {
              t.Error("Altitude should be zero: ", location.Altitude)
            }
            if location.Latitude != 50 {
              t.Error("Latitude should be zero: ", location.Latitude)
            }
            if location.Longitude != 70 {
              t.Error("Longitude should be zero: ", location.Longitude)
            }
          } else {
            t.Error("could not get location: ", err)
          }

        } else {
          t.Error("could not get location: ", session.Error)
        }

      } else {
        t.Error("could not get location: ", session.Error)
      }

    } else {
      t.Error("could not get location: ", session.Error)
    }
  }

}

////////////////////////////////////////////////////////////////
func TestKeys(t *testing.T) {

  // keys := []string{PageDown, PageDown, PageDown, PageDown}
  keys := []string{Control, End}

  // TODO: figure out why this does not work with Firefox.  It should
  for _, session := range getSessions() {
    if session.Url("http://localhost:8080/longpage.html"); session.Error == nil {

      sleepForSeconds(1)

      if session.Keys(keys); session.Error == nil {

        sleepForSeconds(1)
        // sleepForSeconds(4)

      } else {
        t.Error("could not send keys : ", session.Error)
      }

    } else {
      t.Error("could not get url: ", session.Error)
    }

  }

}

////////////////////////////////////////////////////////////////
func TestCookie(t *testing.T) {

  // TODO: why chromedriver doesn't work and firefox does
  for _, session := range getSessions("firefox") {
    if session.Url("http://localhost:8080/cookies.html"); session.Success() {
    // if session.Url("http://www.google.com/"); session.Success() {

      sleepForSeconds(1)

      if session.Cookie(); session.Success() {

        sleepForSeconds(1)

        if cookies, err := session.Cookies(); err == nil {
          if len(cookies) <= 0 {
            t.Error("should have returned at least one cookie", cookies)
            t.Log(session.StringValue())
          }
        } else {
          t.Error("cookie: ", err, cookies)
        }

      } else {
        t.Error("could not get source: ", session.Error)
      }

    } else {
      t.Error("could not get url: ", session.Error)
    }

  }

}













