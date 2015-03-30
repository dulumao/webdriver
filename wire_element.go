package webdriver

import (
  "net/http"
)

// POST /session/:sessionId/element
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element
//
// Search for an element on the page, starting from the document root. The located element will be returned as a WebElement JSON object. The table below lists the locator strategies that each server should support. Each locator must return the first matching element located in the DOM.
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
func (s *Session) Element(using string, value string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/element",
                              &Params{"using": using, "value": value}); err == nil {

    wireResponse, err = s.Do(req)
    if wireResponse != nil {
      wireResponse.Session = s
    }

  }

  return wireResponse, err
}

// POST /session/:sessionId/elements
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/elements
//
// Search for multiple elements on the page, starting from the document root. The located elements will be returned as a WebElement JSON objects. The table below lists the locator strategies that each server should support. Elements should be returned in the order located in the DOM.
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
//      XPathLookupError - If using XPath and the input expression is invalid
func (s *Session) Elements(using string, value string) (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/elements",
                              &Params{"using": using, "value": value}); err == nil {

    wireResponse, err = s.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s
    }

  }

  return wireResponse, err
}

// POST  /session/:sessionId/element/active
//
// https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/element/active
//
// Get the element on the page that currently has focus. The element will be returned as a WebElement JSON object.
//
//    Returns:
//      {ELEMENT:string} A WebElement JSON object for the active element.
//
//    Potential Errors:
//      NoSuchWindow - If the currently selected window has been closed.
func (s *Session) Active() (wireResponse *WireResponse, err error) {

  var req *http.Request
  if req, err = s.PostRequest("/session/:sessionid/element/active", nil); err == nil {

    wireResponse, err = s.Do(req)

    if wireResponse != nil {
      wireResponse.Session = s
    }

  }

  return wireResponse, err
}




























