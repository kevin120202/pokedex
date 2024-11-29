package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {

	c := &Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, err := c.cache[key]
	if !err {
		return nil, false
	}

	return val.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()

		for key, value := range c.cache {
			if time.Since(value.createdAt) > interval {
				delete(c.cache, key)
			}
		}

		c.mu.Unlock()
	}
}
