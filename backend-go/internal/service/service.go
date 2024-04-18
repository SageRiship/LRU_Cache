package service

import (
	"container/list"
	"sync"
	"time"
)

type cacheItem struct {
	key        string
	value      interface{}
	expiration time.Time
}

type LRUCache struct {
	maxSize   int
	evictList *list.List
	cache     map[string]*list.Element
	mutex     sync.Mutex
}

func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		maxSize:   maxSize,
		evictList: list.New(),
		cache:     make(map[string]*list.Element),
	}
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		item := elem.Value.(*cacheItem)
		if item.expiration.After(time.Now()) {
			c.evictList.MoveToFront(elem)
			return item.value, true
		} else {
			c.evictList.Remove(elem)
			delete(c.cache, key)
		}
	}
	return nil, false
}

func (c *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.evictList.MoveToFront(elem)
		elem.Value.(*cacheItem).value = value
		elem.Value.(*cacheItem).expiration = time.Now().Add(expiration)
	} else {
		elem := c.evictList.PushFront(&cacheItem{key, value, time.Now().Add(expiration)})
		c.cache[key] = elem
		if c.evictList.Len() > c.maxSize {
			c.removeOldest()
		}
	}
}

func (c *LRUCache) removeOldest() {
	if c.evictList.Len() == 0 {
		return
	}
	elem := c.evictList.Back()
	if elem != nil {
		c.evictList.Remove(elem)
		key := elem.Value.(*cacheItem).key
		delete(c.cache, key)
	}
}
