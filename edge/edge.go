package main

import (
	"context"
	"flag"
	"fmt"

	"golang.org/x/net/websocket"
	"net/http"
	"im/common/socketio"

	"im/edge/internal/logic"

	"im/edge/internal/config"
	"im/edge/internal/server"
	"im/edge/internal/svc"

	"im/common/libnet"
	"im/common/socket"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/im.yaml", "the config file")

func main() {
	flag.Parse()

	var err error
	var c config.Config
	conf.MustLoad(*configFile, &c)
	srvCtx := svc.NewServiceContext(c)

	logx.DisableStat()

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	tcpServer := server.NewTCPServer(srvCtx)
	wsServer := server.NewWSServer(srvCtx)
	protobuf := libnet.NewIMProtocol()

	tcpServer.Server, err = socket.NewServe(c.TCPName, c.TCPListenOn, protobuf, 0)
	if err != nil {
		panic(err)
	}

	wsServer.Server, err = socketio.NewServe(c.Name, c.WSListenOn, protobuf, c.SendChanSize)
	if err != nil {
		panic(err)
	}
	http.Handle("/ws", websocket.Handler(func(conn *websocket.Conn) {
		conn.PayloadType = websocket.BinaryFrame
		wsServer.HandleRequest(conn)
	}))

	//threading.GoSafe(func() {
	//	tcpServer.HandleRequest()
	//})
	//threading.GoSafe(func() {
	//	tcpServer.KqHeart()
	//})
	go wsServer.Start()
	go tcpServer.HandleRequest()
	go tcpServer.KqHeart()

	fmt.Printf("Starting tcp server at %s, ws server at: %s...\n", c.TCPListenOn, c.WSListenOn)
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	//for _, mq := range listen.Mqs(c, tcpServer.Server) {
	//	serviceGroup.Add(mq)
	//}

	for _, mq := range logic.Consumers(context.Background(), srvCtx, tcpServer.Server, wsServer.Server) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()
}
