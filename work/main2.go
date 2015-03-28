package main

import (
  "encoding/json"
  "fmt"
)

func main() {

  jsonBlob := []byte{123, 34, 115, 101, 115, 115, 105, 111, 110, 73, 100, 34, 58, 34, 55, 98, 50, 53, 99, 101, 51, 54, 48, 52, 55, 49, 55, 57, 49, 50, 49, 49, 101, 49, 51, 102, 100, 57, 54, 52, 98, 53, 101, 48, 57, 51, 92, 117, 48, 48, 51, 67, 34, 125}
  x := []byte{34, 55, 98, 50, 53, 99, 101, 51, 54, 48, 52, 55, 49, 55, 57, 49, 50, 49, 49, 101, 49, 51, 102, 100, 57, 54, 52, 98, 53, 101, 48, 57, 51, 92, 117, 48, 48, 51, 67, 34}

  type Thing struct {
//        SessionID             json.RawMessage `json:"sessionId"`
        SessionID2             string `json:"sessionId"`
  }
  thing := &Thing{}
  err := json.Unmarshal(jsonBlob, thing)
  if err != nil {
    fmt.Println("error:", err)
  }
//  fmt.Printf("%+v", thing)
//  fmt.Println(string(thing.SessionID))
  fmt.Println(string(thing.SessionID2))
  // fmt.Println(string(jsonBlob))
  fmt.Println(string(x))

  var dude string
  err = json.Unmarshal(x, &dude)
  if err != nil {
    fmt.Println("error:", err)
  }

fmt.Println("dude here", dude)
}
