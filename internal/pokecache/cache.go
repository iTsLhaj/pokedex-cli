package pokecache

import (
	"time"
)

func (c *cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, ok
}

func (c *cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			{
				c.mu.Lock()
				for key, entry := range c.entries {
					if entry.createdAt.Add(interval).Before(time.Now()) {
						delete(c.entries, key)
					}
				}
				c.mu.Unlock()
			}
		}
	}
}

func NewCache() *ConfigurableCache {
	cc := &ConfigurableCache{
		Cache:    cache{},
		Interval: time.Minute,
	}
	cc.Cache.entries = make(map[string]cacheEntry)
	go cc.Cache.reapLoop(cc.Interval)
	return cc
}
