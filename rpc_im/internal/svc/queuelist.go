package svc

import (
	"sync"
	"time"

	"context"
	"encoding/json"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type QueueList struct {
	kqs map[string]*kq.Pusher
	l   sync.Mutex
}

func NewQueueList() *QueueList {
	return &QueueList{
		kqs: make(map[string]*kq.Pusher),
	}
}

func (q *QueueList) Update(key string, data kq.KqConf) {
	logx.Infof("QueueList Update key: %+v data: %+v", key, data)
	edgeQueue := kq.NewPusher(data.Brokers, data.Topic)
	q.l.Lock()
	q.kqs[key] = edgeQueue
	q.l.Unlock()
}

func (q *QueueList) Delete(key string) {
	q.l.Lock()
	delete(q.kqs, key)
	q.l.Unlock()
}

func (q *QueueList) Load(key string) (*kq.Pusher, bool) {
	logx.Infof("QueueList Load key: %+v data: %+v", key, q.kqs)
	q.l.Lock()
	defer q.l.Unlock()

	edgeQueue, ok := q.kqs[key]
	return edgeQueue, ok
}

func GetQueueList(conf discov.EtcdConf) *QueueList {
	ql := NewQueueList()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.Hosts,
		DialTimeout: time.Second * 3,
	})
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cli.Get(ctx, conf.Key, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	logx.Infof("QueueList init key is %+v res: %+v", conf.Key, res)
	for _, kv := range res.Kvs {
		var data kq.KqConf
		if err := json.Unmarshal(kv.Value, &data); err != nil {
			logx.Errorf("invalid data key is: %s value is: %s", string(kv.Key), string(kv.Value))
			continue
		}
		if len(data.Brokers) == 0 || len(data.Topic) == 0 {
			continue
		}
		edgeQueue := kq.NewPusher(data.Brokers, data.Topic)

		ql.l.Lock()
		ql.kqs[string(kv.Key)] = edgeQueue
		ql.l.Unlock()
	}

	return ql
}
