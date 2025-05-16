package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	m			map[string]cacheEntry
	mu			sync.Mutex
	interval	time.Duration
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		m:			map[string]cacheEntry{},
		mu:			sync.Mutex{},
		interval: 	interval,
	}
	go cache.reapLoop()
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.m[key] = cacheEntry{
		createdAt:	time.Now(),
		val:		val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool){
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cEntry, exists := cache.m[key]
	if !exists {
		return []byte{}, false
	}
	return cEntry.val, true
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	for {
		currentTick := <-ticker.C
		cache.mu.Lock()
		keysToDelete := []string{}
		for key := range cache.m {
			deltaCreated := currentTick.Sub(cache.m[key].createdAt)
			if deltaCreated > cache.interval {
				keysToDelete = append(keysToDelete, key)
			}
		}
		cache.mu.Unlock()
		cache.mu.Lock()
		for _, key := range keysToDelete {
			delete(cache.m, key)
		}
		cache.mu.Unlock()
	}
}