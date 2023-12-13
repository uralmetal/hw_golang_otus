package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    *sync.Mutex
}

type itemCache struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mutex:    new(sync.Mutex),
	}
}

func (cache lruCache) Set(key Key, value interface{}) bool {
	cache.mutex.Lock()
	item, exist := cache.items[key]
	if exist {
		cacheValue := item.Value.(itemCache)
		cacheValue.value = value
		item.Value = cacheValue
		cache.items[key] = item
		cache.queue.MoveToFront(item)
		defer cache.mutex.Unlock()
		return exist
	}
	if cache.queue.Len() >= cache.capacity {
		oldestItem := cache.queue.Back()
		delete(cache.items, oldestItem.Value.(itemCache).key)
		cache.queue.Remove(oldestItem)
	}
	item = cache.queue.PushFront(itemCache{
		key:   key,
		value: value,
	})
	cache.items[key] = item
	defer cache.mutex.Unlock()
	return exist
}

func (cache lruCache) Get(key Key) (interface{}, bool) {
	cache.mutex.Lock()
	item, exist := cache.items[key]
	if exist {
		cache.queue.MoveToFront(item)
		cacheValue := item.Value.(itemCache)
		defer cache.mutex.Unlock()
		return cacheValue.value, exist
	}
	defer cache.mutex.Unlock()
	return nil, exist
}

func (cache lruCache) Clear() {
	cache.mutex.Lock()
	for k := range cache.items {
		delete(cache.items, k)
	}
	for i := cache.queue.Front(); i != nil; i = i.Next {
		cache.queue.Remove(i)
	}
	defer cache.mutex.Unlock()
}
