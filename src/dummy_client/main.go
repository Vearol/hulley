package main

import (
  "fmt"
  "bytes"
  "io/ioutil"
  "../common"
  "net/http"
  "encoding/json"
  "math/rand"
  "time"
)

func main() {
  events := &common.EventSlice{
    Events: make([]common.Event, 2),
  }

  eventsBuffer, _ := json.Marshal(events)

  url := "http://localhost:8080/postevents"


  randomTime := rand.Intn(3) + 1
  timer := time.NewTimer(time.Second * time.Duration(randomTime))
  for {
    
    <-timer.C
    
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
    
    randomTime = rand.Intn(3) + 1
    timer = time.NewTimer(time.Second * time.Duration(randomTime))
  }
}
