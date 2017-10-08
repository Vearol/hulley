package main

import (
  "net/http"
  "encoding/json"
  "../common"
  "flag"
)

var (
  stdoutFlag = flag.Bool("stdout", false, "Log to stdout and to logfile")
)

func main() {
  err := parseFlags()
  if err != nil {
    flag.PrintDefaults()
  }

  logFilepaths := []string{ "cruckserver.log" }

  if (*stdoutFlag){
    logFilepaths = append(logFilepaths, "stdout")
  }
  
  common.InitLog(logFilepaths)
  if (err != nil){
    common.Log.Error(err.Error())
  }
  
  http.HandleFunc("/postevents", posteventsHandler)

  err = http.ListenAndServe(":8080", nil)
  if err != nil {
    common.Log.Error(err.Error())
  }  
}

func parseFlags() error {
  flag.Parse()
  return nil
}


func posteventsHandler(rw http.ResponseWriter, request *http.Request) {
  decoder := json.NewDecoder(request.Body)

  var events common.EventSlice

  err := decoder.Decode(&events)
  if err != nil {
    common.Log.Error(err.Error())
  } else {
    for _, element := range events.Events {
      eventData := element.String()
      common.Log.Info(eventData)
    }
  } 

  defer request.Body.Close()
}
