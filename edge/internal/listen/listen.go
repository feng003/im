package listen

import (
	"context"
	"im/common/socket"
	"im/edge/internal/config"
	"im/edge/internal/svc"

	// "im/edge/internal/logic/deferMq"

	"github.com/zeromicro/go-zero/core/service"
)

func Mqs(c config.Config, server *socket.Server) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	//kq ：pub sub 消息队列.
	// services = append(services, KqMqs(c, ctx, svcContext, server)...)

	//asynq ： 延迟队列、定时任务
	services = append(services, AsynqMqs(c, ctx, svcContext)...)

	return services
}
