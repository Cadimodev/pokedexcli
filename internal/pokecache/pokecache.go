package pokecache

import (
	"errors"
	"time"
)

func NewCache(interval time.Duration) *Cache {

	c := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		interval:     interval,
		ticker:       time.NewTicker(interval),
	}

	go func() {
		for range c.ticker.C {
			c.reapLoop()
		}
	}()

	return c
}

func (c *Cache) Add(key string, val []byte) error {

	if len(key) == 0 {
		return errors.New("key is empty")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.cacheEntries[key]; ok {
		return errors.New("key already exist")
	}

	c.cacheEntries[key] = cacheEntry{createdAt: time.Now(), val: val}

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {

	if len(key) == 0 {
		return []byte{}, false
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.cacheEntries[key]; !ok {
		return []byte{}, false
	}

	return c.cacheEntries[key].val, true
}

func (c *Cache) reapLoop() {

	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	for key, entry := range c.cacheEntries {

		if now.Sub(entry.createdAt) > c.interval {
			delete(c.cacheEntries, key)
		}
	}
}
