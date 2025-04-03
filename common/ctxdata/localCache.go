package ctxdata

import (
	"sync"
	"time"
)

// 定义一个全局本地缓存
type LocalCache struct {
	sync.RWMutex
	data map[string]cacheItem
}

// 缓存项定义，包含值和过期时间
type cacheItem struct {
	value     string
	expiresAt time.Time
}

// 初始化本地缓存
var localCache = LocalCache{
	data: make(map[string]cacheItem),
}

// 添加到本地缓存
func (lc *LocalCache) Add(userID, token string, duration time.Duration) {
	lc.Lock()
	defer lc.Unlock()
	lc.data[userID] = cacheItem{
		value:     token,
		expiresAt: time.Now().Add(duration),
	}
}

// 从本地缓存获取
func (lc *LocalCache) Get(userID string) (string, bool) {
	lc.RLock()
	defer lc.RUnlock()

	// 检查缓存是否存在以及是否过期
	item, exists := lc.data[userID]
	if !exists || time.Now().After(item.expiresAt) {
		return "", false
	}
	return item.value, true
}

// 从本地缓存删除
func (lc *LocalCache) Delete(userID string) {
	lc.Lock()
	defer lc.Unlock()
	delete(lc.data, userID)
}

// 定期清理过期缓存
func (lc *LocalCache) CleanupExpiredItems() {
	for {
		time.Sleep(1 * time.Minute) // 每分钟清理一次
		lc.Lock()
		for userID, item := range lc.data {
			if time.Now().After(item.expiresAt) {
				delete(lc.data, userID)
			}
		}
		lc.Unlock()
	}
}
