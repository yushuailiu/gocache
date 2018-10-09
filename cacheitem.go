package gocache

import (
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex
	key interface{}
	data interface{}

	lifeSpan time.Duration

	createdOn time.Time
	accessedOn time.Time

	accessCount int64

	aboutToExpire func(key interface{})
}

func NewCacheItem(key interface{}, data interface{}, lifeSpan time.Duration) *CacheItem {
	t := time.Now()
	return &CacheItem{
		key: key,
		data: data,
		lifeSpan: lifeSpan,

		createdOn: t,
		accessedOn: t,
		aboutToExpire: nil,
	}
}

func (item *CacheItem) KeepAlive() {
	item.Lock()
	defer item.Unlock()
	item.accessedOn = time.Now()
	item.accessCount++
}

func (item *CacheItem) LifeSpan() time.Duration {
	return item.lifeSpan
}

func (item *CacheItem) AccessedOn() time.Time {
	item.RLock()
	defer item.RUnlock()
	return item.accessedOn
}

func (item *CacheItem) CreatedOn() time.Time {
	return item.createdOn
}

func (item *CacheItem) AccessCount() int64 {
	item.RLock()
	defer item.RUnlock()
	return item.accessCount
}

func (item *CacheItem) Key() interface{} {
	return item.key
}

func (item *CacheItem) Data() interface{} {
	return item.data
}

func (item *CacheItem) SetAboutToExpireCallback(f func(interface{})) {
	item.Lock()
	defer item.Unlock()
	item.aboutToExpire = f
}


