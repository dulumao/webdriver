package webdriver

import (
  // "encoding/json"
  // "errors"
  "fmt"
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

// GET /session/:sessionId/element/:id/attribute/:name
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/attribute/:name
//
// Get the value of an element's attribute.
func (s *WebElement) AttributeName(name string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.GetRequest(s.BuildElementUrl(
                              fmt.Sprintf("/session/:sessionid/element/:id/attribute/%v", name)), nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// Returns true if WebElement.Value is an empty string.  You could simply check for an empty string
// in your code and call it a day.  However, Blank() encapsulates that logic into a method call as future
// versions of the code may change and this could reduce the amount of future changes in app code.
func (s *WebElement) Blank() bool {
  return s.Value == ""
}

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

// POST /session/:sessionId/element/:id/clear
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/clear
//
// Query for an element's tag name.
func (s *WebElement) Clear() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.PostRequest(s.BuildElementUrl("/session/:sessionid/element/:id/clear"),
                              nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// POST /session/:sessionId/element/:id/click
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/click
//
// Click on an element.
func (s *WebElement) Click() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.PostRequest(s.BuildElementUrl("/session/:sessionid/element/:id/click"),
                              nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// GET /session/:sessionId/element/:id/css/:propertyName
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/css/:propertyName
//
// Query the value of an element's computed CSS property. The CSS property to query should be specified using the CSS property name, not the JavaScript property name (e.g. background-color instead of backgroundColor).
func (s *WebElement) CssProperty(name string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.GetRequest(s.BuildElementUrl(
                              fmt.Sprintf("/session/:sessionid/element/:id/css/%v", name)), nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// GET /session/:sessionId/element/:id/displayed
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/displayed
//
// Determine if an element is currently displayed.
//
//     Returns:
//       {boolean} Whether the element is displayed.
func (s *WebElement) Displayed() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.GetRequest(s.BuildElementUrl("/session/:sessionid/element/:id/displayed"),
                              nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
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

// GET /session/:sessionId/element/:id/enabled
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/enabled
//
// Determine if an element is currently enabled.
//
//     Returns:
//       {boolean} Whether the element is enabled.
func (s *WebElement) Enabled() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.GetRequest(s.BuildElementUrl("/session/:sessionid/element/:id/enabled"),
                              nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// GET /session/:sessionId/element/:id/location
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/location
//
// Determine an element's location on the page. The point (0, 0) refers to the upper-left corner of the page. The element's coordinates are returned as a JSON object with x and y properties.
//
//     Returns:
//       {x:number, y:number} The X and Y coordinates for the element on the page.
func (s *WebElement) Location() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.GetRequest(s.BuildElementUrl("/session/:sessionid/element/:id/location"),
                              nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// GET /session/:sessionId/element/:id/location_in_view
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/location_in_view
//
// Determine an element's location on the screen once it has been scrolled into view.
//
//     Note:
//       This is considered an internal command and should only be used to determine an element's location for correctly generating native events.
//
//     Returns:
//       {x:number, y:number} The X and Y coordinates for the element on the page.
func (s *WebElement) LocationInView() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.GetRequest(s.BuildElementUrl("/session/:sessionid/element/:id/location_in_view"),
                              nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// GET /session/:sessionId/element/:id/name
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/name
//
// Query for an element's tag name.
func (s *WebElement) Name() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.GetRequest(s.BuildElementUrl("/session/:sessionid/element/:id/name"),
                              nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// POST /session/:sessionId/element/:id/value
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/value
//
// Send a sequence of key strokes to an element.
//
// NOTE:  The api was designed in such a way as to have each method name match the corresponding
// command or method name in the JsonWireProtocol.  However, the method name Value() would conflict
// with the field name Value on WebElement, so, I chose to leave the field name in tact and
// use the method name PostValue().  The reason is I had a large chunk of the api written and
// discovered this conflict while trying to add Value() to WebElement.  Bummer.
func (s *WebElement) PostValue(value []string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.PostRequest(s.BuildElementUrl("/session/:sessionid/element/:id/value"),
                              &Params{"value": value}); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// GET /session/:sessionId/element/:id/selected
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/selected
//
// Determine if an OPTION element, or an INPUT element of type checkbox or radiobutton is currently selected.
//
//     Returns:
//       {boolean} Whether the element is selected.
func (s *WebElement) Selected() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.GetRequest(s.BuildElementUrl("/session/:sessionid/element/:id/selected"),
                              nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// GET /session/:sessionId/element/:id/size
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/size
//
// Determine an element's size in pixels. The size will be returned as a JSON object with width and height properties.
//
//     Returns:
//       {width:number, height:number} The width and height of the element, in pixels.
func (s *WebElement) Size() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.GetRequest(s.BuildElementUrl("/session/:sessionid/element/:id/size"),
                              nil); err == nil {

    wireResponse, err = s.Session.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s.Session
    }

  }

  return wireResponse, err
}

// POST /session/:sessionId/element/:id/submit
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/:id/submit
//
// Submit a FORM element. The submit command may also be applied to any element that is a descendant of a FORM element.
func (s *WebElement) Submit() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.Session.PostRequest(s.BuildElementUrl("/session/:sessionid/element/:id/submit"),
                              nil); err == nil {

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










