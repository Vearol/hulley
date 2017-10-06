package common

import (
  s "strings"
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

func (event Event) ToString() string{
  timestamp := strconv.Itoa(event.Timestamp)
  eventId := strconv.Itoa(event.EventId)
  userId := strconv.Itoa(event.UserId)

  return s.Join([]string{timestamp, eventId, userId}, "; ")
}

