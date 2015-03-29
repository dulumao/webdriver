package webdriver

import (
  "fmt"
  // "testing"
  "html/template"
  "net/http"
)

////////////////////////////////////////////////////////////////
func startWebServer() {

  http.HandleFunc("/", handlerIndex)
  http.HandleFunc("/element.html", handlerElement)
  http.HandleFunc("/elements.html", handlerElements)
  http.HandleFunc("/source.html", handlerSource)
  http.HandleFunc("/step01.html", handlerStep01)
  http.HandleFunc("/step02.html", handlerStep02)
  http.HandleFunc("/step03.html", handlerStep03)

  fmt.Println("web server running at: http://localhost:8080")

  http.ListenAndServe(":8080", nil)

}

////////////////////////////////////////////////////////////////
func handlerIndex(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/index.html")

  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerElement(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/element.html")

  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerElements(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/elements.html")

  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerSource(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/source.html")

  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerStep01(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/step01.html")

  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerStep02(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/step02.html")

  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerStep03(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/step03.html")

  t.Execute(w, nil)

}
