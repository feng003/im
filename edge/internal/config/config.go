package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf

	TCPName      string
	TCPListenOn  string
	WSListenOn   string
	SendChanSize int

	//RedisConf struct {
	//	Host string
	//	Pass string
	//	Type string
	//}

	KqConf kq.KqConf
	//MqConf kq.KqConf
	Etcd discov.EtcdConf

	ImRpc zrpc.RpcClientConf
}
