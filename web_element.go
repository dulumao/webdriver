package webdriver

import (
  // "fmt"
  "net/http"
  "strings"
)

type (

  // WebElement - An object in the WebDriver API that represents a DOM element on the page.
  //
  // WebElement JSON Object - The JSON representation of a WebElement for transmission over the wire.
  //     Key Type  Description
  //     ELEMENT string  The opaque ID assigned to the element by the server. This ID should be used in all subsequent commands issued against the element.
  WebElement struct {
    Session             *Session
    Value               string `json:"element"`
  }

)

// BuildElementUrl() is similar to BuildFullUrl(), except it simply does a search / replace
// on the :id value of the current WebElement.  Relies on the current value of Value as the element :id.
//
//   // given:
//     Value = "{my-hex-value-or-some-custom-id-value}"
//
//   // the following call
//   BuildElement("/session/:sessionid/element/:id/text")
//
//   // would produce
//   /session/:sessionid/element/{my-hex-value-or-some-custom-id-value}/text
func (s *WebElement) BuildElementUrl(url string) string {
  return strings.Replace(url, ":id", s.Value, -1)
}

// POST /session/:sessionId/element/:id/element
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/element
//
// Search for an element on the page, starting from the identified element. The located element will be returned as a WebElement JSON object. The table below lists the locator strategies that each server should support. Each locator must return the first matching element located in the DOM.
//
//    Strategy              Description
//      class name          Returns an element whose class name contains the search value; compound class names are not permitted.
//      css selector        Returns an element matching a CSS selector.
//      id                  Returns an element whose ID attribute matches the search value.
//      name                Returns an element whose NAME attribute matches the search value.
//      link text           Returns an anchor element whose visible text matches the search value.
//      partial link text   Returns an anchor element whose visible text partially matches the search value.
//      tag name            Returns an element whose tag name matches the search value.
//      xpath               Returns an element matching an XPath expression.
//
//    JSON Parameters:
//      using - {string} The locator strategy to use.
//      value - {string} The The search target.
//
//    Returns:
//      {ELEMENT:string} A WebElement JSON object for the located element.
//
//    Potential Errors:
//      NoSuchWindow - If the currently selected window has been closed.
//      NoSuchElement - If the element cannot be found.
//      XPathLookupError - If using XPath and the input expression is invalid
func (s *WebElement) Element(using string, value string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.PostRequest(s.BuildElementUrl("/session/:sessionid/element/:id/element"),
                              &Params{"using": using, "value": value}); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// POST /session/:sessionId/element/:id/elements
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/elements
//
// Search for multiple elements on the page, starting from the identified element. The located elements will be returned as a WebElement JSON objects. The table below lists the locator strategies that each server should support. Elements should be returned in the order located in the DOM.
//
//    Strategy              Description
//      class name          Returns an element whose class name contains the search value; compound class names are not permitted.
//      css selector        Returns an element matching a CSS selector.
//      id                  Returns an element whose ID attribute matches the search value.
//      name                Returns an element whose NAME attribute matches the search value.
//      link text           Returns an anchor element whose visible text matches the search value.
//      partial link text   Returns an anchor element whose visible text partially matches the search value.
//      tag name            Returns an element whose tag name matches the search value.
//      xpath               Returns an element matching an XPath expression.
//
//    JSON Parameters:
//      using - {string} The locator strategy to use.
//      value - {string} The The search target.
//
//    Returns:
//      {Array.<{ELEMENT:string}>} A list of WebElement JSON objects for the located elements.
//
//    Potential Errors:
//      NoSuchWindow - If the currently selected window has been closed.
//      NoSuchElement - If the element cannot be found.
//      XPathLookupError - If using XPath and the input expression is invalid
func (s *WebElement) Elements(using string, value string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.PostRequest(s.BuildElementUrl("/session/:sessionid/element/:id/elements"),
                              &Params{"using": using, "value": value}); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// GET /session/:sessionId/element/:id/text
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/text
//
// Returns the visible text for the element.
func (s *WebElement) Text() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.GetRequest(s.BuildElementUrl("/session/:sessionid/element/:id/text"),
                              nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}












