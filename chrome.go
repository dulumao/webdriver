package webdriver

import (
  "fmt"
  "os"
  "os/exec"
  "time"
)

type (

  Chrome struct {
    AdbPort int
    Host string
    LogPath string
    PathExec string
    Port int
    PortServer string
    Process *os.Process
    RemoveOnClose string
    Silent bool
    Timeout float64
    UrlBase string
    Verbose bool
    WhiteList string

    *Sessions

    *Wire

  }

)

// Sets the default values for a Client.  Chrome struct includes
// an anonymous Client struct that is initialized in this method.
func (s *Chrome) SetDefaults() (err error) {
  if s.Host == "" {s.Host = "localhost"}
  if s.PathExec == "" {s.PathExec = "chromedriver"}
  if s.Port == 0 {s.Port = 9515}
  if s.Timeout == 0 {s.Timeout = 60}

  s.Sessions = &Sessions{}
  s.Sessions.SetDefaults()
  s.Wire = &Wire{}
  s.Wire.SetDefaults()

  return err
}

// ChromeDriver 2.14.313457 (3d645c400edf2e2c500566c9aa096063e707c9cf)
//
// Options
//   --port=PORT                     port to listen on
//   --adb-port=PORT                 adb server port
//   --log-path=FILE                 write server log to file instead of stderr, increases log level to INFO
//   --verbose                       log verbosely
//   --version                       print the version number and exit
//   --silent                        log nothing
//   --url-base                      base URL path prefix for commands, e.g. wd/url
//   --port-server                   address of server to contact for reserving a port
//   --whitelisted-ips               comma-separated whitelist of remote IPv4 addresses which are allowed to connect to ChromeDriver
//
func (s *Chrome) Run() (err error) {

  if err = s.SetDefaults(); err == nil {

    options := s.buildOptions()

    cmd := exec.Command(s.PathExec, options...)

    cmd.Start()

    s.Process = cmd.Process

    // shouldn't make a difference if UrlBase is blank or not
    s.BaseUrl = fmt.Sprintf("http://%v:%v%v", s.Host, s.Port, s.UrlBase)
    s.Sessions.BaseUrl = fmt.Sprintf("http://%v:%v%v", s.Host, s.Port, s.UrlBase)

    err = waitForConnect(s.Host, s.Port, s.Timeout * float64(time.Second))

  }

  return err
}

func (s *Chrome) buildOptions() (options []string) {

  if s.AdbPort > 0 {
    options = append(options, fmt.Sprintf("-adb-port=%d", s.AdbPort))
  }

  if s.LogPath != "" {
    options = append(options, fmt.Sprintf("-log-path=%v", s.LogPath))
  }

  if s.Port > 0 {
    options = append(options, fmt.Sprintf("-port=%d", s.Port))
  }

  if s.PortServer != "" {
    options = append(options, fmt.Sprintf("-port-server=%v", s.PortServer))
  }

  if s.Silent {
    options = append(options, "--silent")
  }

  if s.UrlBase != "" {
    options = append(options, fmt.Sprintf("-url-base=%v", s.UrlBase))
  }

  if s.Verbose {
    options = append(options, "--verbose")
  }

  if s.WhiteList != "" {
    options = append(options, fmt.Sprintf("-whitelisted-ips=%v", s.WhiteList))
  }

  return options
}

// TODO Figure out how to shutdown the browser as well as the webdriver
// currently, it seems like just the webdriver is being shutdown
// Kills the currently running webdriver if it is running.
func (s *Chrome) Close() (err error) {

  if s.Process != nil {

    for _, v := range s.Sessions.List {
      v.Delete()
    }

    s.Process.Kill()
  }

  return err
}

func (s *Chrome) GetSessions() (*Sessions) {
  return s.Sessions
}
























