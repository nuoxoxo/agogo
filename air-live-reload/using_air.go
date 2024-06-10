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
  // var N int
  var msg string

  log.Println("method/", r.Method, ending)
  log.Println("URL   /", r.URL, ending)
  for key, val := range r.Header {
    if len(val) != 1 {
      print(key, " - ", val)
    }
    if len(val[0]) < 42 {
      msg = val[0]
    } else {
      msg = val[0][:42]
      msg += " ... "
    }
    log.Println("header/", key, msg)
  }
  // log.Println("header/ ", r.Header, ending)

  w.Write([]byte("hello world"))
}

