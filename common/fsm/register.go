package fsm

import "context"

var (
	stateMachineMap = make(map[BusinessName]map[State]map[State]*TransitionEntry)
)

// 注册状态转移过程
func Register(bussiness BusinessName, currentStateInt64 int64, nextStateInt64 int64, handler ExecHandler) {

	currentState := State(currentStateInt64)
	nextState := State(nextStateInt64)
	if stateMachineMap[bussiness] == nil {
		stateMachineMap[bussiness] = make(map[State]map[State]*TransitionEntry)
	}

	if stateMachineMap[bussiness][currentState] == nil {
		stateMachineMap[bussiness][currentState] = make(map[State]*TransitionEntry)
	}

	if stateMachineMap[bussiness][currentState][nextState] == nil {
		stateMachineMap[bussiness][currentState][nextState] = &TransitionEntry{
			Handler: handler,
		}
	}
}

// 定义什么也不做的函数
func DoNothing(ctx context.Context, req interface{}) (interface{}, error) {

	return nil, nil
}
