package main

import (
  "net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  generateGif(w)
}

func main() {
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":3000", nil); err != nil {
    panic(err)
  }
}
