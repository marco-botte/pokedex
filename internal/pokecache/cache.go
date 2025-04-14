package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       *sync.Mutex
	interval time.Duration
	store    map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	var store = make(map[string]cacheEntry)
	mu := &sync.Mutex{}
	cache := &Cache{
		mu:       mu,
		interval: interval,
		store:    store,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.mu.Lock()
		for key, entry := range c.store {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.store, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.store[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.store[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
