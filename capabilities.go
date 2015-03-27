package webdriver

type (

  // ActualCapabilities represents the json object returned from the
  // webdriver during a call to create a new session.  it contains all
  // of the actual capabilities of the current environment
  ActualCapabilities struct {
    AcceptSslCerts                bool `json:"acceptSslCerts"`
    ApplicationCacheEnabled       bool `json:"applicationCacheEnabled"`
    BrowserName                 string `json:"browserName"`
    BrowserConnectionEnabled      bool `json:"browserConnectionEnabled"`
    CssSelectorsEnabled           bool `json:"cssSelectorsEnabled"`
    DatabaseEnabled               bool `json:"databaseEnabled"`
    LocationContextEnabled        bool `json:"locationContextEnabled"`
    HandlesAlerts                 bool `json:"handlesAlerts"`
    JavascriptEnabled             bool `json:"javascriptEnabled"`
    NativeEvents                  bool `json:"nativeEvents"`
    Platform                    string `json:"platform"`
    Rotatable                     bool `json:"rotatable"`
    TakesScreenshot               bool `json:"takesScreenshot"`
    Version                     string `json:"version"`
    WebStorageEnabled             bool `json:"webStorageEnabled"`

    // TODO: there is a slot in the spec for a Proxy JSON object,
    // however, it doesn't seem to be supported right now

  }

  // TODO: there are a ton more possibilities defined on the spec
  // however, the doc itself states:
  // https://code.google.com/p/selenium/wiki/DesiredCapabilities
  // DesiredCapabilities  
  // (still under work) A specification of DesiredCapabilities
  // and their content when used in JsonWireProtocol Updated Feb 26, 2015

  // for now, decided to implement as a simple map instead of a struct
  // as empty structs would cause the webdriver extension to crash.
  Capabilities map[string]interface{}

  // Capabilities struct {
  //   AcceptSslCerts                bool `json:"acceptSslCerts"`
  //   ApplicationCacheEnabled       bool `json:"applicationCacheEnabled"`
  //   BrowserName                 string `json:"browserName"`
  //   BrowserConnectionEnabled      bool `json:"browserConnectionEnabled"`
  //   CssSelectorsEnabled           bool `json:"cssSelectorsEnabled"`
  //   DatabaseEnabled               bool `json:"databaseEnabled"`
  //   ElementScrollBehavior          int `json:"elementScrollBehavior"`
  //   LocationContextEnabled        bool `json:"locationContextEnabled"`
  //   HandlesAlerts                 bool `json:"handlesAlerts"`
  //   JavascriptEnabled             bool `json:"javascriptEnabled"`
  //   NativeEvents                  bool `json:"nativeEvents"`
  //   Platform                    string `json:"platform"`
  //   Rotatable                     bool `json:"rotatable"`
  //   TakesScreenshot               bool `json:"takesScreenshot"`
  //   UnexpectedAlertBehaviour    string `json:"unexpectedAlertBehaviour"`
  //   Version                     string `json:"version"`
  //   WebStorageEnabled             bool `json:"webStorageEnabled"`

  //  // TODO: there is a slot in the spec for a Proxy JSON object,
  //  // however, it doesn't seem to be supported right now

  // }

)
