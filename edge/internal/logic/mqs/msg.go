package mqs

//type MsgLogic struct {
//	ctx    context.Context
//	svcCtx *svc.ServiceContext
//	server *socket.Server
//	logx.Logger
//}
//
//func NewMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgLogic {
//	return &MsgLogic{
//		ctx:    ctx,
//		svcCtx: svcCtx,
//		Logger: logx.WithContext(ctx),
//	}
//}
//
//func (l *MsgLogic) Consume(_, val string) error {
//	logx.Infof("[Consume] succ val: %s", val)
//	return nil
//}
