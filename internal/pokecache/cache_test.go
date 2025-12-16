package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	c := NewCache()

	key := "foo"
	val := []byte("bar")

	c.Cache.Add(key, val)

	got, ok := c.Cache.Get(key)
	if !ok {
		t.Fatalf("expected cache hit")
	}

	if string(got) != "bar" {
		t.Fatalf("expected 'bar', got '%s'", string(got))
	}
}

func TestCacheReapDeletesExpiredEntry(t *testing.T) {
	c := NewCache()

	// override interval for fast test
	c.Interval = 20 * time.Millisecond
	c.Cache.entries = make(map[string]cacheEntry)

	go c.Cache.reapLoop(c.Interval)

	key := "expire-me"
	c.Cache.Add(key, []byte("bye"))

	// wait longer than interval
	time.Sleep(50 * time.Millisecond)

	_, ok := c.Cache.Get(key)
	if ok {
		t.Fatalf("expected cache entry to be deleted")
	}
}
