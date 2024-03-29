// This package is an attempt at a Golang implementation for the Selenium
// webdriver: http://www.seleniumhq.org/ and aims to fully suport the
// Json Wire Protocol https://code.google.com/p/selenium/wiki/JsonWireProtocol
package webdriver

type (

  // A webdriver client that supports the Json Wire Protocol.  All clients should implement this interface.
  Client interface {

    Close() error

    GetSessions() []*Session

    Run() error

    Session(capabilities ...*Capabilities) (session *Session, err error)

    SetDefaults() error

    Status() *Wire

    WireSessions() *Wire

  }

)
