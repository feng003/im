package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/threading"
	"im/common/discovery"
	"im/rpc_im/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	BizRedis  *redis.Redis
	QueueList *QueueList
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 第一次初始化
	queueList := GetQueueList(c.QueueEtcd)
	threading.GoSafe(func() {
		discovery.QueueDiscoveryProc(c.QueueEtcd, queueList)
	})
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:    c,
		BizRedis:  rds,
		QueueList: NewQueueList(),
	}
}
