package webdriver

import (
  "encoding/json"
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
)

////////////////////////////////////////////////////////////////
// tests 302 redirect.
// redirects 5 times, then, it should succeed.
func TestRedirect01(t *testing.T) {

  count := 0

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    count += 1
    if count < 5 {
      w.Header().Set("Location", fmt.Sprintf("%v%v", r.Host, r.URL))
      w.WriteHeader(302)
      } else {

        body, _ := json.Marshal(&WireResponse{
                                    Name: "redirectTest",
                                    Value: []byte("\"this is a test\""),
                                    })

        w.Write(body)
      }
  }))

  defer ts.Close()

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = ts.URL

  if req, wtf := s.NewRequest("GET", "/", nil); wtf == nil {
    if wireResponse, err := s.Do(req); err == nil {

      if wireResponse.StringValue() != "this is a test" {
        t.Error("wireResponse.StringValue() should be this is a test => ", wireResponse.StringValue())
      }
    }
  }

}

////////////////////////////////////////////////////////////////
// tests 302 redirect.
// redirects until max redirects have exceeded.
// wireResponse should be nil
func TestRedirect02(t *testing.T) {

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Location", fmt.Sprintf("%v%v", r.Host, r.URL))
      w.WriteHeader(302)
  }))

  defer ts.Close()

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = ts.URL

  if req, wtf := s.NewRequest("GET", "/", nil); wtf == nil {
    if _, err := s.Do(req); err == nil {
      t.Error("this should have failed and did not.  err should NOT be nil")
    }
  } else {
    t.Error("Unknown error", wtf)
  }

}

////////////////////////////////////////////////////////////////
// tests 301 redirect.
// redirects 5 times, then, it should succeed.
func TestRedirect03(t *testing.T) {

  count := 0

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    count += 1
    if count < 5 {
      w.Header().Set("Location", fmt.Sprintf("%v%v", r.Host, r.URL))
      w.WriteHeader(301)
      } else {

        body, _ := json.Marshal(&WireResponse{
                                    Name: "redirectTest",
                                    Value: []byte("\"this is a test\""),
                                    })

        w.Write(body)
      }
  }))

  defer ts.Close()

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = ts.URL

  if req, wtf := s.NewRequest("GET", "/", nil); wtf == nil {
    if wireResponse, err := s.Do(req); err == nil {

      if wireResponse.StringValue() != "this is a test" {
        t.Error("wireResponse.StringValue() should be this is a test => ", wireResponse.StringValue())
      }
    }
  }

}

////////////////////////////////////////////////////////////////
// tests 301 redirect.
// redirects until max redirects have exceeded.
// wireResponse should be nil
func TestRedirect04(t *testing.T) {

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Location", fmt.Sprintf("%v%v", r.Host, r.URL))
      w.WriteHeader(301)
  }))

  defer ts.Close()

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = ts.URL

  if req, wtf := s.NewRequest("GET", "/", nil); wtf == nil {
    if _, err := s.Do(req); err == nil {
      t.Error("this should have failed and did not.  err should NOT be nil")
    }
  } else {
    t.Error("Unknown error", wtf)
  }

}

////////////////////////////////////////////////////////////////
func TestNewRequestDelete(t *testing.T) {

  var req *http.Request
  var err error

  payload, _ := json.Marshal(&Params{"this": "that", "the": "other"})

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = "http://localhost:7055"
  s.SessionID = "my-session-id"

  if req, err = s.NewRequest("DELETE", "/session/:sessionid/nothing", payload); err == nil {

    if req.Method != "DELETE" {
      t.Error("wtf", "req.Method should be DELETE => ", req.Method)
    }

    if req.URL.Host != "localhost:7055" {
      t.Error("wtf", "req.URL.Host should equal localhost:7055 => ", req.URL.Host)
    }

    if req.URL.Path != "/session/my-session-id/nothing" {
      t.Error("wtf", "req.URL.Path should equal /session/my-session-id/nothing => ", req.URL.Path)
    }

    if req.Header["Accept"][0] != "application/json" {
      t.Error("wtf", "req.Header Accept should be application/json => ", req.Header["Accept"])
    }

    if req.Header["Accept-Charset"][0] != "utf-8" {
      t.Error("wtf", "req.Header Accept-Charset should be utf-8 => ", req.Header["Accept-Charset"])
    }

    if req.Header["Content-Type"][0] != "application/json;charset=utf-8" {
      t.Error("wtf", "req.Header Content-Type should be application/json;charset=utf-8 => ", req.Header["Content-Type"])
    }

  } else {
    t.Error("error should have been nil", err)
  }

}

