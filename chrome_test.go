package webdriver

import (
  "testing"
)

////////////////////////////////////////////////////////////////
func TestChromeSetDefaults(t *testing.T) {

  s := &Chrome{}

  if s.Host != "" {
    t.Error("Host should be blank by default", s.Host)
  }

  if s.Port > 0 {
    t.Error("Port should be zero", s.Port)
  }

  if s.PathExec != "" {
    t.Error("PathExec should be blank by default", s.PathExec)
  }

  if s.Timeout > 0 {
    t.Error("Timeout should be zero", s.Timeout)
  }

  s.SetDefaults()

  if s.Host != "localhost" {
    t.Error("Host should equal localhost", s.Host)
  }

  if s.Port <= 0 {
    t.Error("Port should be 9515", s.Port)
  }

  if s.PathExec != "chromedriver" {
    t.Error("PathExec should equal chromedriver", s.PathExec)
  }

  if s.Timeout <= 0 {
    t.Error("Timeout should be 60", s.Timeout)
  }

}

////////////////////////////////////////////////////////////////
func TestBuildOptions(t *testing.T) {

  s := &Chrome{
    AdbPort: 1000,
    LogPath: "/tmp/chromedriver.log",
    Port: 9515,
    PortServer: "localhost",
    Silent: true,
    UrlBase: "/hub",
    Verbose: true,
    WhiteList: "127.0.0.1, 127.0.0.2",
  }

  options := s.buildOptions()
  if options == nil {
    t.Error("wtf")
  }

  // TODO: validate the options
}

////////////////////////////////////////////////////////////////
// starts another chrome webdriver on another port, then, shuts it down
// to call Close() and Delete() on All of it's sessions.
func TestCreateDestroyChromeSessionWithCapabilities(t *testing.T) {

  if !no_chrome && !env_circle {
    client := &Chrome{
                        LogPath: "support/drivers/chromedriver.log",
                        PathExec: "support/drivers/chromedriver",
                        Verbose: true,
                        Port: 9516,
                      }

    if err := client.Run(); err == nil {

      if session, err := client.Session(&Capabilities{"Platform": "Linux"}, &Capabilities{"Platform": "Linux"}); err == nil {

        if len(client.GetSessions()) <= 0 {
          t.Error("client should have at least one session")
        }

        if wire := session.DeleteSession(); wire.Error != nil {
          t.Error("could not delete session: ", wire.Error)
        }

      } else {
        t.Error(err)
      }

      client.Close()

    }
  }

}















