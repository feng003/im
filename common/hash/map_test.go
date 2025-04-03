package hash

import "testing"

func TestNewShardingMap(t *testing.T) {
	// Create a ShardingMap with string keys and int values
	sm := NewShardingMap[string, int](10)

	// Set values
	sm.Set("key1", 100)
	sm.Set("key2", 200)

	// Get values
	if value, ok := sm.Get("key1"); ok {
		println("key1:", value)
	} else {
		println("key1 not found")
	}

	// Delete a key
	sm.Del("key1")

	// Try to get the deleted key
	if _, ok := sm.Get("key1"); !ok {
		println("key1 deleted")
	}
}