////////////////////////////////////////////////////////////////
func TestNewRequestGet(t *testing.T) {

  var req *http.Request
  var err error

  payload, _ := json.Marshal(&Params{"this": "that", "the": "other"})

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = "http://localhost:7055"
  s.SessionID = "my-session-id"

  if req, err = s.NewRequest("GET", "/session/:sessionid/nothing", payload); err == nil {

    if req.Method != "GET" {
      t.Error("wtf", "req.Method should be GET => ", req.Method)
    }

    if req.URL.Host != "localhost:7055" {
      t.Error("wtf", "req.URL.Host should equal localhost:7055 => ", req.URL.Host)
    }

    if req.URL.Path != "/session/my-session-id/nothing" {
      t.Error("wtf", "req.URL.Path should equal /session/my-session-id/nothing => ", req.URL.Path)
    }

    if req.Header["Accept"][0] != "application/json" {
      t.Error("wtf", "req.Header Accept should be application/json => ", req.Header["Accept"])
    }

    if req.Header["Accept-Charset"][0] != "utf-8" {
      t.Error("wtf", "req.Header Accept-Charset should be utf-8 => ", req.Header["Accept-Charset"])
    }

    if len(req.Header["Content-Type"]) != 0 {
      t.Error("wtf", "req.Header Content-Type should be missing => ", req.Header["Content-Type"])
    }

  } else {
    t.Error("error should have been nil", err)
  }

}

////////////////////////////////////////////////////////////////
func TestNewRequestPost(t *testing.T) {

  var req *http.Request
  var err error

  payload, _ := json.Marshal(&Params{"this": "that", "the": "other"})

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = "http://localhost:7055"
  s.SessionID = "my-session-id"

  if req, err = s.NewRequest("POST", "/session/:sessionid/nothing", payload); err == nil {

    if req.Method != "POST" {
      t.Error("wtf", "req.Method should be POST => ", req.Method)
    }

    if req.URL.Host != "localhost:7055" {
      t.Error("wtf", "req.URL.Host should equal localhost:7055 => ", req.URL.Host)
    }

    if req.URL.Path != "/session/my-session-id/nothing" {
      t.Error("wtf", "req.URL.Path should equal /session/my-session-id/nothing => ", req.URL.Path)
    }

    if req.Header["Accept"][0] != "application/json" {
      t.Error("wtf", "req.Header Accept should be application/json => ", req.Header["Accept"])
    }

    if req.Header["Accept-Charset"][0] != "utf-8" {
      t.Error("wtf", "req.Header Accept-Charset should be utf-8 => ", req.Header["Accept-Charset"])
    }

    if req.Header["Content-Type"][0] != "application/json;charset=utf-8" {
      t.Error("wtf", "req.Header Content-Type should be application/json;charset=utf-8 => ", req.Header["Content-Type"])
    }

  } else {
    t.Error("error should have been nil", err)
  }

}

////////////////////////////////////////////////////////////////
func TestDeleteRequest(t *testing.T) {

  var req *http.Request
  var err error

  payload, _ := json.Marshal(&Params{"this": "that", "the": "other"})

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = "http://localhost:7055"
  s.SessionID = "my-session-id"

  if req, err = s.DeleteRequest("/session/:sessionid/nothing", payload); err == nil {

    if req.Method != "DELETE" {
      t.Error("wtf", "req.Method should be DELETE => ", req.Method)
    }

    if req.URL.Host != "localhost:7055" {
      t.Error("wtf", "req.URL.Host should equal localhost:7055 => ", req.URL.Host)
    }

    if req.URL.Path != "/session/my-session-id/nothing" {
      t.Error("wtf", "req.URL.Path should equal /session/my-session-id/nothing => ", req.URL.Path)
    }

    if req.Header["Accept"][0] != "application/json" {
      t.Error("wtf", "req.Header Accept should be application/json => ", req.Header["Accept"])
    }

    if req.Header["Accept-Charset"][0] != "utf-8" {
      t.Error("wtf", "req.Header Accept-Charset should be utf-8 => ", req.Header["Accept-Charset"])
    }

    if req.Header["Content-Type"][0] != "application/json;charset=utf-8" {
      t.Error("wtf", "req.Header Content-Type should be application/json;charset=utf-8 => ", req.Header["Content-Type"])
    }

  } else {
    t.Error("error should have been nil", err)
  }

}

