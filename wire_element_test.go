package webdriver

import (
  // "fmt"
  // "log"
  // "encoding/json"
  // "strings"
  "testing"
)

////////////////////////////////////////////////////////////////
func TestElement(t *testing.T) {

  var err error
  var wireResponse *WireResponse

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/element.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("class name", "main-div-class"); err == nil {
        if webElement, err := wireResponse.WebElement(); err == nil {
          if webElement == nil || webElement.Value == "" {
            t.Error("could not find element", webElement.Value)
          }
        } else {
          t.Error("could not find element")
        }

      }

      if wireResponse, err = v.Element("id", "main-div-id"); err == nil {
        if webElement, err := wireResponse.WebElement(); err == nil {
          if webElement == nil || webElement.Value == "" {
            t.Error("could not find element", webElement.Value)
          }
        } else {
          t.Error("could not find element")
        }

      }

      if wireResponse, err = v.Element("name", "main-div-name"); err == nil {
        if webElement, err := wireResponse.WebElement(); err == nil {
          if webElement == nil || webElement.Value == "" {
            t.Error("could not find element", webElement.Value)
          }
        } else {
          t.Error("could not find element")
        }

      }

      if wireResponse, err = v.Element("link text", "link text"); err == nil {
        if webElement, err := wireResponse.WebElement(); err == nil {
          if webElement == nil || webElement.Value == "" {
            t.Error("could not find element", webElement.Value)
          }
        } else {
          t.Error("could not find element")
        }

      }

      if wireResponse, err = v.Element("partial link text", "partial"); err == nil {
        if webElement, err := wireResponse.WebElement(); err == nil {
          if webElement == nil || webElement.Value == "" {
            t.Error("could not find element", webElement.Value)
          }
        } else {
          t.Error("could not find element")
        }

      }

      if wireResponse, err = v.Element("tag name", "mytag"); err == nil {
        if webElement, err := wireResponse.WebElement(); err == nil {
          if webElement == nil || webElement.Value == "" {
            t.Error("could not find element", webElement.Value)
          }
        } else {
          t.Error("could not find element")
        }

      }

      if wireResponse, err = v.Element("xpath", ".//div[@id='main-div-xpath']"); err == nil {
        if webElement, err := wireResponse.WebElement(); err == nil {
          if webElement == nil || webElement.Value == "" {
            t.Error("could not find element", webElement.Value)
          }
        } else {
          t.Error("could not find element")
        }

      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestElements(t *testing.T) {

  var err error
  var wireResponse *WireResponse

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/elements.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Elements("class name", "myclass"); err == nil {

        var list []*WebElement
        if list, err = wireResponse.WebElements(); err == nil {

          if len(list) < 4 {
            t.Error("Should have found 4 elements")
          }
        } else {
          t.Error("could not find elements")
        }

      }

    }
  }

}
