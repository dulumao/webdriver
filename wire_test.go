package webdriver

import (
  // "fmt"
  "log"
  "encoding/json"
  "strings"
  "testing"
)

////////////////////////////////////////////////////////////////
func TestStringValue(t *testing.T) {

  for _, v := range sessions {
    if wireResponse, err := v.GetSession(); err == nil {

      var value map[string]interface{}
      if err = json.Unmarshal(wireResponse.Value, &value); err != nil {
        t.Error(err)
      } else {

        if _, ok := value["platform"]; ok {
          x := strings.ToLower(value["platform"].(string))
          if x != "linux" {
            t.Error("result should have contained platform: Linux")
          }
        } else {
          t.Error("result should have contained platform: Linux")
        }
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestUrlSourceTitle(t *testing.T) {

  for _, v := range sessions {
    if _, err := v.Url("http://localhost:8080/source.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err := v.Source(); err == nil {
        value, _ := wireResponse.UnmarshalValue()
        if !strings.Contains(value, "<div>verify source is working</div>") {
          t.Error("should contain: <div>verify source is working</div> => ", value)
        }
      }

      if wireResponse, err := v.Title(); err == nil {
        if wireResponse.StringValue() != "title check" {
          t.Error("<title> tag should be title check =>", wireResponse.StringValue())
        }
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestSessions(t *testing.T) {

  for _, v := range sessions {

    sleepForSeconds(1)

    // TODO revisit this test later
    // should respond with an array of sessions
    if _, err := v.Sessions(); err == nil {
    // if wireResponse, err := v.Sessions(); err == nil {
      // if wireResponse != nil {
      // }
    }

    // status is another tough one.
    // not all webdrivers will respond the same
    // revisit later
    if wireResponse, err := v.Status(); err == nil {
      if wireResponse != nil {

        var value map[string]interface{}
        if err = json.Unmarshal(wireResponse.Value, &value); err != nil {
          t.Error(err)
        } else {
          if value == nil {
            t.Error("Status should have at least repsonded with something.")
          }
        }

      }
    }

  }

}

////////////////////////////////////////////////////////////////
func TestNav01(t *testing.T) {

  for _, v := range sessions {

    log.Println("session id", v.SessionID)

    if _, err := v.Url("http://localhost:8080/step01.html"); err == nil {

      sleepForSeconds(2)

      if wireResponse, err := v.Title(); err == nil {
        sleepForSeconds(2)
        if wireResponse.StringValue() != "step 01" {
          t.Error("<title> tag should be step 01 =>", wireResponse.StringValue())
        }
      }

    }

    if _, err := v.Url("http://localhost:8080/step02.html"); err == nil {

      sleepForSeconds(2)

      if wireResponse, err := v.Title(); err == nil {
        if wireResponse.StringValue() != "step 02" {
          t.Error("<title> tag should be step 02 =>", wireResponse.StringValue())
        }
      }

    }

    if _, err := v.Url("http://localhost:8080/step03.html"); err == nil {

      sleepForSeconds(2)

      if wireResponse, err := v.Title(); err == nil {
        if wireResponse.StringValue() != "step 03" {
          t.Error("<title> tag should be step 03 =>", wireResponse.StringValue())
        }
      }

    }

    sleepForSeconds(2)

    if _, err := v.Back(); err == nil {
      sleepForSeconds(2)
      if wireResponse, err := v.Title(); err == nil {
        if wireResponse.StringValue() != "step 02" {
          t.Error("<title> tag should be step 02 =>", wireResponse.StringValue())
        }
      }
    }

    sleepForSeconds(2)

    if _, err := v.Back(); err == nil {
      sleepForSeconds(2)
      if wireResponse, err := v.Title(); err == nil {
        if wireResponse.StringValue() != "step 01" {
          t.Error("<title> tag should be step 01 =>", wireResponse.StringValue())
        }
      }
    }

    sleepForSeconds(2)

    if _, err := v.Forward(); err == nil {
      sleepForSeconds(2)
      if wireResponse, err := v.Title(); err == nil {
        if wireResponse.StringValue() != "step 02" {
          t.Error("<title> tag should be step 02 =>", wireResponse.StringValue())
        }
      }
    }

    sleepForSeconds(2)

    if _, err := v.Forward(); err == nil {
      sleepForSeconds(2)
      if wireResponse, err := v.Title(); err == nil {
        if wireResponse.StringValue() != "step 03" {
          t.Error("<title> tag should be step 03 =>", wireResponse.StringValue())
        }
      }
    }

    sleepForSeconds(2)

    if _, err := v.Refresh(); err == nil {

      sleepForSeconds(2)

      if wireResponse, err := v.Title(); err == nil {
        if wireResponse.StringValue() != "step 03" {
          t.Error("<title> tag should be step 03 =>", wireResponse.StringValue())
        }
      }
    }

    sleepForSeconds(2)

  }

}


////////////////////////////////////////////////////////////////
func TestKeys(t *testing.T) {

  var err error
  var wireResponse *WireResponse

  // TODO: think of a better way to test keys being sent
  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/longpage.html"); err == nil {

      sleepForSeconds(1)

      // keys := []string{PageDown, PageDown, PageDown, PageDown}
      keys := []string{Control, End}

      if wireResponse, err = v.Keys(keys); err == nil && wireResponse.Success() {

      sleepForSeconds(4)
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestGeoLocation(t *testing.T) {

  var err error
  var location *Location
  var wireResponse *WireResponse

  // TODO: revisit
  // looks like firefox not supporting this very well.

  t.Skip()

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/index.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.SetLocation(&Location{Altitude: 0, Latitude: 50, Longitude: 70}); err == nil && wireResponse.Success() {

        if wireResponse, err = v.Location(); err == nil && wireResponse.Success() {

          if location, err = wireResponse.Location(); err == nil {

            // TODO: should compare the values here
            t.Log("location: ", location)

          } else {
            t.Error(err, wireResponse.StringValue())
          }

        } else {
          t.Error(err, wireResponse.StringValue())
        }

      } else {
        t.Error(err, wireResponse.StringValue())
      }

    }
  }

}


////////////////////////////////////////////////////////////////
func TestCookie(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  var cookies []*Cookie

  // TODO: why chromedriver doesn't work and firefox does
  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/cookies.html"); err == nil {
    // if _, err = v.Url("http://www.google.com/"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Cookie(); err == nil && wireResponse.Success() {

        if cookies, err = wireResponse.Cookies(); err == nil {

          if len(cookies) <= 0 {
            t.Error("Server should have at least one cookie", wireResponse, cookies)
          }

        } else {
          t.Error(err, wireResponse)
        }

      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestSetCookie(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  // var cookies []*Cookie

  // TODO: complete the testing
  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/index.html"); err == nil {
    // if _, err = v.Url("http://www.google.com/"); err == nil {

      sleepForSeconds(1)

      cookie := &Cookie{
                   Domain: "localhost",
                  // Expires: time.Date(2020, 11, 23, 1, 5, 3, 0, time.UTC),
                     Name: "setmain",
                     Path: "/",
                    Value: "this-is-my-cookie-value-not-hard-to-decrypt",
                }

      if wireResponse, err = v.SetCookie(cookie); err == nil && wireResponse.Success() {

        // if cookies, err = wireResponse.Cookies(); err == nil {

        //   if len(cookies) <= 0 {
        //     t.Error("Server should have at least one cookie", wireResponse, cookies)
        //   }

        // } else {
        //   t.Error(err, wireResponse)
        // }

      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestDeleteCookies(t *testing.T) {

  var err error
  var wireResponse *WireResponse

  // TODO: complete the testing
  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/index.html"); err == nil {
    // if _, err = v.Url("http://www.google.com/"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.DeleteCookies(); err == nil && wireResponse.Success() {

        // if cookies, err = wireResponse.Cookies(); err == nil {

        //   if len(cookies) <= 0 {
        //     t.Error("Server should have at least one cookie", wireResponse, cookies)
        //   }

        // } else {
        //   t.Error(err, wireResponse)
        // }

      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestDeleteCookie(t *testing.T) {

  var err error
  var wireResponse *WireResponse

  // TODO: complete the testing
  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/index.html"); err == nil {
    // if _, err = v.Url("http://www.google.com/"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.DeleteCookie("main"); err == nil && wireResponse.Success() {

        // if cookies, err = wireResponse.Cookies(); err == nil {

        //   if len(cookies) <= 0 {
        //     t.Error("Server should have at least one cookie", wireResponse, cookies)
        //   }

        // } else {
        //   t.Error(err, wireResponse)
        // }

      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}























