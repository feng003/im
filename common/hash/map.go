package hash

import (
	"hash/maphash"
	"sync"
)

// Generic ShardingMap with type parameters K (key) and V (value)
type ShardingMap[K comparable, V any] struct {
	locks  []sync.RWMutex
	shards []map[K]V
	seed   maphash.Seed
}

// NewShardingMap creates a new generic ShardingMap with the specified number of shards.
func NewShardingMap[K comparable, V any](size int) *ShardingMap[K, V] {
	sm := &ShardingMap[K, V]{
		locks:  make([]sync.RWMutex, size),
		shards: make([]map[K]V, size),
		seed:   maphash.MakeSeed(),
	}
	for i := 0; i < size; i++ {
		sm.shards[i] = make(map[K]V)
	}
	return sm
}

// hashKey generates a hash for a key using maphash.
func (m *ShardingMap[K, V]) hashKey(key K) uint64 {
	var h maphash.Hash
	h.SetSeed(m.seed)
	h.WriteString(any(key).(string)) // Convert key to string if necessary
	return h.Sum64()
}

// getShardIdx determines which shard a key belongs to.
func (m *ShardingMap[K, V]) getShardIdx(key K) uint64 {
	hash := m.hashKey(key)
	return hash % uint64(len(m.shards))
}

// Set adds or updates a key-value pair in the map.
func (m *ShardingMap[K, V]) Set(key K, value V) {
	idx := m.getShardIdx(key)
	m.locks[idx].Lock()
	defer m.locks[idx].Unlock()
	m.shards[idx][key] = value
}

// Get retrieves the value for a key and a boolean indicating if the key exists.
func (m *ShardingMap[K, V]) Get(key K) (V, bool) {
	idx := m.getShardIdx(key)
	m.locks[idx].RLock()
	defer m.locks[idx].RUnlock()
	value, ok := m.shards[idx][key]
	return value, ok
}

// Del removes a key-value pair from the map.
func (m *ShardingMap[K, V]) Del(key K) {
	idx := m.getShardIdx(key)
	m.locks[idx].Lock()
	defer m.locks[idx].Unlock()
	delete(m.shards[idx], key)
}
