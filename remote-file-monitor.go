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

  if len(url) == 0 {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  resp, err := http.Get(url)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  } else {
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      return
    }

    var contents = string(body)

    fmt.Fprintf(w, "%d", len(contents))
    return
  }
}
