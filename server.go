package main

import (
  "fmt"
  "net/http"
)

var content =
`
<h1>
  Welcome to my Golang application
</h1>
<p>
  You are not logged in
</p>
`

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, content)
}

func main() {
  if loggedIn := true; loggedIn {
    content =
    `
    <h1>
      Welcome to my Golang application
    </h1>
    <p>
      You are logged in
    </p>
    `
  }

  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
