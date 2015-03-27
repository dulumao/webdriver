package webdriver

import (
  "fmt"
  "os"
  "strconv"
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
  duration, _ := time.ParseDuration(strconv.Itoa(value))
  fmt.Println("sleeping: ", value, " seconds...")
  time.Sleep(duration * time.Second)
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

    sessions = make([]*Session, 0)

    sleepForSeconds(2)

    if sessionChrome, err = clientChrome.NewSession(); err != nil {
      fmt.Println("wtf?", err)
    } else {
      sessions = append(sessions, sessionChrome)
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

    sessions = make([]*Session, 0)

    sleepForSeconds(2)

    if sessionFirefox, err = clientFirefox.NewSession(); err != nil {
      fmt.Println("wtf?", err)
    } else {
      sessions = append(sessions, sessionFirefox)
    }

  }

}

////////////////////////////////////////////////////////////////
func TestMain(m *testing.M) {

  // setting the environment variable NOFIREFOX to anything
  // will set no_firefox = true
  if os.Getenv("NOCHROME") != "" {
    no_chrome = true
  } else {
    startChrome()
  }

  if os.Getenv("NOFIREFOX") != "" {
    no_firefox = true
  } else {
    startFirefox()
  }

  x := m.Run()

  if !no_chrome {
    clientChrome.Close()
  }

  if !no_firefox {
    clientFirefox.Close()
  }

  os.Exit(x)

}















