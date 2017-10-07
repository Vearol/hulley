package main

import (
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
  } else {
    for _, element := range events.Events {
      eventData := element.ToString()
      common.Log.Info(eventData)
    }
  } 

  defer request.Body.Close()
}


func main() {
  
  common.InitLog("cruckserver.log")
  
  http.HandleFunc("/postevents", posteventsHandler)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

