package webdriver

import (
  // "log"
  "testing"
)

////////////////////////////////////////////////////////////////
func TestWork(t *testing.T) {

  var client Client
  // var wire *Wire

  client = &Chrome{
                      LogPath: "support/drivers/chromedriver.log",
                      PathExec: "support/drivers/chromedriver",
                      Verbose: true,
                    }

  if err := client.Run(); err == nil {

    defer client.Close()

    if value, err := client.Status().StringValue(); err == nil {
      t.Log("success!!", value)
    } else {
      t.Error("error: ", err)
    }

    if session, err := client.Session(); err == nil {
      if session.Url("http://www.google.com/").Title(); session.Error == nil {
        if title, err := session.StringValue(); err == nil {
          t.Log("title", title)
        }
      }
    }

  }


}
