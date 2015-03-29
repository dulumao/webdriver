package webdriver

import (
  // "fmt"
  // "log"
  // "encoding/json"
  "strings"
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
  var wireSubResponse2 *WireResponse

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/elements.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Elements("class name", "myclass"); err == nil {

        var list []*WebElement
        if list, err = wireResponse.WebElements(); err == nil {

          if len(list) < 4 {
            t.Error("Should have found 4 elements")
          } else {

            all_text := "my class 1_my class 2_my class 3_my class 4_"

            for _, v := range list {

              if wireSubResponse2, err = v.Text(); err == nil {
                if !strings.Contains(all_text, wireSubResponse2.StringValue()) {
                  t.Error("text not found in all_text variable", wireSubResponse2.StringValue())
                }
              }
            }

          }

        } else {
          t.Error("could not find elements")
        }

      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestSubElement(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  var wireSubResponse *WireResponse
  var wireSubResponse2 *WireResponse
  var webElement *WebElement
  var subElement *WebElement

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/sub-element.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("id", "main-div-id"); err == nil {
        if webElement, err = wireResponse.WebElement(); err == nil {
          if webElement == nil || webElement.Value == "" {
            t.Error("could not find element", webElement.Value)
          } else {

            if wireSubResponse, err = webElement.Element("id", "div-sub"); err == nil {
              if subElement, err = wireSubResponse.WebElement(); err == nil {
                if wireSubResponse2, err = subElement.Text(); err == nil {
                  if wireSubResponse2.StringValue() != "main div sub" {
                    t.Error("text should be main div sub", wireSubResponse2.StringValue())
                  }
                } else {
                  t.Error("could not find element")
                }

              } else {
                t.Error("could not find element")
              }

            } else {
              t.Error("could not find element")
            }

          }
        } else {
          t.Error("could not find element")
        }

      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestSubElements(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  var wireSubResponse *WireResponse
  var wireSubResponse2 *WireResponse
  var webElement *WebElement

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/sub-elements.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("id", "main-div-id"); err == nil {
        if webElement, err = wireResponse.WebElement(); err == nil {
          if webElement == nil || webElement.Value == "" {
            t.Error("could not find element", webElement.Value)
          } else {

            if wireSubResponse, err = webElement.Elements("class name", "myclass"); err == nil {

              var list []*WebElement
              if list, err = wireSubResponse.WebElements(); err == nil {
                if len(list) < 4 {
                  t.Error("Should have found 4 elements")
                } else {

                  all_text := "my class 1_my class 2_my class 3_my class 4_"

                  for _, v := range list {

                    if wireSubResponse2, err = v.Text(); err == nil {
                      if !strings.Contains(all_text, wireSubResponse2.StringValue()) {
                        t.Error("text not found in all_text variable", wireSubResponse2.StringValue())
                      }
                    }
                  }
                }

              }

            } else {
              t.Error("could not find element")
            }

          }
        } else {
          t.Error("could not find element")
        }

      }

    }
  }

}
