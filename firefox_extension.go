package webdriver

import (
  "archive/zip"
  "encoding/xml"
  "errors"
  "fmt"
  "log"
  "io"
  "io/ioutil"
  "os"
  "path/filepath"
  "strings"
)

type (

  Extension struct {
    BaseDir string
    ConfigList []string
    ConfigPolicy string
    Path string
    Name string
    Policy string
    TargetPath string
  }

  InstallManifest struct {
    Description InstallManifestDescription
  }

  InstallManifestDescription struct {
    ID string `xml:"id"`
  }

)

// Installs the webdriver.xpi extension into the current firefox profile directory.
func (s *Firefox) installExtension() error {

  var err error

  log.Println("installExtension()...")

  // verifies the profile directory exists.  if not, creates it
  // also, verifies profiledir/extensions exists.  if not, creates it
  // then, if the policy is to remove the extension; it is removed.
  if err = s.verifyExtensionsPath(); err == nil && s.Extension.Policy == "remove" {

    // directories exist and the policy is to remove and install the extension
    if s.Extension.Name, err = s.extractExtensionName(); err == nil {

      if err = s.prepareExtensionTarget(); err == nil {

        // just return the error.  no need to if logic, etc.
        err = s.extractExtensionContents()

      }

    }

  }

  return err
}

func (s *Firefox) verifyExtensionsPath() (err error) {

// The target ProfileDir and the extensions subdirectory need to exist
// in order to install and use the webdriver plugin.  So, verifyExtensionsPath
// will test for the existence of the full path to the extensions directory
// and attempt to create it if it does not exist.  That would create the ProfileDir
// as well during the process thus killing two birds with one stone.
//
// Returns and error if there is a problem, otherwise, returns nil
  log.Println("verifyExtensionsPath() verifying profile and extensions directory exists: ", s.Extension.BaseDir)

  // check if the full path to the extensions directory exists
  if _, err = os.Stat(s.Extension.BaseDir); err != nil {

    log.Println("profile / extensions directory DOES NOT exist.  attempting to created it")
    // the full path to the extensions directory DOES NOT exist
    // attempt to create it with MkdirAll.  That way, the ProfileDir
    // will be created as well if it does not exist (two birds one stone)
    err = os.MkdirAll(s.Extension.BaseDir, s.DirPermissions)
    if err != nil {
      buffer := fmt.Sprintf("verifyExtensionsPath() => %v DOES NOT exist and unable to create\n", s.Extension.BaseDir)
      buffer = fmt.Sprint("%v%v", buffer, err.Error())
      err = errors.New(buffer)
      log.Println(err)
    } else {
      log.Println("Directory created! ", s.Extension.BaseDir)
    }
  } else {
    log.Println("verifyExtensionsPath() profile directory ALREADY exists", s.Extension.BaseDir)
  }

  return err
}




// Extracts the entire contents of webdriver.xpi
func (s *Firefox) extractExtensionContents() (err error) {

  log.Println("Installing the extension at: ", s.Extension.TargetPath)

  var zipFile *zip.ReadCloser
  total := 0

  if _, err = os.Stat(s.Extension.Path); err == nil {
    if zipFile, err = zip.OpenReader(s.Extension.Path); err == nil {
      defer zipFile.Close()

      for _, f := range zipFile.File {

        if err = s.writeExtensionFile(f); err != nil {
          break
        }

        total += 1

      }

    }
  }

  log.Println("Total files installed: ", total)

  return err
}

// Writes the contents of a single file contained in webdriver.xpi
func (s *Firefox) writeExtensionFile(f *zip.File) (err error) {

  // seperated this into a function so I could use defer, etc.
  var file io.ReadCloser

  if file, err = f.Open(); err == nil {
    defer file.Close()

    targetFileSpec := filepath.Join(s.Extension.TargetPath, f.Name)
    if f.FileInfo().IsDir() {
      err = os.MkdirAll(targetFileSpec, s.DirPermissions)
    } else {

      var targetFile *os.File
      if targetFile, err = os.OpenFile(targetFileSpec,
                                        os.O_WRONLY | os.O_CREATE,
                                        s.FilePermissions); err == nil {

        defer targetFile.Close()
        _, err = io.Copy(targetFile, file)

      }

    }

  }

  return err
}

// Extracts the name of the extension from within install.rdf
// which is a file contained in webdriver.xpi
func (s *Firefox) extractExtensionName() (extensionName string, err error) {

  var zipFile *zip.ReadCloser
  var file io.ReadCloser

  if _, err = os.Stat(s.Extension.Path); err == nil {
    // there is a file named install.rdf within the webdriver extension
    // the contents of that file contain the plugin name
    // open the zip archive (webdriver extension) and
    // extract the contents of of install.rdf to obtain the file name
    if zipFile, err = zip.OpenReader(s.Extension.Path); err == nil {
      defer zipFile.Close()

      for _, f := range zipFile.File {

        if extensionName == "" {
          if strings.ToLower(f.Name) == "install.rdf" {

            if file, err = f.Open(); err == nil {

              var buffer []byte
              if buffer, err = ioutil.ReadAll(file); err == nil {

                manifest := InstallManifest{}
                if err = xml.Unmarshal(buffer, &manifest); err == nil {
                  extensionName = manifest.Description.ID
                }
              }

              file.Close()

            }
          }
        } else {
          break
        }
      }

    }
  }

  return extensionName, err
}

// Prepares the extension target installation directory.  if it exists, then,
// it is removed and created again.
func (s *Firefox) prepareExtensionTarget() (err error) {

  s.Extension.TargetPath = filepath.Join(s.Extension.BaseDir, s.Extension.Name)
  log.Println("Preparing the extension target installation directory: ", s.Extension.TargetPath)

  if _, err2 := os.Stat(s.Extension.TargetPath); err2 == nil {

    err = os.RemoveAll(s.Extension.TargetPath)

  }

  if err == nil {

    err = os.MkdirAll(s.Extension.TargetPath, s.DirPermissions)
    if err != nil {
      buffer := fmt.Sprintf("prepareExtensionTarget() => %v unable to create extension target directory\n", s.Extension.TargetPath)
      buffer = fmt.Sprint("%v%v", buffer, err.Error())
      err = errors.New(buffer)
    }

  }

  return err
}

// The spec recommends removing the following files:
//      compatibility.ini
//      extensions.cache
//      extensions.ini
//      extensions.json
//      extensions.rdf
//      extensions.sqlite
//      extensions.sqlite-journal
//
// By default, removeExtensionConfig() will remove the above list.  You can add / remove
// files to / from the list, however, it will ONLY remove files in the root of the extensions
// directory.
//
// You have the option of bypassing by setting Extension.ConfigPolicy = ""
func (s *Firefox) removeExtensionConfig() (err error) {

  if s.Extension.ConfigPolicy == "remove" {

    for _, v := range s.Extension.ConfigList {

      if err == nil {
        targetFileSpec := filepath.Join(s.ProfileDir, v)
        // nil means the file/directory exists
        if _, err2 := os.Stat(targetFileSpec); err2 == nil {
          err = os.Remove(targetFileSpec)
        }
      }
    }

  }

  return err
}






















