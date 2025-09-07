package store

import (
	"sync"

	"github.com/Archii0/lumo/pkg/logger"
	"go.uber.org/zap"
)

var (
	mu    sync.RWMutex
	store = make(map[string]string)
)

func Set(key, value string) {
	mu.Lock()
	defer mu.Unlock()
	store[key] = value

	logger.Log.Info("Set key in store",
		zap.String("key", key),
		zap.String("value", value),
	)
}

func Get(key string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()
	val, ok := store[key]

	logger.Log.Info("Get key from store",
		zap.String("key", key),
		zap.Bool("found", ok),
	)

	return val, ok
}

func Delete(key string) {
	mu.Lock()
	defer mu.Unlock()
	delete(store, key)

	logger.Log.Info("Delete key from store",
		zap.String("key", key),
	)
}
