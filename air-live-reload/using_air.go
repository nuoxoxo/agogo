package main

import (
  "log" // for printing req info
  "net/http"
)

func main() {

  http.HandleFunc("/", handle_index)
  http.ListenAndServe(":10086", nil)
}

// helper

func handle_index (w http.ResponseWriter, r * http.Request) {

  ending := " \n"

  log.Println("method/ ", r.Method, ending)
  log.Println("URL   / ", r.URL, ending)
  log.Println("header/ ", r.Header, ending)

  w.Write( []byte("hello world") )
}

