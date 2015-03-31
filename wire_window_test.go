package webdriver

import (
  "testing"
)

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















