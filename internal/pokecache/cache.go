package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Entries map[string]CacheEntry
	Mu      sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries: make(map[string]CacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Entries[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	entry, found := c.Entries[key]
	if !found {
		return nil, false
	}
	return entry.Val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for t := range ticker.C {
		c.Mu.Lock()
		for key, entry := range c.Entries {
			if t.Sub(entry.CreatedAt) > interval {
				delete(c.Entries, key)
			}
		}
		c.Mu.Unlock()
	}
}
