package main

import (
  "fmt"
  "os"
  "io"
  "log"
  "net/http"
  "encoding/json"
  "../common"
)

func posteventsHandler(rw http.ResponseWriter, request *http.Request) {
  decoder := json.NewDecoder(request.Body)

  var events common.EventSlice

  err := decoder.Decode(&events)
  if err != nil {
    log.Println(err)
  } else {
    log.Println(events)
  }

  defer request.Body.Close()
}


func main() {
  logfile, err := setupLogging("cruckserver.log", true)
  if err != nil {
    defer logfile.Close()
  }

  
  http.HandleFunc("/postevents", posteventsHandler)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupLogging(logFilepath string, useStdout bool) (f *os.File, err error) {
  f, err = os.OpenFile(logFilepath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
    fmt.Println("error opening file: %v", logFilepath)
    return nil, err
  }

  if useStdout {
    mw := io.MultiWriter(os.Stdout, f)
    log.SetOutput(mw)
  } else {
    log.SetOutput(f)
  }

  log.Println("------------------------------")
  log.Println("cruck server log started")

  return f, err
}
