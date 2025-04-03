package fsm

import (
	"context"
	"testing"
)

func Test_Fsm(t *testing.T) {

	const activityBase BusinessName = "activity"
	var (
		Configuring State = 1
		Testing     State = 2
		CanOnline   State = 3

		submitToConfig ExecHandler = func(ctx context.Context, param interface{}) (interface{}, error) {
			return nil, nil
		}
		submitToTest ExecHandler = func(ctx context.Context, param interface{}) (interface{}, error) {
			return nil, nil
		}
		testDone ExecHandler = func(ctx context.Context, param interface{}) (interface{}, error) {
			return nil, nil
		}
	)

	// 配置中 -> 待上线
	Register(activityBase, int64(Configuring), int64(CanOnline), submitToConfig)
	// 配置中 -> 测试中
	Register(activityBase, int64(Configuring), int64(Testing), submitToTest)
	// 测试中 -> 待上线
	Register(activityBase, int64(Testing), int64(CanOnline), testDone)
}

func testDone(ctx context.Context, req interface{}) (interface{}, error) {

	// 检查参数类型
	//request, ok := req.(*testDoneFsmParam)
	//if !ok {
	//	return 0, ParamConvertInvalidError
	//}
	//fmt.Println(request)

	// 执行业务逻辑

	return nil, nil
}
