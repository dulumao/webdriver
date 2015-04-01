package webdriver

import (
  "fmt"
  // "testing"
  "html/template"
  "net/http"
  "time"
)

////////////////////////////////////////////////////////////////
func startWebServer() {

  http.HandleFunc("/", handlerIndex)
  http.HandleFunc("/cookies.html", handlerCookies)
  http.HandleFunc("/element.html", handlerElement)
  http.HandleFunc("/elements.html", handlerElements)
  http.HandleFunc("/frame.html", handlerFrame)
  http.HandleFunc("/form01.html", handlerForm01)
  http.HandleFunc("/form01-post.html", handlerForm01Post)
  http.HandleFunc("/longpage.html", handlerLongPage)
  http.HandleFunc("/sub-element.html", handlerSubElement)
  http.HandleFunc("/sub-elements.html", handlerSubElements)
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
func handlerCookies(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/cookies.html")

  http.SetCookie(w,
    &http.Cookie{
       Domain: "localhost",
      Expires: time.Date(2020, 11, 23, 1, 5, 3, 0, time.UTC),
     HttpOnly: true,
         Name: "main",
         Path: "/",
        Value: "this-is-my-cookie-value-not-hard-to-decrypt",
       Secure: false,
    })

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
func handlerFrame(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/frame.html")

  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerForm01(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/form01.html")

  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerForm01Post(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/form01-post.html")
  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerLongPage(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/longpage.html")

  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerSubElement(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/sub-element.html")

  t.Execute(w, nil)

}

////////////////////////////////////////////////////////////////
func handlerSubElements(w http.ResponseWriter, r *http.Request) {

  t, _ := template.ParseFiles("templates/sub-elements.html")

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
