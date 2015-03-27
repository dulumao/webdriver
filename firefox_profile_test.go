package webdriver

import (
  // "fmt"
  // "os"
  // "os/exec"
  "testing"
)

//////////////////////////////////////////////////////////////
func TestParseUserJS(t *testing.T) {

  client := &Firefox{UserJS: "support/drivers/user.js", Extension: &Extension{Path: "support/webdriver.xpi"}}

  client.SetDefaults()

  if data, err := client.parseUserJS(); err == nil {

    if data["browser.newtab.url"] != "\"about:blank\"" {
      t.Error("browser.newtab.url should equal \"about:blank\"", data["browser.newtab.url"])
    }

  } else {
    t.Error("parseUserJS()", err)
  }

}

//////////////////////////////////////////////////////////////
func TestParseUserJSNoSemiColon(t *testing.T) {

  client := &Firefox{UserJS: "support/drivers/user-no-semi-colon.js", Extension: &Extension{Path: "support/webdriver.xpi"}}

  client.SetDefaults()

  if data, err := client.parseUserJS(); err == nil {

    if data["browser.newtab.url"] != "\"about:blank\"" {
      t.Error("browser.newtab.url should equal \"about:blank\"", data["browser.newtab.url"])
    }

  } else {
    t.Error("parseUserJS()", err)
  }

}

//////////////////////////////////////////////////////////////
func TestDefaultUserJS(t *testing.T) {

  client := &Firefox{UserJS: "support/drivers/user.js", Extension: &Extension{Path: "support/webdriver.xpi"}}

  data := client.defaultUserJS()

  if data == nil {
    t.Error("Um, something bad happened...")
  }

  if data["browser.startup.homepage"] != "\"about:blank\"" {
    t.Error("browser.startup.homepage should equal about:blank: ", data["browser.startup.homepage"])
  }

}

//////////////////////////////////////////////////////////////
func TestBuildUserJS(t *testing.T) {

  client := &Firefox{UserJS: "support/drivers/user.js", Extension: &Extension{Path: "support/webdriver.xpi"}}

  client.UserJSPolicy = "merge"

  data, _ := client.buildUserJS()

  if data == nil {
    t.Error("Um, something bad happened...")
  }

  if data["browser.startup.homepage"] != "\"my_home_page\"" {
    t.Error("browser.startup.homepage should equal my_home_page: ", data["browser.startup.homepage"])
  }

}

//////////////////////////////////////////////////////////////
func TestBuildUserJSNoMerge(t *testing.T) {

  client := &Firefox{UserJS: "support/drivers/user.js", Extension: &Extension{Path: "support/webdriver.xpi"}}

  data, _ := client.buildUserJS()

  if data == nil {
    t.Error("Um, something bad happened...")
  }

  if data["browser.startup.homepage"] != "\"my_home_page\"" {
    t.Error("browser.startup.homepage should equal my_home_page: ", data["browser.startup.homepage"])
  }

}

//////////////////////////////////////////////////////////////
func TestBuildUserJSSetNoExist(t *testing.T) {

  client := &Firefox{UserJS: "xsupport/drivers/user.js", Extension: &Extension{Path: "support/webdriver.xpi"}}

  data, _ := client.buildUserJS()

  if data == nil {
    t.Error("Um, something bad happened...")
  }

  if data["browser.startup.homepage"] != "\"about:blank\"" {
    t.Error("browser.startup.homepage should equal about:blank: ", data["browser.startup.homepage"])
  }

}

//////////////////////////////////////////////////////////////
func TestBuildUserJSNotSet(t *testing.T) {

  client := &Firefox{Extension: &Extension{Path: "support/webdriver.xpi"}}

  data, _ := client.buildUserJS()

  if data == nil {
    t.Error("Um, something bad happened...")
  }

  if data["browser.startup.homepage"] != "\"about:blank\"" {
    t.Error("browser.startup.homepage should equal about:blank: ", data["browser.startup.homepage"])
  }

}














