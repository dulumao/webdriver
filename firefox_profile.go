package webdriver

import (
  "bufio"
  "errors"
  "fmt"
  "log"
  "os"
  "path/filepath"
  "strings"
)

// Configures the Firefox profile directory.  Installs the extension.  Copies in user.js settings, etc.
func (s *Firefox) configProfile() (err error) {

  log.Println("Configuring profile...")

  if err = s.installExtension(); err == nil {

    if err = s.removeExtensionConfig(); err == nil {

      // builds the contents of the user.js file
      var userJS map[string]interface{}
      if userJS, err = s.buildUserJS(); err == nil {

        // this is how we are telling webdriver which port to use
        userJS["webdriver_firefox_port"] = s.Port

        targetFileSpec := filepath.Join(s.ProfileDir, "user.js")

        var file *os.File

        if file, err = os.OpenFile(targetFileSpec, os.O_WRONLY | os.O_CREATE, s.FilePermissions); err == nil {

          defer file.Close()

          log.Println("configProfile() writing user.js to: ", targetFileSpec)

          for k, v := range userJS {
            file.WriteString(fmt.Sprintf("user_pref(\"%v\", %v);\n", k, v))
          }

        }

      }

    }
  }

  return err
}

// Parses the contents of a user.js file into a map.  Developer has the option to use a custom
// user.js file.  Setting Firefox.UserJS to a full path to a file will cause buildUserJS() to load
// the contents of that file.  if Firefox.UserJSPolicy is set to "merge", then, the contents of the
// custom user.js file is merged with the default values.  If Firefox.UserJSPolicy is set to anything
// other than "merge", the contents of the custom user.js are used "as is".  However, regardless of the
// configuration, webdriver_firefox_port is ALWAYS assigned during configProfile()
func (s *Firefox) buildUserJS() (values map[string]interface{}, err error) {

  values = s.defaultUserJS()

  if s.UserJS != "" {

    log.Println("buildUserJS() user.js set to: ", s.UserJS)

    var userJS map[string]interface{}

    if userJS, err = s.parseUserJS(); err == nil {

      log.Println("buildUserJS() user.js parsed ok")
      // this means a valid UserJS file was found and parsed

      if s.UserJSPolicy == "merge" {

        log.Println("buildUserJS() merging user.js with default values")

        for k, v := range userJS {
          values[k] = v
        }

      } else {
        log.Println("buildUserJS() overwriting default values with user.js")
        values = userJS
      }

    } else {
      log.Println("buildUserJS() could not read: ", s.UserJS, err)
    }

  } else {
    log.Println("buildUserJS() user.js not set.  using default values")
  }

  return values, err
}