////////////////////////////////////////////////////////////////
func TestGetRequest(t *testing.T) {

  var req *http.Request
  var err error

  payload, _ := json.Marshal(&Params{"this": "that", "the": "other"})

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = "http://localhost:7055"
  s.SessionID = "my-session-id"

  if req, err = s.GetRequest("/session/:sessionid/nothing", payload); err == nil {

    if req.Method != "GET" {
      t.Error("wtf", "req.Method should be GET => ", req.Method)
    }

    if req.URL.Host != "localhost:7055" {
      t.Error("wtf", "req.URL.Host should equal localhost:7055 => ", req.URL.Host)
    }

    if req.URL.Path != "/session/my-session-id/nothing" {
      t.Error("wtf", "req.URL.Path should equal /session/my-session-id/nothing => ", req.URL.Path)
    }

    if req.Header["Accept"][0] != "application/json" {
      t.Error("wtf", "req.Header Accept should be application/json => ", req.Header["Accept"])
    }

    if req.Header["Accept-Charset"][0] != "utf-8" {
      t.Error("wtf", "req.Header Accept-Charset should be utf-8 => ", req.Header["Accept-Charset"])
    }

    if len(req.Header["Content-Type"]) != 0 {
      t.Error("wtf", "req.Header Content-Type should be missing => ", req.Header["Content-Type"])
    }

  } else {
    t.Error("error should have been nil", err)
  }

}

////////////////////////////////////////////////////////////////
func TestPostRequest(t *testing.T) {

  var req *http.Request
  var err error

  payload, _ := json.Marshal(&Params{"this": "that", "the": "other"})

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = "http://localhost:7055"
  s.SessionID = "my-session-id"

  if req, err = s.PostRequest("/session/:sessionid/nothing", payload); err == nil {

    if req.Method != "POST" {
      t.Error("wtf", "req.Method should be POST => ", req.Method)
    }

    if req.URL.Host != "localhost:7055" {
      t.Error("wtf", "req.URL.Host should equal localhost:7055 => ", req.URL.Host)
    }

    if req.URL.Path != "/session/my-session-id/nothing" {
      t.Error("wtf", "req.URL.Path should equal /session/my-session-id/nothing => ", req.URL.Path)
    }

    if req.Header["Accept"][0] != "application/json" {
      t.Error("wtf", "req.Header Accept should be application/json => ", req.Header["Accept"])
    }

    if req.Header["Accept-Charset"][0] != "utf-8" {
      t.Error("wtf", "req.Header Accept-Charset should be utf-8 => ", req.Header["Accept-Charset"])
    }

    if req.Header["Content-Type"][0] != "application/json;charset=utf-8" {
      t.Error("wtf", "req.Header Content-Type should be application/json;charset=utf-8 => ", req.Header["Content-Type"])
    }

  } else {
    t.Error("error should have been nil", err)
  }

}

////////////////////////////////////////////////////////////////
func TestBuildFullUrl(t *testing.T) {

  s := &Wire{}
  s.SetDefaults()

  s.BaseUrl = "http://localhost:7055"
  s.SessionID = "no_session"

  if url := s.BuildFullUrl("/status"); url != "http://localhost:7055/status" {
    t.Error("BuildFullUrl should be: http://localhost:7055/status  value is => ", url)
  }

}

////////////////////////////////////////////////////////////////
func TestBuildFullUrlWithSessionID(t *testing.T) {

  s := &Wire{}
  s.SetDefaults()
  s.BaseUrl = "http://localhost:7055"
  s.SessionID = "my_session_id"

  if url := s.BuildFullUrl("/session/:sessionid"); url != "http://localhost:7055/session/my_session_id" {
    t.Error("BuildFullUrl should be: http://localhost:7055/session/my_session_id  value is => ", url)
  }

}





















