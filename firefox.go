package webdriver

import (
  "errors"
  "fmt"
  "log"
  "io/ioutil"
  "net"
  "os"
  "os/exec"
  "os/user"
  "path/filepath"
  "strings"
  "time"
)

type (

  Firefox struct {
    // will probably always be localhost and could have just hard coded
    // it throughout the lib, however, using a struct
    // field in case future may require changing it
    DirPermissions os.FileMode
    Extension *Extension
    FilePermissions os.FileMode
    Host string
    LockingPort int
    PathExec string
    Port int
    Process *os.Process
    ProfileDir string
    RemoveOnClose string
    Timeout float64
    UserJS string
    UserJSPolicy string

    *Sessions

    *Wire

  }

)

func (s *Firefox) SetDefaults() (err error) {

  if s.DirPermissions == 0 {s.DirPermissions = 0770}
  if s.FilePermissions == 0 {s.FilePermissions = 0600}
  if s.Host == "" {s.Host = "localhost"}
  if s.LockingPort == 0 {s.LockingPort = 7054}
  if s.PathExec == "" {s.PathExec = "firefox"}
  if s.Port == 0 {s.Port = 7055}
  if s.Timeout == 0 {s.Timeout = 60}
  if s.UserJSPolicy == "" {s.UserJSPolicy = "merge"}

  log.Println("Firefox SetDefaults()")
  s.Sessions = &Sessions{}
  s.Sessions.SetDefaults()
  s.Wire = &Wire{}
  s.Wire.SetDefaults()

  if s.Extension == nil {
    s.Extension = &Extension{}
  }

  // a full path to the firefox extension file: webdriver.xpi is REQUIRED
  if s.Extension.Path == "" {
    err = errors.New("WARNING: You need to set the path to the location of the Selenium plugin extension.\nExample: client := &webdriver.Firefox{Extension: &webdriver.Extension{Path: \"~/webdriver.xpi\"}}")
    log.Println(err)
    return err
  }

  // user may have defined it with a tilde to point to their home directory
  // find out the absolute path to their home directory and search/replace the tilde with the absolute path.
  if strings.HasPrefix(s.Extension.Path, "~") {
    var currentUser *user.User

    currentUser, err = user.Current()

    s.Extension.Path = strings.Replace(s.Extension.Path, "~", currentUser.HomeDir, 1)
    log.Println("Extension.Path has been set to: ", s.Extension.Path)
  }

  if s.Extension.Policy == "" {s.Extension.Policy = "remove"}
  if s.Extension.ConfigPolicy == "" {s.Extension.ConfigPolicy = "remove"}
  if len(s.Extension.ConfigList) <= 0 {
    s.Extension.ConfigList = []string{
                                "compatibility.ini",
                                "extensions.cache",
                                "extensions.ini",
                                "extensions.json",
                                "extensions.rdf",
                                "extensions.sqlite",
                                "extensions.sqlite-journal",
    }
  }

  // if the user has set the ProfileDir, then, it is up to them to set
  // if the ProfileDir should be removed on close or not.
  // the default is to leave it blank if ProfileDir is set.  this will
  // prevent the ProfileDir from being deleted.
  if s.ProfileDir == "" {

    s.ProfileDir, err = ioutil.TempDir(os.TempDir(), "webdriver-firefox")

      // we are using a temporary profile and directory.  set to RemoveOnClose to "remove"
    // ONLY if it has NOT BEEN set already
    if s.RemoveOnClose == "" {
      s.RemoveOnClose = "remove"
    }

}

  s.Extension.BaseDir = filepath.Join(s.ProfileDir, "extensions")

  log.Println("        Profile directory: ", s.ProfileDir)
  log.Println("Extensions base directory: ", s.Extension.BaseDir)
  log.Println("        Extensions Policy: ", s.Extension.Policy)
  log.Println("  Extensions ConfigPolicy: ", s.Extension.ConfigPolicy)

  return err
}

func (s *Firefox) Run() (err error) {

  var listener net.Listener
  if err = s.SetDefaults(); err == nil {

    log.Println("Waiting for lock at: ", s.Host, s.LockingPort)
    if listener, err = waitForLock(s.Host, s.LockingPort, s.Timeout * float64(time.Second)); err == nil {

      defer listener.Close()

      log.Println("Lock succeeded!!")

      if s.Port, err = findNextAvailablePort(s.Host, s.Port, s.Timeout * float64(time.Second)); err == nil {

        log.Println("Firefox extension should listen at: ", s.Host, s.Port)

        if err = s.configProfile(); err == nil {

          cmd := exec.Command(s.PathExec,
                              "-silent",
                              "-no-remote",
                              "-profile",
                              s.ProfileDir)

          // run is suppose to wait until the process has finished
          log.Println("Starting firefox in -silent mode (as per the spec)")
          cmd.Run()

          cmd = exec.Command(s.PathExec,
                              "-no-remote",
                              "-foreground",
                              "-profile",
                              s.ProfileDir)

          log.Println("All systems Go!!  Starting Firefox... (for real this time)")
          cmd.Start()

          s.Process = cmd.Process

          log.Println("Firefox started pid: ", s.Process.Pid)

          log.Println("Waiting to connect to webdriver at: ", s.Host, s.Port)
          err = waitForConnect(s.Host, s.Port, s.Timeout * float64(time.Second))

          s.BaseUrl = fmt.Sprintf("http://%v:%v/hub", s.Host, s.Port)
          s.Sessions.BaseUrl = fmt.Sprintf("http://%v:%v/hub", s.Host, s.Port)

          log.Println("Look Mom, no hands!! Firefox should now be running with webdriver at: ", s.BaseUrl)
        }

      }

    }

  }

  return err
}

// Kills the currently running webdriver if it is running.
func (s *Firefox) Close() (err error) {

    log.Println("Close()")

  if s.Process != nil {
    log.Println("Close() killing process")
    s.Process.Kill()

    time.Sleep(2 * time.Second)
  }

  log.Println("RemoveOnClose: ", s.RemoveOnClose)

  if s.RemoveOnClose == "remove" {

    if _, err = os.Stat(s.ProfileDir); err == nil {

      log.Println("RemoveOnClose: YES ", s.RemoveOnClose, s.ProfileDir)
      err = os.RemoveAll(s.ProfileDir)

    }

  }

  return err
}


























