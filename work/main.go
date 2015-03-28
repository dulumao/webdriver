package main

import (
  "fmt"
  // "unicode/utf8"
  "reflect"
)

func main() {
  // a := []byte{123, 34, 115, 101, 115, 115, 105, 111, 110, 73, 100, 34, 58, 34, 55, 98, 50, 53, 99, 101, 51, 54, 48, 52, 55, 49, 55, 57, 49, 50, 49, 49, 101, 49, 51, 102, 100, 57, 54, 52, 98, 53, 101, 48, 57, 51, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 48, 44, 34, 118, 97, 108, 117, 101, 34, 58}
  a := []byte{123, 34, 115, 101, 115, 115, 105, 111, 110, 73, 100, 34, 58, 34, 55, 98, 50, 53, 99, 101, 51, 54, 48, 52, 55, 49, 55, 57, 49, 50, 49, 49, 101, 49, 51, 102, 100, 57, 54, 52, 98, 53, 101, 48, 57, 51, 34, 125}
  // x := []byte{34, 92, 117, 48, 48, 51, 67, 104, 116, 109, 108, 62}
  x := []byte{92, 117, 48, 48, 51, 67}
  y := "\"\u003Chtml>"
  z := string(x)
  // fmt.Println("Hello, playground", string(x))
  // fmt.Println(len(x))
  // fmt.Println("start")

  fmt.Println("dude")
  fmt.Println(len(x), reflect.TypeOf(x), x)
  fmt.Println(len(y), reflect.TypeOf(y), y)
  fmt.Println(len(z), reflect.TypeOf(z), z)
  fmt.Println(len(a), reflect.TypeOf(a), string(a))

  fmt.Println("............ loop x")
  for i, v := range x {
    fmt.Println(i, v)
  }

  fmt.Println("............ loop y")
  for i, v := range y {
    fmt.Println(i, v)
  }

  fmt.Println("............ loop z")
  for i, v := range z {
    fmt.Println(i, v)
  }

  // wtf, _ := utf8.DecodeRune(x)
  // fmt.Println("unicode", wtf)

}
