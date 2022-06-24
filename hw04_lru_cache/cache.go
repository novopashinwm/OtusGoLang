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

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	ci := &cacheItem{key: key, value: value}
	if item, ok := cache.items[key]; ok {
		item.Value = ci
		cache.queue.MoveToFront(item)
		return true
	}

	if cache.queue.Len() == cache.capacity {
		lastItem, ok := cache.queue.Back().Value.(*cacheItem)
		if !ok {
			return false
		}
		cache.queue.Remove(cache.queue.Back())
		delete(cache.items, lastItem.key)
	}
	pushFront := cache.queue.PushFront(ci)
	cache.items[key] = pushFront
	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	if item, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(item)
		ci, ok := item.Value.(*cacheItem)
		if !ok {
			return nil, false
		}
		return ci.value, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}

type cacheItem struct {
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
