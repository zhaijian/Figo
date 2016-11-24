package Figo

import "github.com/quexer/utee"

type Cache interface {
	Put(key, val interface{})
	Get(key interface{}) interface{}
}

type CacheObj struct {
	put func(key, val interface{})
	get func(key interface{}) interface{}
}

func (p *CacheObj) Put(key, val interface{}) {
	p.put(key, val)
}

func (p *CacheObj) Get(key interface{}) interface{} {
	return p.get(key)
}

func NewCacheObj(put func(key, val interface{}), get func(key interface{}) interface{}) *CacheObj {
	return &CacheObj{
		put: put,
		get: get,
	}
}

func NewTimerCache(ttl int, expireCb func(key, value interface{})) *CacheObj {
	tc := utee.NewTimerCache(ttl, expireCb)
	put := func(key, val interface{}) {
		tc.Put(key, val)
	}
	get := func(key interface{}) interface{} {
		return tc.Get(key)
	}
	return NewCacheObj(put, get)
}