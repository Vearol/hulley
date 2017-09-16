package common

type Event struct {
  Timestamp int `json:"timestamp"`
  EventId int `json:"eventId"`
  UserId int `json:"userId"`
}

type EventSlice struct {
  Events []Event `json:"events"`
}
