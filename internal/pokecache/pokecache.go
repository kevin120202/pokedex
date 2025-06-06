// Package pokecache provides a caching mechanism for PokeAPI responses
package pokecache

import (
	"sync"
	"time"
)

// Cache represents a thread-safe cache with automatic cleanup
// It stores API responses and removes them after a specified duration
type Cache struct {
	// cache is a map that stores the cached data
	// key: URL or identifier for the cached data
	// value: cacheEntry containing the data and when it was cached
	cache map[string]cacheEntry

	// mux is a mutex (mutual exclusion lock) that ensures
	// only one goroutine can access the cache at a time
	// This prevents race conditions when multiple goroutines
	// try to read or write to the cache simultaneously
	mux *sync.Mutex
}

// cacheEntry represents a single cached item
type cacheEntry struct {
	// createdAt stores when this entry was added to the cache
	// Used to determine if the entry is too old and should be removed
	createdAt time.Time

	// val stores the actual cached data as a byte slice
	// This allows us to cache any type of data
	val []byte
}

// NewCache creates and returns a new Cache instance
// interval: how often the cache should check for and remove old entries
func NewCache(interval time.Duration) Cache {
	// Create a new Cache with an empty map and a new mutex
	c := Cache{
		cache: make(map[string]cacheEntry), // Initialize empty map
		mux:   &sync.Mutex{},               // Initialize mutex
	}

	// Start a goroutine that periodically cleans up old entries
	// This runs in the background and doesn't block the main program
	go c.reapLoop(interval)

	return c
}

// Add stores a new value in the cache
// key: identifier for the cached data
// value: the data to cache
func (c *Cache) Add(key string, value []byte) {
	// Lock the mutex to ensure thread-safe access
	// Other goroutines will wait here if they try to access the cache
	c.mux.Lock()
	// Ensure the mutex is unlocked when the function returns
	// This is important to prevent deadlocks
	defer c.mux.Unlock()

	// Store the new entry in the cache
	// Include the current time so we know when it was added
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(), // Use UTC for consistency
		val:       value,            // Store the actual data
	}
}

// Get retrieves a value from the cache
// Returns the value and whether it was found
func (c *Cache) Get(key string) ([]byte, bool) {
	// Lock the mutex for thread-safe access
	c.mux.Lock()
	// Ensure the mutex is unlocked when the function returns
	defer c.mux.Unlock()

	// Try to get the value from the cache
	// ok will be true if the key exists, false otherwise
	val, ok := c.cache[key]
	return val.val, ok
}

// reapLoop runs in the background and periodically cleans up old entries
// interval: how often to check for old entries
func (c *Cache) reapLoop(interval time.Duration) {
	// Create a ticker that triggers at the specified interval
	ticker := time.NewTicker(interval)
	// Loop forever, waiting for the ticker to trigger
	for range ticker.C {
		// When the ticker triggers, clean up old entries
		c.reap(time.Now().UTC(), interval)
	}
}

// reap removes entries that are older than the specified interval
// now: current time
// last: how old an entry can be before it's removed
func (c *Cache) reap(now time.Time, last time.Duration) {
	// Lock the mutex for thread-safe access
	c.mux.Lock()
	// Ensure the mutex is unlocked when the function returns
	defer c.mux.Unlock()

	// Check each entry in the cache
	for k, v := range c.cache {
		// If the entry is older than the interval, remove it
		// now.Add(-last) gives us the cutoff time
		// If the entry was created before this time, it's too old
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k) // Remove the old entry
		}
	}
}
