package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
		createdAt	time.Time
		value		[]byte
	}

func NewCache(interval time.Duration) Cache {
	cac := Cache{
		cache: 	make(map[string]cacheEntry),
		mutex:	&sync.Mutex{},
	}
	go cac.reapLoop(interval)
	return cac
}

func (cac *Cache) Add(key string, value []byte) {
	cac.mutex.Lock()
	defer cac.mutex.Unlock()
	cac.cache[key] = cacheEntry {
		createdAt: 	time.Now().UTC(),
		value: 		value,
	}
}

func (cac *Cache) Get(key string) ([]byte, bool) {
	cac.mutex.Lock()
	defer cac.mutex.Unlock()
	value, ok := cac.cache[key]
	return value.value, ok
}

func (cac *Cache) reapLoop(interval time.Duration) {
	tick := time.NewTicker(interval)
	for range tick.C {
		cac.reap(time.Now().UTC(), interval)
	}
}

func (cac *Cache) reap(current time.Time, previous time.Duration) {
	cac.mutex.Lock()
	defer cac.mutex.Unlock()
	for i, j := range cac.cache {
		if j.createdAt.Before(current.Add(-previous)) {
			delete(cac.cache, i)
		}
	}
}