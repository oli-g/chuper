package chuper

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("not found")
)

type Cache interface {
	Get(key string) (interface{}, error)

	Set(key string, value interface{}) error

	SetNX(key string, value interface{}) (bool, error)

	Delete(key string) error
}

type MemoryCache struct {
	sync.Mutex

	items map[string]interface{}
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		items: map[string]interface{}{},
	}
}

func (r *MemoryCache) Get(key string) (interface{}, error) {
	r.Lock()
	defer r.Unlock()

	value, ok := r.items[key]
	if !ok {
		return nil, ErrNotFound
	}

	return value, nil
}

func (r *MemoryCache) Set(key string, value interface{}) error {
	r.Lock()
	defer r.Unlock()

	r.items[key] = value
	return nil
}

func (r *MemoryCache) SetNX(key string, value interface{}) (bool, error) {
	r.Lock()
	defer r.Unlock()

	_, ok := r.items[key]
	if ok {
		return false, nil
	} else {
		r.items[key] = value
		return true, nil
	}
}

func (r *MemoryCache) Delete(key string) error {
	r.Lock()
	defer r.Unlock()

	delete(r.items, key)
	return nil
}
