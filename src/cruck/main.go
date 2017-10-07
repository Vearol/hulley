package main

import (
  "net/http"
  "encoding/json"
  "../common"
)

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


func main() {  
  common.InitLog("cruckserver.log")
  
  http.HandleFunc("/postevents", posteventsHandler)

  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    common.Log.Error(err.Error())
  }  
}

