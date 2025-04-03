package deferMq

import (
	"context"
	"im/edge/internal/svc"
)

type AsynqTask struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsynqTask(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqTask {
	return &AsynqTask{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
//
//func (l *AsynqTask) Start() {
//
//	logx.Infof("AsynqTask start")
//
//	srv := asynq.NewServer(
//		asynq.RedisClientOpt{Addr: l.svcCtx.Config.RedisConf.Host, Password: l.svcCtx.Config.RedisConf.Pass},
//		asynq.Config{
//			Concurrency: 10,
//			Queues: map[string]int{
//				"critical": 6,
//				"default":  3,
//				"low":      1,
//			},
//		},
//	)
//
//	mux := asynq.NewServeMux()
//
//	//关闭民宿订单任务
//	//mux.HandleFunc(asynqmq.TypeHomestayOrderCloseDelivery, l.closeHomestayOrderStateMqHandler)
//
//	if err := srv.Run(mux); err != nil {
//		logx.Errorf("could not run server: %v", err)
//	}
//}
//
//func (l *AsynqTask) Stop() {
//	logx.Infof("AsynqTask stop")
//}
