package storage

import "time"

type Event struct {
	ID              string
	Title           string
	BeginTimestamp  time.Time
	EndTimestamp    time.Time
	Description     string
	UserID          string
	RemindTimestamp time.Time
}

type storage interface {
	CreateEvent(event Event) error
	GetEventByID(id string) Event
	UpdateEventByID(event Event) error
	DeleteEventByID(id string) error
	GetAllEvents() []Event
}

type Storage struct {
	storage
}
