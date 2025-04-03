package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"im/edge/internal/config"
	"im/rpc_im/client/imservice"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
	ImRpc  imservice.ImService
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := zrpc.MustNewClient(c.ImRpc)

	// redis
	//conf := redis.RedisConf{
	//	Host: c.RedisConf.Host,
	//	Type: c.RedisConf.Type,
	//	Pass: c.RedisConf.Pass,
	//}
	//r, _ := redis.NewRedis(conf)

	return &ServiceContext{
		//Redis:  r,
		Config: c,
		ImRpc:  imservice.NewImService(client),
	}
}
