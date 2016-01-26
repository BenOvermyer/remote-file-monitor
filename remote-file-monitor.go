package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
)

func main() {
  var port = os.Getenv("RMF_PORT")
  if (len(port) == 0) {
    port = "8090"
  }
  http.HandleFunc("/size", sizeHandler)
  http.ListenAndServe(":" + port, nil)
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
