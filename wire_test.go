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
