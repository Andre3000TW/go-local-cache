package localcache

import (
	"sync"
	"time"
)

const expirationTime = 30 * time.Second

type cache struct {
	mtx sync.Mutex
	val map[string]any
}

func (c *cache) Get(key string) (any, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	val, ok := c.val[key]
	return val, ok
}

func (c *cache) Set(key string, val any) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.val[key] = val
	time.AfterFunc(expirationTime, func() {
		c.mtx.Lock()
		defer c.mtx.Unlock()

		delete(c.val, key)
	})
}

func New() Cache {
	return &cache{
		mtx: sync.Mutex{},
		val: make(map[string]any),
	}
}
