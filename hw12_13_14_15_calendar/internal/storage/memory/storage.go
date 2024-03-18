package memorystorage

import (
	"github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/storage"
	"sync"
)

type MemoryStorage struct {
	storage.Storage
	storage map[string]storage.Event
	mu      sync.RWMutex //nolint:unused
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (memoryStorage *MemoryStorage) CreateEvent(event storage.Event) error {
	defer memoryStorage.mu.Unlock()
	memoryStorage.mu.Lock()
	memoryStorage.storage[event.ID] = event
	return nil
}

func (memoryStorage *MemoryStorage) GetEventByID(id string) storage.Event {
	defer memoryStorage.mu.Unlock()
	memoryStorage.mu.Lock()
	return memoryStorage.storage[id]
}

func (memoryStorage *MemoryStorage) UpdateEventByID(event storage.Event) error {
	defer memoryStorage.mu.Unlock()
	memoryStorage.mu.Lock()
	memoryStorage.storage[event.ID] = event
	return nil
}

func (memoryStorage *MemoryStorage) DeleteEventByID(id string) error {
	defer memoryStorage.mu.Unlock()
	memoryStorage.mu.Lock()
	delete(memoryStorage.storage, id)
	return nil
}

func (memoryStorage *MemoryStorage) GetAllEvents() []storage.Event {
	events := make([]storage.Event, len(memoryStorage.storage))
	for _, value := range memoryStorage.storage {
		events = append(events, value)
	}
	return events
}
