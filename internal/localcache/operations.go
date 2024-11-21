package localcache

import (
	"time"

	"github.com/google/uuid"
)

// Add to local cache
func (lc *LocalCache) Add(id uuid.UUID, obj interface{}) {
	lc.cache[id] = &CachedObject{
		expire: time.Now().Add(3 * time.Hour),
		object: obj,
	}
}

// Get from local cache
func (lc *LocalCache) Get(id uuid.UUID) interface{} {
	obj, ok := lc.cache[id]
	if ok {
		return obj.object
	}
	return nil
}

// Get from local cache
func (lc *LocalCache) Remove(id uuid.UUID) interface{} {
	_, ok := lc.cache[id]
	if ok {
		delete(lc.cache, id)
	}
	return nil
}
