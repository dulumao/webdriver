package webdriver

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

// Convenience method that wraps NewRequest()
func (s *Wire) DeleteRequest(url string, payload interface{}) (req *http.Request, err error) {
  return s.NewRequest("DELETE", url, payload)
}

// Convenience method that wraps NewRequest()
func (s *Wire) GetRequest(url string, payload interface{}) (req *http.Request, err error) {
  return s.NewRequest("GET", url, payload)
}

// Convenience method that wraps NewRequest()
func (s *Wire) PostRequest(url string, payload interface{}) (req *http.Request, err error) {
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
func (s *Wire) NewRequest(method string, url string, payload interface{}) (req *http.Request, err error) {

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

// Submits a request to a JsonWireProtocol server (selenium webdriver)
// and reads the response back into a WireResponse if the server
// responds with status code 200.
func (s *Wire) Do(req *http.Request) (wireResponse *WireResponse, err error) {

  // never be nil
  wireResponse = &WireResponse{}

  var resp *http.Response
  if resp, err = http.DefaultClient.Do(req); err == nil {

    fmt.Println("status: ", resp.StatusCode)

    wireResponse.HttpStatusCode = resp.StatusCode

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


























