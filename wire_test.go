package webdriver

import (
  "fmt"
  // "encoding/json"
  // "net/http"
  // "net/http/httptest"
  "testing"
)

////////////////////////////////////////////////////////////////
func TestStringValue(t *testing.T) {

  for _, v := range sessions {
    fmt.Println("================================= session", v)
    if wireResponse, err := v.GetCapabilities(); err == nil {
      fmt.Println("wtf?", wireResponse.StringValue())
    }
  }

}

// ////////////////////////////////////////////////////////////////
// func TestStringValue(t *testing.T) {

//   wireResponse := &WireResponse{Value: []byte("\"this is a test\"")}

//   if wireResponse.StringValue() != "this is a test" {
//     t.Error("wtf?: StringValue() should have returned: this is a test  => ", wireResponse.StringValue())
//   }

//   wireResponse = &WireResponse{Value: []byte("this is a another test")}

//   if wireResponse.StringValue() != "this is a another test" {
//     t.Error("wtf?: StringValue() should have returned: this is a another test  => ", wireResponse.StringValue())
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestSource(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/source" {
//       t.Error("url path should be /session/my-session-id/source =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.Source(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", len(wireResponse.StringValue()))
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestTitle(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/title" {
//       t.Error("url path should be /session/my-session-id/title =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"my title\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.Title(); err == nil {
//     if wireResponse.StringValue() != "my title" {
//       t.Error("should have responded with my title", len(wireResponse.StringValue()))
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestStatus(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/status" {
//       t.Error("url path should be /status =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL

//   if wireResponse, err := s.Status(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", len(wireResponse.StringValue()))
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestSessions(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/sessions" {
//       t.Error("url path should be /sessions =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL

//   if wireResponse, err := s.Sessions(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", len(wireResponse.StringValue()))
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestGetSession(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id" {
//       t.Error("url path should be /session/my-session-id =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.GetSession(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", len(wireResponse.StringValue()))
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestDeleteSession(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id" {
//       t.Error("url path should be /session/my-session-id =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.DeleteSession(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", len(wireResponse.StringValue()))
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestTimeouts(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/timeouts" {
//       t.Error("url path should be /session/my-session-id/timeouts =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.Timeouts("script", 10000); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", len(wireResponse.StringValue()))
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestTimeoutsAsyncScript(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/timeouts/async_script" {
//       t.Error("url path should be /session/my-session-id/timeouts/async_script =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.TimeoutsAsyncScript(10000); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", len(wireResponse.StringValue()))
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestTimeoutsImplicitWait(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/timeouts/implicit_wait" {
//       t.Error("url path should be /session/my-session-id/timeouts/implicit_wait =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.TimeoutsImplicitWait(10000); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", len(wireResponse.StringValue()))
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestWindowHandle(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/window_handle" {
//       t.Error("url path should be /session/my-session-id/window_handle =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.WindowHandle(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", wireResponse.StringValue())
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestWindowHandles(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/window_handles" {
//       t.Error("url path should be /session/my-session-id/window_handles =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.WindowHandles(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", wireResponse.StringValue())
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestUrl(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/url" {
//       t.Error("url path should be /session/my-session-id/url =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.Url("http://www.example.com/"); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", wireResponse.StringValue())
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestGetUrl(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/url" {
//       t.Error("url path should be /session/my-session-id/url =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.GetUrl(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", wireResponse.StringValue())
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestForward(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/forward" {
//       t.Error("url path should be /session/my-session-id/forward =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.Forward(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", wireResponse.StringValue())
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestBack(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/back" {
//       t.Error("url path should be /session/my-session-id/back =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.Back(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", wireResponse.StringValue())
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestRefresh(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/refresh" {
//       t.Error("url path should be /session/my-session-id/refresh =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.Refresh(); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", wireResponse.StringValue())
//     }
//   }

// }

// ////////////////////////////////////////////////////////////////
// func TestExecute(t *testing.T) {

//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//     if r.URL.Path != "/session/my-session-id/execute" {
//       t.Error("url path should be /session/my-session-id/execute =>", r.URL.Path)
//     }

//     body, _ := json.Marshal(&WireResponse{
//                                 Name: "test",
//                                 Status: 0,
//                                 Value: []byte("\"ok\""),
//                                 })

//     w.Write(body)
//   }))

//   defer ts.Close()

//   s := &Wire{}
//   s.BaseUrl = ts.URL
//   s.SessionID = "my-session-id"

//   if wireResponse, err := s.Execute("alert('dude')", ""); err == nil {
//     if wireResponse.StringValue() != "ok" {
//       t.Error("should have responded with ok", wireResponse.StringValue())
//     }
//   }

// }







































