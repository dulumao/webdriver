package webdriver

import (
  "encoding/json"
  "testing"
)

////////////////////////////////////////////////////////////////
func TestGetStatus(t *testing.T) {

  for _, client := range clients {
    if wire := client.Status(); wire.Error == nil {

      var buffer map[string]interface{}
      if err := json.Unmarshal(wire.Response.Value, &buffer); err != nil {
        t.Error("could not Unmarshal: ", err)
      } else {

        if _, ok := buffer["build"]; !ok {
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

  for _, session := range getSessions("chrome") {
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
