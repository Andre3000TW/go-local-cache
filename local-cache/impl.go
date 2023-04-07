package localcache

import (
	"sync"
	"time"
)

const duration = 30 * time.Second

type cache struct {
	mtx   sync.Mutex
	items map[string]*cacheItem
}

type cacheItem struct {
	val   interface{}
	timer *time.Timer
}

func (c *cache) Get(key string) interface{} {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if item, ok := c.items[key]; ok {
		return item.val
	}

	return nil
}

func (c *cache) Set(key string, val interface{}) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if item, ok := c.items[key]; ok {
		item.timer.Stop()
	}

	c.items[key] = &cacheItem{
		val: val,
		timer: time.AfterFunc(duration, func() {
			c.mtx.Lock()
			defer c.mtx.Unlock()

			delete(c.items, key)
		}),
	}
}

func New() Cache {
	return &cache{
		mtx:   sync.Mutex{},
		items: make(map[string]*cacheItem),
	}
}
