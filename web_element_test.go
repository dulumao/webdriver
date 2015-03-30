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
  var webElement *WebElement

  keys := []string{"class name", "id", "name", "link text", "partial link text", "tag name", "xpath"}
  text_values := []string{"main class", "main id", "main name", "link text", "partial link text", "my tag", "main id xpath"}
  values := []string{"main-div-class", "main-div-id", "main-div-name", "link text", "partial", "mytag", ".//div[@id='main-div-xpath']"}

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/element.html"); err == nil {

      sleepForSeconds(1)

      for i, key := range keys {

        if wireResponse, err = v.Element(key, values[i]); err == nil && wireResponse.Success() {
          if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {
              if wireResponse, err = webElement.Text(); err == nil && wireResponse.Success() {
                if wireResponse.StringValue() != text_values[i] {
                  t.Error("StringValue does not match", wireResponse.StringValue())
                }

              } else {
                t.Error(err, wireResponse.HttpStatusCode)
              }

          } else {
            t.Error(err, webElement)
          }

        } else {
          t.Error(err, wireResponse.HttpStatusCode)
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
        if list, err = wireResponse.WebElements(); err == nil && len(list) >= 4 {

          all_text := "my class 1_my class 2_my class 3_my class 4_"

          for _, v := range list {

            if wireSubResponse2, err = v.Text(); err == nil {
              if !strings.Contains(all_text, wireSubResponse2.StringValue()) {
                t.Error("text not found in all_text variable", wireSubResponse2.StringValue())
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
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {
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
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {
          if wireSubResponse, err = webElement.Elements("class name", "myclass"); err == nil {

            var list []*WebElement
            if list, err = wireSubResponse.WebElements(); err == nil && len(list) >= 4 {

              all_text := "my class 1_my class 2_my class 3_my class 4_"

              for _, v := range list {
                if wireSubResponse2, err = v.Text(); err == nil {
                  if !strings.Contains(all_text, wireSubResponse2.StringValue()) {
                    t.Error("text not found in all_text variable", wireSubResponse2.StringValue())
                  }
                }
              }

            } else {
              t.Error("should have produced at least four elements: ", len(list))
            }

          } else {
            t.Error("could not find element")
          }

        } else {
          t.Error("could not find element")
        }

      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestElementName(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  // var wireSubResponse *WireResponse
  // var wireSubResponse2 *WireResponse
  var webElement *WebElement

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/sub-elements.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("id", "main-div-id"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {
          if wireResponse, err = webElement.Name(); err == nil && wireResponse.Success() {
            if wireResponse.StringValue() != "div" {
              t.Error("StringValue() should equal div: ", wireResponse.StringValue())
            }
          } else {
            t.Error(err, wireResponse)
          }
        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestElementAttributeNameClear(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  // var wireSubResponse *WireResponse
  // var wireSubResponse2 *WireResponse
  var webElement *WebElement

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("id", "clear-me"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {
          if wireResponse, err = webElement.AttributeName("value"); err == nil && wireResponse.Success() {
            if wireResponse.StringValue() != "clear me" {
              t.Error("StringValue() should equal clear me: ", wireResponse.StringValue())
            } else {
              if wireResponse, err = webElement.Clear(); err == nil && wireResponse.Success() {
                if wireResponse, err = webElement.AttributeName("value"); err == nil && wireResponse.Success() {
                  if wireResponse.StringValue() != "" {
                    t.Error("StringValue() should equal clear me: ", wireResponse.StringValue())
                  }
                }

              }

            }

          } else {
            t.Error(err, wireResponse)
          }
        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}


////////////////////////////////////////////////////////////////
func TestPostValue(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  // var wireSubResponse *WireResponse
  // var wireSubResponse2 *WireResponse
  var webElement *WebElement

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("id", "first-name"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

          keys := []string{"t", "e", "s", "t"}

          if wireResponse, err = webElement.PostValue(keys); err == nil && wireResponse.Success() {

            sleepForSeconds(1)

            if wireResponse, err = webElement.AttributeName("value"); err == nil && wireResponse.Success() {
              if wireResponse.StringValue() != "test" {
                t.Error("StringValue() should be test: ", wireResponse.StringValue())
              }
            } else {
              t.Error(err, wireResponse)
            }

          } else {
            t.Error(err, wireResponse)
          }
        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestActive(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  var webElement *WebElement

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("id", "first-name"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

          if wireResponse, err = webElement.Click(); err == nil && wireResponse.Success() {

            sleepForSeconds(1)

            if wireResponse, err = v.Active(); err == nil && wireResponse.Success() {
              if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

                if wireResponse, err = webElement.AttributeName("id"); err == nil && wireResponse.Success() {
                  if wireResponse.StringValue() != "first-name" {
                    t.Error("StringValue() should be first-name: ", wireResponse.StringValue())
                  }
                } else {
                  t.Error(err, wireResponse)
                }

              } else {
                t.Error(err, webElement)
              }

            } else {
              t.Error(err, wireResponse)
            }

          } else {
            t.Error(err, wireResponse)
          }
        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestSelected(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  var webElement *WebElement

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("id", "checked-yes"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

          if wireResponse, err = webElement.Selected(); err == nil && wireResponse.Success() {

            if wireResponse.StringValue() != "true" {
              t.Error("StringValue() should be true: ", wireResponse.StringValue())
            }

          } else {
            t.Error(err, wireResponse)
          }

        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

      if wireResponse, err = v.Element("id", "checked-no"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

          if wireResponse, err = webElement.Selected(); err == nil && wireResponse.Success() {

            if wireResponse.StringValue() != "false" {
              t.Error("StringValue() should be false: ", wireResponse.StringValue())
            }

          } else {
            t.Error(err, wireResponse)
          }

        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestEnabled(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  var webElement *WebElement

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("id", "first-name"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

          if wireResponse, err = webElement.Enabled(); err == nil && wireResponse.Success() {

            if wireResponse.StringValue() != "true" {
              t.Error("StringValue() should be true: ", wireResponse.StringValue())
            }

          } else {
            t.Error(err, wireResponse)
          }

        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

      if wireResponse, err = v.Element("id", "protected"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

          if wireResponse, err = webElement.Enabled(); err == nil && wireResponse.Success() {

            if wireResponse.StringValue() != "false" {
              t.Error("StringValue() should be false: ", wireResponse.StringValue())
            }

          } else {
            t.Error(err, wireResponse)
          }

        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestDisplayed(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  var webElement *WebElement

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("id", "first-name"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

          if wireResponse, err = webElement.Displayed(); err == nil && wireResponse.Success() {

            if wireResponse.StringValue() != "true" {
              t.Error("StringValue() should be true: ", wireResponse.StringValue())
            }

          } else {
            t.Error(err, wireResponse)
          }

        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

      if wireResponse, err = v.Element("id", "hidden"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

          if wireResponse, err = webElement.Displayed(); err == nil && wireResponse.Success() {

            if wireResponse.StringValue() != "false" {
              t.Error("StringValue() should be false: ", wireResponse.StringValue())
            }

          } else {
            t.Error(err, wireResponse)
          }

        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}

////////////////////////////////////////////////////////////////
func TestLocation(t *testing.T) {

  var err error
  var wireResponse *WireResponse
  var webElement *WebElement
  var point *Point

  for _, v := range sessions {
    if _, err = v.Url("http://localhost:8080/form01.html"); err == nil {

      sleepForSeconds(1)

      if wireResponse, err = v.Element("id", "first-name"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

          if wireResponse, err = webElement.Location(); err == nil && wireResponse.Success() {
            if point, err = wireResponse.Point(); err != nil {
              t.Error(err, point)
            } else {
              t.Log("location", point)
            }

          } else {
            t.Error(err, wireResponse)
          }

        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

      if wireResponse, err = v.Element("id", "last-name"); err == nil && wireResponse.Success() {
        if webElement, err = wireResponse.WebElement(); err == nil && !webElement.Blank() {

          if wireResponse, err = webElement.LocationInView(); err == nil && wireResponse.Success() {
            if point, err = wireResponse.Point(); err != nil {
              t.Error(err, point)
            } else {
              t.Log("location in view", point)
            }

          } else {
            t.Error(err, wireResponse)
          }

        } else {
          t.Error(err, webElement)
        }
      } else {
        t.Error(err, wireResponse)
      }

    }
  }

}
