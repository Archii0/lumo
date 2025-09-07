package store

import "sync"

var (
	mu    sync.RWMutex
	store = make(map[string]string)
)

func Set(key, value string) {
	mu.Lock()
	defer mu.Unlock()
	store[key] = value
}

func Get(key string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()
	val, ok := store[key]
	return val, ok
}

func Delete(key string) {
	mu.Lock()
	defer mu.Unlock()
	delete(store, key)
}
