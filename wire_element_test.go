package webdriver

import (
  "fmt"
  // "log"
  // "encoding/json"
  // "strings"
  "testing"
)

////////////////////////////////////////////////////////////////
func TestElement(t *testing.T) {

  for _, v := range sessions {
    if _, err := v.Url("http://localhost:8080/element.html"); err == nil {

      sleepForSeconds(1)

      // if wireResponse, err := v.Source(); err == nil {
      //   value, _ := wireResponse.UnmarshalValue()
      //   if !strings.Contains(value, "<div>verify source is working</div>") {
      //     t.Error("should contain: <div>verify source is working</div> => ", value)
      //   }
      // }

      if wireResponse, err := v.Element("id", "main-div"); err == nil {

webElement, err := wireResponse.WebElement()

fmt.Println("wireResponse", webElement, err)
        // if wireResponse.StringValue() != "title check" {
        //   t.Error("<title> tag should be title check =>", wireResponse.StringValue())
        // }
      }

    }
  }

}

