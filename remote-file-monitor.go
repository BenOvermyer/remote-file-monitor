package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
  http.HandleFunc("/size", sizeHandler)
  http.ListenAndServe(":8090", nil)
}

func sizeHandler(w http.ResponseWriter, r *http.Request) {
  var url = r.FormValue("url")
  resp, err := http.Get(url)

  if err != nil {
    fmt.Fprintf(w, "%s", err)
  } else {
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
      fmt.Fprintf(w, "%s", err)
    }

    var contents = string(body)

    fmt.Fprintf(w, "%d", len(contents))
  }
}
