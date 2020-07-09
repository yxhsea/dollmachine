// Package memstore 基于内存的k-v存储
package memstore

import (
	"sync"
	"time"
)

// CacheStrategy 缓存策略
type CacheStrategy uint8

// 缓存策略
const (
	CacheStrategyAbsolute CacheStrategy = iota
	CacheStrategyRelative
)

type value struct {
	key      string
	v        interface{}
	expired  int64 // 失效时间
	d        time.Duration
	strategy CacheStrategy
}

// MemStore 内存存储
type MemStore struct {
	lock        *sync.RWMutex
	store       map[string]value
	gcFrequency time.Duration
}

// NewMemStore creates new instance
func NewMemStore() *MemStore {
	m := &MemStore{
		store:       make(map[string]value),
		lock:        &sync.RWMutex{},
		gcFrequency: time.Second * 30,
	}
	go m.gc()
	return m
}

// SetGCFrequency 设置缓存数据回收频率
func (m *MemStore) SetGCFrequency(d time.Duration) {
	m.gcFrequency = d
}

// Set 默认保存24小时
func (m *MemStore) Set(k string, v interface{}) {
	m.SetV1(k, v, CacheStrategyAbsolute, time.Hour*24)
}

// SetV1 sets value
func (m *MemStore) SetV1(k string, v interface{}, strategy CacheStrategy, d time.Duration) {
	m.lock.Lock()
	defer m.lock.Unlock()
	val := value{key: k, v: v, strategy: strategy, d: d, expired: time.Now().Add(d).Unix()}
	m.store[k] = val
}

// Get gets value
func (m *MemStore) Get(k string) (interface{}, bool) {
	m.lock.RLock()
	v, has := m.store[k]
	m.lock.RUnlock()
	if !has {
		return nil, false
	}
	if v.strategy == CacheStrategyRelative {
		m.lock.Lock()
		v.expired = time.Now().Add(v.d).Unix()
		m.store[k] = v
		m.lock.Unlock()
	}
	return v.v, true
}

func (m *MemStore) gc() {
	ticker := time.Tick(m.gcFrequency)
	for now := range ticker {
		nowUnix := now.Unix()
		var expiredKeys []string
		maxCount := 50
		i := 0
		for k, v := range m.store {
			if v.expired <= nowUnix {
				expiredKeys = append(expiredKeys, k)
			}
			i++
			if i >= maxCount {
				break
			}
		}
		m.lock.Lock()
		for index := range expiredKeys {
			delete(m.store, expiredKeys[index])
		}
		m.lock.Unlock()
	}
}
