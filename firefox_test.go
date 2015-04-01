package webdriver

import (
  "os"
  "path/filepath"
  "testing"
)

////////////////////////////////////////////////////////////////
func TestFirefoxSetDefaults(t *testing.T) {

  s := &Firefox{}

  if s.DirPermissions > 0 {
    t.Error("DirPermissions should be zero", s.DirPermissions)
  }

  if s.FilePermissions > 0 {
    t.Error("FilePermissions should be zero", s.FilePermissions)
  }

  if s.Host != "" {
    t.Error("Host should be blank by default", s.Host)
  }

  if s.LockingPort > 0 {
    t.Error("LockingPort should be zero", s.LockingPort)
  }

  if s.PathExec != "" {
    t.Error("PathExec should be blank by default", s.PathExec)
  }

  if s.Port > 0 {
    t.Error("Port should be zero", s.Port)
  }

  if s.Timeout > 0 {
    t.Error("Timeout should be zero", s.Timeout)
  }

  s.SetDefaults()

  if s.DirPermissions != 0770 {
    t.Error("DirPermissions should equal 0770", s.DirPermissions)
  }

  if s.FilePermissions != 0600 {
    t.Error("FilePermissions should equal 0770", s.FilePermissions)
  }

  if s.Host != "localhost" {
    t.Error("Host should equal localhost", s.Host)
  }

  if s.LockingPort <= 0 {
    t.Error("LockingPort should be 7054", s.LockingPort)
  }

  if s.PathExec != "firefox" {
    t.Error("PathExec should equal firefox", s.PathExec)
  }

  if s.Port <= 0 {
    t.Error("Port should be 9515", s.Port)
  }

  if s.Timeout <= 0 {
    t.Error("Timeout should be 60", s.Timeout)
  }

}

////////////////////////////////////////////////////////////////
func TestFirefoxExtensionPathToHomeDir(t *testing.T) {

  s := &Firefox{
                Extension: &Extension{
                  Path: "~/webdriver.xpi",
                },
              }

  s.SetDefaults()

}

////////////////////////////////////////////////////////////////
// starts another firefox webdriver on another port, then, shuts it down
// to call Close() and Delete() on All of it's sessions.
func TestCreateDestroyFirefoxSession(t *testing.T) {

  if !no_firefox {
    client := &Firefox{
                        Extension: &Extension{
                          Path: "support/drivers/webdriver.xpi",
                        },
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

////////////////////////////////////////////////////////////////
func TestVerifyExtensionsPathCreate(t *testing.T) {

  client := &Firefox{
                      ProfileDir: "/tmp/user01",
                      Extension: &Extension{
                        Path: "support/drivers/webdriver.xpi",
                      },
                    }

  client.SetDefaults()

  if err := client.verifyExtensionsPath(); err != nil {
    t.Error(err)
  }

  if err := os.RemoveAll("/tmp/user01"); err != nil {
    t.Error(err)
  }

}

////////////////////////////////////////////////////////////////
func TestVerifyExtensionsPathAlreadyExists(t *testing.T) {

  if err := os.MkdirAll("/tmp/user01/extensions", 0770); err != nil {
    t.Error(err)
  }

  client := &Firefox{
                      ProfileDir: "/tmp/user01",
                      Extension: &Extension{
                        Path: "support/drivers/webdriver.xpi",
                      },
                    }

  client.SetDefaults()

  if err := client.verifyExtensionsPath(); err != nil {
    t.Error(err)
  }

  if err := os.RemoveAll("/tmp/user01"); err != nil {
    t.Error(err)
  }

}

////////////////////////////////////////////////////////////////
func TestVerifyExtensionsPathPermissionsFailure(t *testing.T) {

  if env_volatile {
    client := &Firefox{
                        ProfileDir: "/root/user01",
                        Extension: &Extension{
                          Path: "support/drivers/webdriver.xpi",
                        },
                      }

    client.SetDefaults()

    if err := client.verifyExtensionsPath(); err == nil {
      t.Error("this should have failed due to permissions.  should not write to /root directory")
    }
  } else {
    t.Skip("Skipping VOLATILE test...")
  }

}

////////////////////////////////////////////////////////////////
func TestExtractExtensionContentsFailure(t *testing.T) {

  if env_volatile {
    client := &Firefox{
                        ProfileDir: "/root/user01",
                        Extension: &Extension{
                          Path: "support/drivers/webdriver.xpi",
                        },
                      }

    client.SetDefaults()

    client.Extension.TargetPath = "/root/user01/extensions"

    if err := client.extractExtensionContents(); err == nil {
      t.Error("this should have failed due to permissions.  should not write to /root directory")
    }
  } else {
    t.Skip("Skipping VOLATILE test...")
  }

}

////////////////////////////////////////////////////////////////
func TestPrepareExtensionTargetRemoveAllFailure(t *testing.T) {

  if env_volatile {
    client := &Firefox{
                        ProfileDir: "/root/user01",
                        Extension: &Extension{
                          Path: "support/drivers/webdriver.xpi",
                        },
                      }

    client.SetDefaults()

    // client.Extension.TargetPath = "/root"

    if err := client.prepareExtensionTarget(); err == nil {
      t.Error("this should have failed due to permissions.  should not remove to /root directory")
    }
  } else {
    t.Skip("Skipping VOLATILE test...")
  }

}

////////////////////////////////////////////////////////////////
func TestPrepareExtensionTargetSuccess(t *testing.T) {

  if err := os.MkdirAll("/tmp/user01/extensions/myextension", 0770); err != nil {
    t.Error(err)
  }

  client := &Firefox{
                      ProfileDir: "/tmp/user01",
                      Extension: &Extension{
                        Path: "support/drivers/webdriver.xpi",
                      },
                    }

  client.SetDefaults()

  client.Extension.Name = "myextension"

  if err := client.prepareExtensionTarget(); err != nil {
    t.Error(err)
  }

  if err := os.RemoveAll("/tmp/user01"); err != nil {
    t.Error(err)
  }

}

////////////////////////////////////////////////////////////////
func TestRemoveExtensionConfig(t *testing.T) {

  if err := os.MkdirAll("/tmp/user01", 0770); err != nil {
    t.Error(err)
  }

  client := &Firefox{
                      ProfileDir: "/tmp/user01",
                      Extension: &Extension{
                        Path: "support/drivers/webdriver.xpi",
                      },
                    }

  client.SetDefaults()

  for _, v := range client.Extension.ConfigList {
    targetFileSpec := filepath.Join(client.ProfileDir, v)
    if f, ferr := os.Create(targetFileSpec); ferr != nil {
      t.Error(ferr)
    } else {
      f.Close()
    }
  }

  // client.Extension.Name = "myextension"

  if err := client.removeExtensionConfig(); err != nil {
    t.Error(err)
  }

  for _, v := range client.Extension.ConfigList {

    targetFileSpec := filepath.Join(client.ProfileDir, v)

    if _, err2 := os.Stat(targetFileSpec); err2 == nil {
      t.Error("file should not exist.")
    }

  }

  if err := os.RemoveAll("/tmp/user01"); err != nil {
    t.Error(err)
  }

}














