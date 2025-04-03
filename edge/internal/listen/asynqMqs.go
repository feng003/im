package listen

import (
	"context"
	"im/edge/internal/config"
	"im/edge/internal/svc"

	// "im/edge/internal/logic/deferMq"

	"github.com/zeromicro/go-zero/core/service"
)

// asynq
// 定时任务、延迟任务
func AsynqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//监听延迟队列
		// deferMq.NewAsynqTask(ctx, svcContext),

		//监听定时任务
	}

}
