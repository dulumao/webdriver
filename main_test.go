package webdriver

import (
  "fmt"
  "log"
  "os"
  // "strconv"
  "testing"
  "time"
)

var no_chrome bool
var no_firefox bool

var clientChrome Client
var sessionChrome *Session

var clientFirefox Client
var sessionFirefox *Session

var clients []Client
var sessions []*Session

////////////////////////////////////////////////////////////////
func sleepForSeconds(value int) {
  duration, _ := time.ParseDuration(fmt.Sprintf("%ds", value))
  log.Println("sleeping: ", duration, " seconds...")
  time.Sleep(duration)
}

////////////////////////////////////////////////////////////////
func startChrome() {

  clientChrome = &Chrome{
                      LogPath: "support/drivers/chromedriver.log",
                      PathExec: "support/drivers/chromedriver",
                      Verbose: true,
                    }

  if err := clientChrome.Run(); err == nil {

    clients = append(clients, clientChrome)

    sleepForSeconds(2)

    if sessionChrome, err = clientChrome.NewSession(); err != nil {
      log.Println("chrome", err)
    } else {
      sessions = append(sessions, sessionChrome)
      log.Println("added chrome to list:", len(sessions), sessionChrome)
    }

  }

}

////////////////////////////////////////////////////////////////
func startFirefox() {

  clientFirefox = &Firefox{
                      // ProfileDir: "/tmp/mywebdriver",
                      // UserJS: "user.js",
                      // RemoveOnClose: "none",
                      Extension: &Extension{
                        Path: "support/drivers/webdriver.xpi",
                      },
                    }

  if err := clientFirefox.Run(); err == nil {

    clients = append(clients, clientFirefox)

    sleepForSeconds(2)

    if sessionFirefox, err = clientFirefox.NewSession(); err != nil {
      log.Println("firefox", err)
    } else {
      sessions = append(sessions, sessionFirefox)
      log.Println("added firefox to list:", len(sessions), sessionFirefox)
    }

  }

}

////////////////////////////////////////////////////////////////
func TestMain(m *testing.M) {

  sessions = make([]*Session, 0)

  // setting the environment variable NOFIREFOX to anything
  // will set no_firefox = true
  if os.Getenv("NOCHROME") != "" {
    no_chrome = true
  } else {
    startChrome()
  }

no_firefox = true
  // if os.Getenv("NOFIREFOX") != "" {
  //   no_firefox = true
  // } else {
  //   startFirefox()
  // }

  log.Println("total number of sessions: ", len(sessions))

  go startWebServer()

  x := m.Run()

  if !no_chrome {
    clientChrome.Close()
  }

  if !no_firefox {
    clientFirefox.Close()
  }

  os.Exit(x)

}















