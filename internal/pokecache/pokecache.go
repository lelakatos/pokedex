package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CachedValues map[string]cacheEntry
	mu           sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		CachedValues: make(map[string]cacheEntry),
		mu:           sync.Mutex{},
	}
	go c.reapLoop(interval)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.CachedValues[key] = newEntry

	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.CachedValues[key]

	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for tickerTime := range ticker.C {
		for k, v := range c.CachedValues {
			if diff := tickerTime.Sub(v.createdAt); diff > interval {
				c.mu.Lock()
				delete(c.CachedValues, k)
				c.mu.Unlock()
			}
		}
	}
}
