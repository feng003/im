package server

import (
	"github.com/zeromicro/go-zero/core/logx"
	"im/common/discovery"
	"im/common/socket"
	"im/edge/client"
	"im/edge/internal/svc"
)

type TCPServer struct {
	svcCtx *svc.ServiceContext
	Server *socket.Server
}

func NewTCPServer(svcCtx *svc.ServiceContext) *TCPServer {
	return &TCPServer{svcCtx: svcCtx}
}

func (srv *TCPServer) HandleRequest() {
	for {
		session, err := srv.Server.Accept()
		if err != nil {
			panic(err)
		}
		cli := client.NewClient(srv.Server.Manager, session, srv.svcCtx.ImRpc)
		go srv.sessionLoop(cli)
	}
}

func (srv *TCPServer) sessionLoop(client *client.Client) {
	message, err := client.Receive()
	if err != nil {
		logx.Errorf("[sessionLoop] client.Receive error: %v", err)
		_ = client.Close()
		return
	}

	logx.Infof("[sessionLoop] receive message: %+v", message)
	// 登录校验
	err = client.Login(message)
	if err != nil {
		logx.Errorf("[sessionLoop] client.Login error: %v", err)
		_ = client.Close()
		return
	}

	go func() {
		err := client.HeartBeat()
		if err != nil {
			logx.Errorf("[sessionLoop] client.HeartBeat error: %v", err)
		}
	}()

	for {
		message, err := client.Receive()
		logx.Infof("[sessionLoop] receive message: %+v and error is %v \n", message, err)
		if err != nil {
			logx.Errorf("[sessionLoop] client.Receive error: %v", err)
			_ = client.Close()
			return
		}
		err = client.HandlePackage(message)
		if err != nil {
			logx.Errorf("[sessionLoop] client.HandleMessage error: %v", err)
		}
	}
}

func (srv *TCPServer) KqHeart() {
	work := discovery.NewQueueWorker(srv.svcCtx.Config.Etcd.Key, srv.svcCtx.Config.Etcd.Hosts, srv.svcCtx.Config.KqConf)
	work.HeartBeat()
}
