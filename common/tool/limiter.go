package tool

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity     int
	tokens       int
	fillInterval time.Duration
	mutex        sync.Mutex
	ticker       *time.Ticker
}

func NewTokenBucket(capacity int, fillInterval time.Duration) *TokenBucket {
	bucket := &TokenBucket{
		capacity:     capacity,
		tokens:       capacity,
		fillInterval: fillInterval,
		ticker:       time.NewTicker(fillInterval),
	}

	go bucket.refill()

	return bucket
}

func (tb *TokenBucket) refill() {
	for range tb.ticker.C {
		tb.mutex.Lock()
		if tb.tokens < tb.capacity {
			tb.tokens++
		}
		tb.mutex.Unlock()
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}
