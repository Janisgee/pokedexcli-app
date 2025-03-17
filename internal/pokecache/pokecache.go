package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mux   *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)

	return c

}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()

	defer c.mux.Unlock()
	c.cache[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()

	defer c.mux.Unlock()
	// Check if map included "key"
	data, exist := c.cache[key]
	return data.val, exist
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, val := range c.cache {
		if val.createdAt.Before(now.Add(-last)) {
			delete(c.cache, key)
		}
	}
}
