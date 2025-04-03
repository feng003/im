package svc

import (
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueList(t *testing.T) {
	// TODO
	// 初始化QueueList
	queueList := NewQueueList()

	// 定义测试数据
	key := "edge_01"
	conf := kq.KqConf{
		Brokers: []string{"127.0.0.1:9092"},
		Topic:   "topic-edge-01",
	}

	// Test Update
	queueList.Update(key, conf)
	pusher, ok := queueList.Load(key)
	assert.True(t, ok, "Expected key to exist after update")
	assert.NotNil(t, pusher, "Expected pusher to be non-nil after update")
	fmt.Printf("pusher: %+v ok is %+v \n", pusher, ok)

	// Test Load for non-existent key
	load, ok := queueList.Load("nonExistentKey")
	assert.False(t, ok, "Expected key to not exist")
	fmt.Printf("load: %+v\n", load)

	// Test Delete
	queueList.Delete(key)
	_, ok = queueList.Load(key)
	assert.False(t, ok, "Expected key to not exist after delete")
}
