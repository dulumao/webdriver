package webdriver

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
)

type (

  WireHTTP struct {

    // a url pointing to a running server supporting the JsonWireProtocol.
    // typically, http://localhost:7055 for firefox.
    BaseUrl string

    // represents a JsonWireProtocol Session ID.
    // The Session struct includes *WireHTTP, so, SessionID is available
    // to individual sessions.
    //
    // Most of the JsonWireProtocol API calls require a session id.
    // Only a couple do not.  GetFullUrl() will search for :sessionid
    // and replace it with SessionID during API calls.
    //
    // By default, SessionID is "", so, there should be no impact
    // for API calls that do not require a :sessionid
    SessionID string
  }

)

// Convenience method that wraps NewRequest()
func (s *WireHTTP) DeleteRequest(url string, payload interface{}) (req *http.Request, err error) {
  return s.NewRequest("DELETE", url, payload)
}

// Convenience method that wraps NewRequest()
func (s *WireHTTP) GetRequest(url string, payload interface{}) (req *http.Request, err error) {
  return s.NewRequest("GET", url, payload)
}

// Convenience method that wraps NewRequest()
func (s *WireHTTP) PostRequest(url string, payload interface{}) (req *http.Request, err error) {
  return s.NewRequest("POST", url, payload)
}

// Constructs a new http.Request for the defined method and url including
// a payload.  Default http headers required by JsonWireProtocol are added
// for you based on the type of method (GET, POST, etc.)
//
// method - The type of request GET, POST, etc.
//
// url - The url of the request without the host and port.
// Host, port, and session id are included automatically.
//
// payload - JSON values to be included in the request.
//
func (s *WireHTTP) NewRequest(method string, url string, payload interface{}) (req *http.Request, err error) {

  var body []byte

  if payload == nil {
    payload = map[string]interface{}{}
  }

  if body, err = json.Marshal(payload); err == nil {

    if req, err = http.NewRequest(method, s.BuildFullUrl(url), bytes.NewBuffer(body)); err == nil {

      req.Header.Set("Accept", "application/json")
      req.Header.Set("Accept-charset", "utf-8")

      if method == "POST" || method == "DELETE" {
        req.Header.Add("Content-Type", "application/json;charset=utf-8")
      }

    }
  }

  return req, err
}

// Builds a complete url for a request including host and port.
// Relies on the current value of BaseUrl and SessionID.
//
//   // given:
//     BaseUrl = "http://localhost:7055"
//     SessionID = "my-session-id"
//
//   // the following call
//   BuildFullUrl("/session/:sessionid/forward")
//
//   // would produce
//   http://localhost:7055/session/my-session-id/forward
//
func (s *WireHTTP) BuildFullUrl(url string) string {
  return fmt.Sprintf("%v%v", s.BaseUrl, strings.Replace(url, ":sessionid", s.SessionID, -1))
}

// Submits a request to a JsonWireProtocol server (selenium webdriver)
// and reads the response back into a WireResponse if the server
// responds with status code 200.
func (s *WireHTTP) Do(req *http.Request) (wireResponse *WireResponse, err error) {

  var resp *http.Response
  if resp, err = http.DefaultClient.Do(req); err == nil {

    wireResponse = &WireResponse{}

    fmt.Println("status: ", resp.StatusCode)

    // looking at the code for the Do method of the DefaultClient in the
    // http package.  It looks like I shouldn't have to be concerned with
    // redirects as it appears to handle them.
    if resp.StatusCode == 200 {

      var buffer []byte
      if buffer, err = ioutil.ReadAll(resp.Body); err == nil {

        err = json.Unmarshal(buffer, wireResponse)

      }
    }
  }

  return wireResponse, err
}


























