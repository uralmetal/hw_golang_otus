package app

import (
	"context"
	"github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/logger"
	"github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/storage"
)

type App struct { // TODO
	storage storage.Storage
	logger  logger.Logger
}

//type Logger interface { // TODO
//}
//
//type Storage interface { // TODO
//	memorystorage.MemoryStorage
//}

func New(logger logger.Logger, storage storage.Storage) *App {
	return &App{
		logger:  logger,
		storage: storage,
	}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	defer ctx.Done()
	return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

func (a *App) UpdateEvent(event storage.Event) error {
	return a.storage.UpdateEventByID(event)
}

func (a *App) DeleteEvent(id string) error {
	return a.storage.DeleteEventByID(id)
}

func (a *App) GetEventByID(id string) storage.Event {
	return a.storage.GetEventByID(id)
}

//func (a *App) GetEventsByTimeRange(beginTimeRange string, endTimeRange string) []storage.Event {
//	return []storage.Event{}
//}
