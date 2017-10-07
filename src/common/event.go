package common

import (
  "strings"
  "strconv"
)

type Event struct {
  Timestamp int `json:"timestamp"`
  EventId int `json:"eventId"`
  UserId int `json:"userId"`
}

type EventSlice struct {
  Events []Event `json:"events"`
}

func (event *Event) String() string {
  timestamp := strconv.Itoa(event.Timestamp)
  eventId := strconv.Itoa(event.EventId)
  userId := strconv.Itoa(event.UserId)

  result := strings.Join([]string { timestamp, eventId, userId }, "; ")
  return result
}

