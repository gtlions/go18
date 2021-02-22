package ebi

import (
	"sync"

	"github.com/gtlions/gos10i"
)

type BodyMap map[string]interface{}

var mu sync.RWMutex

func (bm BodyMap) Set(key string, value interface{}) {
	mu.Lock()
	bm[key] = value
	mu.Unlock()
}
func (bm BodyMap) Get(key string) string {
	if bm == nil {
		return ""
	}
	mu.RLock()
	defer mu.RUnlock()
	value, ok := bm[key]
	if !ok {
		return ""
	}
	v, ok := value.(string)
	if !ok {
		return gos10i.X2String(value)
	}
	return v
}
func (bm BodyMap) Remove(key string) {
	mu.Lock()
	delete(bm, key)
	mu.Unlock()
}
