package webdriver

import (
  // "encoding/json"
  // "fmt"
  "net"
  "net/http"
  "net/http/httptest"
  "net/url"
  "strconv"
  "strings"
  "testing"
)

// ////////////////////////////////////////////////////////////////
// TODO: revisit this later.  for now, testing the timeouts working
// seems a bit trivial opposed to the need to complete the development
// quickly.  don't want to wait for the timeout during testing right now.
func TestWaitForConnectSuccess(t *testing.T) {

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  }))

  defer ts.Close()

  url, _ := url.Parse(ts.URL)
  pair := strings.Split(url.Host, ":")
  port, _ := strconv.Atoi(pair[1])

  if err := waitForConnect(pair[0], port, 1); err != nil {
    t.Error("Unable to waitForConnect", err)
  }

}

////////////////////////////////////////////////////////////////
func TestWaitForConnectFailure(t *testing.T) {

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  }))

  defer ts.Close()

  if err := waitForConnect("localhost", 1, 0); err == nil {
    t.Error("Unable to waitForConnect", err)
  }

}

////////////////////////////////////////////////////////////////
func TestWaitForLockSuccess(t *testing.T) {

  if listener, err := waitForLock("localhost", 4444, 10); err != nil {
    t.Error("Unable to waitForLock", err)
  } else {
    listener.Close()
  }

}

////////////////////////////////////////////////////////////////
func TestWaitForLockFailure(t *testing.T) {

  if listener, err := net.Listen("tcp", "localhost:4444"); err != nil {
    t.Error(err)
  } else {
    defer listener.Close()
  }

  if _, err := waitForLock("localhost", 4444, 1); err == nil {
    t.Error("This should have been an error waitForLock", err)
  } else {
    // listener2.Close()
  }

}

////////////////////////////////////////////////////////////////
func TestFindNextAvailablePortSuccess(t *testing.T) {

  if port, err := findNextAvailablePort("localhost", 44444, 1); err != nil {
    t.Error("Unable to findNextAvailablePort", err)
  } else if port != 44444 {
    t.Error("Port should have been 44444 ", port)
  }

}

////////////////////////////////////////////////////////////////
func TestFindNextAvailablePortFailure(t *testing.T) {

  if _, err := findNextAvailablePort("localhost", 0, 1); err == nil {
    t.Error("Unable to findNextAvailablePort", err)
  }

}

////////////////////////////////////////////////////////////////
func TestFindNextAvailablePortTimeout(t *testing.T) {

  if listener, err := net.Listen("tcp", "localhost:44444"); err != nil {
    t.Error(err)
  } else {
    defer listener.Close()
  }

  if _, err := findNextAvailablePort("localhost", 44444, 0); err == nil {
    t.Error("Unable to findNextAvailablePort", err)
  }

}



























