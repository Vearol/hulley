package main

import (
  "fmt"
  "bytes"
  "io/ioutil"
  "../common"
  "net/http"
  "encoding/json"
)

func main() {
  events := &common.EventSlice{
    Events: make([]common.Event, 4),
  }

  eventsBuffer, _ := json.Marshal(events)

  url := "http://localhost:8080/postevents"
  
  request, err := http.NewRequest("POST", url, bytes.NewBuffer(eventsBuffer))
  request.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    panic(err)
  }
  
  defer response.Body.Close()

  fmt.Println("response Status:", response.Status)
  fmt.Println("response Headers:", response.Header)
  body, _ := ioutil.ReadAll(response.Body)
  fmt.Println("response Body:", string(body))
}
