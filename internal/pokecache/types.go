package pokecache

import (
	"sync"
	"time"
)

type (
	cacheEntry struct {
		createdAt time.Time
		val       []byte
	}

	cache struct {
		mu      sync.Mutex
		entries map[string]cacheEntry
	}

	ConfigurableCache struct {
		Cache    cache
		Interval time.Duration
	}
)