// Parses the contents of a user.js file into a map.
func (s *Firefox) defaultUserJS() (values map[string]interface{}) {
  values = map[string]interface{}{
          "app.update.auto":                                        "false",
          "app.update.enabled":                                     "false",
          // "browser.cache.disk.capacity":                            "0",
          // "browser.cache.disk.enable":                              "false",
          // "browser.cache.memory.enable":                            "true",
          // "browser.dom.window.dump.enabled":                        "true",
          "browser.download.manager.showWhenStarting":              "false",
          "browser.EULA.3.accepted":                                "true",
          "browser.EULA.override":                                  "true",
          "browser.link.open_external":                             "2",
          "browser.link.open_newwindow":                            "2",
          "browser.newtab.url":                                     "\"about:blank\"",
          // "browser.newtabpage.enabled":                             "false",
          "browser.offline":                                        "false",
          "browser.safebrowsing.enabled":                           "false",
          "browser.safebrowsing.malware.enabled":                   "false",
          "browser.search.update":                                  "false",
          "browser.sessionstore.resume_from_crash":                 "false",
          "browser.shell.checkDefaultBrowser":                      "false",
          "browser.startup.homepage":                               "\"about:blank\"",
          // "browser.startup.homepage_override.mstone":               "ignore",
          "browser.startup.page":                                   "0",
          // "browser.tabs.insertRelatedAfterCurrent":                 "false",
          "browser.tabs.warnOnClose":                               "false",
          "browser.tabs.warnOnOpen":                                "false",

          "datareporting.healthreport.service.enabled":             "false",
          "datareporting.healthreport.uploadEnabled":               "false",
          "datareporting.healthreport.service.firstRun":            "false",
          "datareporting.healthreport.logging.consoleEnabled":      "false",
          "datareporting.policy.dataSubmissionEnabled":             "false",
          "datareporting.policy.dataSubmissionPolicyAccepted":      "false",

          "devtools.errorconsole.enabled":                          "true",
          "dom.disable_open_during_load":                           "false",
          // "dom.max_chrome_script_run_time":                         "30",
          // "dom.max_script_run_time":                                "30",
          // "dom.report_all_js_exceptions":                           "true",
          "extensions.autoDisableScopes":                           "10",
          "extensions.blocklist.enabled":                           "false",
          "extensions.logging.enabled":                             "true",
          "extensions.update.enabled":                              "false",
          "extensions.update.notifyUser":                           "false",
          // "javascript.options.showInConsole":                       "true",
          // "network.http.max-connections-per-server":                "10",
          "network.http.phishy-userpass-length":                    "255",
          "network.manage-offline-status":                          "false",
          "offline-apps.allow_by_default":                          "true",
          "prompts.tab_modal.enabled":                              "false",
          "security.csp.enable":                                    "false",
          "security.fileuri.origin_policy":                         "3",
          "security.fileuri.strict_origin_policy":                  "false",
          "security.warn_entering_secure":                          "false",
          "security.warn_entering_secure.show_once":                "false",
          "security.warn_entering_weak":                            "false",
          "security.warn_entering_weak.show_once":                  "false",
          "security.warn_leaving_secure":                           "false",
          "security.warn_leaving_secure.show_once":                 "false",
          "security.warn_submit_insecure":                          "false",
          // "security.warn_submit_insecure.show_once":                "false",
          "security.warn_viewing_mixed":                            "false",
          "security.warn_viewing_mixed.show_once":                  "false",
          "signon.rememberSignons":                                 "false",
          "startup.homepage_welcome_url":                           "\"about:blank\"",
          "toolkit.networkmanager.disable":                         "true",
          "toolkit.telemetry.enabled":                              "false",
          "toolkit.telemetry.prompted":                             "2",
          "toolkit.telemetry.rejected":                             "true",
          "webdriver_accept_untrusted_certs":                       "true",
          "webdriver_assume_untrusted_issuer":                      "true",
          "webdriver_enable_native_events":                         "false",
          "webdriver_unexpected_alert_behaviour":                   "\"dismiss\"",
// "webdriver.firefox.logfile": "/tmp/firefox.log",
        }



  return values
}

// Parses the contents of a user.js file into a map.
func (s *Firefox) parseUserJS() (values map[string]interface{}, err error) {

  values = make(map[string]interface{})

  if _, err = os.Stat(s.UserJS); err == nil {

    log.Println("Parsing user.js file at: ", s.UserJS)

    // sample
    // user_pref("browser.startup.homepage", "about:blank");
    // user_pref("dom.max_chrome_script_run_time", 30);
    // user_pref("dom.max_script_run_time", 30);
    // user_pref("dom.report_all_js_exceptions", true);
    // user_pref("javascript.options.showInConsole", true);
    // user_pref("network.http.max-connections-per-server", 10);

    var file *os.File
    if file, err = os.Open(s.UserJS); err == nil {

      defer file.Close()

      scanner := bufio.NewScanner(file)
      for scanner.Scan() {

        // get the full line
        buffer := scanner.Text()

        // strip away user_pref("
        buffer = strings.Replace(buffer, "user_pref(\"", "", 1)

        // if the string ends with ), slice it and remove the last )
        if strings.HasSuffix(buffer, ")") {
          buffer = buffer[0:len(buffer) - 1]
        }

        // if the string ends with ), slice it and remove the last );
        if strings.HasSuffix(buffer, ");") {
          buffer = buffer[0:len(buffer) - 2]
        }

        // now, should have something like:
        // browser.startup.homepage", "about:blank"
        // dom.max_chrome_script_run_time", 30
        // dom.max_script_run_time", 30
        // dom.report_all_js_exceptions", true

        // split the string into two pieces based on ",
        pair := strings.Split(buffer, "\",")

        key := strings.Trim(pair[0], " ")
        value := strings.Trim(pair[1], " ")

        values[key] = value
      }

      err = scanner.Err()

    }

  } else {
    err = errors.New("parseUserJS() => UserJS file not found:")
    log.Println("parseUserJS() => UserJS file not found: ", s.UserJS)
  }

  return values, err
}






























