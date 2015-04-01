package webdriver

import (
  "fmt"
  "log"
  "os"
  "testing"
  "time"
)

var no_chrome bool
var no_firefox bool
var env_circle bool
var env_volatile bool

var clients []Client
var sessions map[string]*Session

////////////////////////////////////////////////////////////////
func getSessions(keys... string) (list []*Session) {

  if len(keys) > 0 {

    for _, v := range keys {
      if value, ok := sessions[v]; ok {
        list = append(list, value)
      }
    }

  } else {
    for _, v := range sessions {
      list = append(list, v)
    }
  }

  return list
}

////////////////////////////////////////////////////////////////
func sleepForSeconds(value int) {
  duration, _ := time.ParseDuration(fmt.Sprintf("%ds", value))
  log.Println("sleeping: ", duration, " seconds...")
  time.Sleep(duration)
}

////////////////////////////////////////////////////////////////
func startChrome() {

  if os.Getenv("NOCHROME") != "" {
    no_chrome = true
  } else {

    client := &Chrome{
                      LogPath: "support/drivers/chromedriver.log",
                      PathExec: "support/drivers/chromedriver",
                      Verbose: true,
                    }

    if err := client.Run(); err == nil {

      clients = append(clients, client)

      sleepForSeconds(2)

      if session, err := client.Session(); err == nil {

        // rigging the tests for now
        // if !env_circle {
          sessions["chrome"] = session
          log.Println("added chrome to session list")
        // }

      } else {
        log.Println("cannot establish chrome session: ", err)
      }

    }

  }

}

////////////////////////////////////////////////////////////////
func startFirefox() {

  if os.Getenv("NOFIREFOX") != "" {
    no_firefox = true
  } else {

    client := &Firefox{
                      // ProfileDir: "/tmp/mywebdriver",
                      // UserJS: "user.js",
                      // RemoveOnClose: "none",
                      Extension: &Extension{
                        Path: "support/drivers/webdriver.xpi",
                      },
                    }

    if err := client.Run(); err == nil {

      clients = append(clients, client)

      sleepForSeconds(2)

      if session, err := client.Session(); err == nil {

        sessions["firefox"] = session
        log.Println("added firefox to session list")

      } else {
        log.Println("cannot establish firefox session: ", err)
      }


    }
  }

}

////////////////////////////////////////////////////////////////
func TestMain(m *testing.M) {

  sessions = make(map[string]*Session, 0)

  if os.Getenv("ENV_CIRCLECI") != "" {
    env_circle = true
  }

  // tests that use the /root directory are volatile, because, they attempt to remove entire
  // directory structures using removeall().  you have to set the environment variable VOLATILE=true
  // to get these tests to run.  otherwise, they are skipped.
  // the reason for using the /root directory is permissions.  parts of the code for installing
  // firefox will remove entire directories.  i'm using the /root to test that code, because, the
  // default user would not have permissions to nuke a directory owned by root.
  // I plan to develop a better solution using a mock file system.  for now, just using something quick and easy.
  if os.Getenv("VOLATILE") != "" {
    env_volatile = true
  }

  startChrome()

  startFirefox()

  go startWebServer()

  log.Println("Active sessions: ", len(sessions))

  x := m.Run()

  for _, v := range clients {
    v.Close()
  }

  os.Exit(x)

}















