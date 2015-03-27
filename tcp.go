package webdriver

import (
  "errors"
  "fmt"
  "net"
  "time"
)

// Waits until a tcp connection can be made to a server
// or until the timeout has been exceeded.
func waitForConnect(host string, port int, timeout float64) (err error) {

  began := time.Now()

  address := fmt.Sprintf("%v:%d", host, port)

  for {

    // try to connect to the host/port
    if connection, err := net.Dial("tcp", address); err == nil {

      err = connection.Close()
      break

    }

    time.Sleep(1 * time.Second)

    // making it here means the lock failed and err is not nil
    // that can and will happen and we should try again unless the timeout
    // period has expired
    if float64(time.Since(began)) > timeout {
      return errors.New("waitForConnect() => timeout expired waiting to connect")
    }

  }

  return err
}

func waitForLock(host string, port int, timeout float64) (listener net.Listener, err error) {

  began := time.Now()

  address := fmt.Sprintf("%v:%d", host, port)

  for {

    // try to lock the port by listening on it
    if listener, err := net.Listen("tcp", address); err == nil {
      // if no error, then, it must be listening and is considered locked
      return listener, nil
    }

    time.Sleep(1 * time.Second)

    // making it here means the lock failed and err is not nil
    // that can and will happen and we should try again unless the timeout
    // period has expired
    if float64(time.Since(began)) > timeout {
      err = errors.New("WaitForLock() => timeout expired trying to lock mutex port")
      break
    }

  }

  return nil, err
}

func findNextAvailablePort(host string, port int, timeout float64) (newPort int, err error) {

  newPort = 0

  if port > 0 {

    began := time.Now()

    // the current value of the client.Port is considered the desired port
    // however, we need to find the next available port for use,
    // therefore, i uses the current desired port as a starting point
    for i := port; i < 65535; i++ {

      address := fmt.Sprintf("%v:%d", host, i)

      // try to lock the port by listening on it
      if listener, err := net.Listen("tcp", address); err == nil {

        // assign the port and close
        newPort = i

        // no need to defer here
        listener.Close()

        // if no error, then, it must be listening and is considered locked
        break
      }

      time.Sleep(1 * time.Second)

      // making it here means the lock failed and err is not nil
      // that can and will happen and we should try again unless the timeout
      // period has expired
      if float64(time.Since(began)) > timeout {
        err = errors.New("findNextAvailablePort() => timeout expired while determining next available port")
        break
      }

    }

  } else {
    err = errors.New(fmt.Sprintf("findNextAvailablePort() => Port needs to be set: ", port))
  }

  return newPort, err
}
