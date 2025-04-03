package fsm

import (
	"context"
	"fmt"
)

type Fsm struct {

	// 业务名
	Business BusinessName

	// 当前状态
	CurrentState State

	// 要转移到的状态
	NextState State
}

type TransitionEntry struct {

	// 执行状态转移的函数
	Handler ExecHandler
}

// 新建状态机实例，需要指定业务名

func NewFsm(business BusinessName) *Fsm {

	return &Fsm{

		Business: business,
	}

}

// 设置当前状态及要转移到的状态

func (f *Fsm) SetState(currentState, nextState int64) {

	f.setCurrentState(State(currentState))

	f.setNextState(State(nextState))

}

func (f *Fsm) setCurrentState(currentState State) {

	f.CurrentState = currentState

}

func (f *Fsm) setNextState(nextState State) {

	f.NextState = nextState

}

// 进行状态转移
func (f *Fsm) Transfer(ctx context.Context, param interface{}) (interface{}, error) {

	stateMap, ok := stateMachineMap[f.Business]
	if !ok {
		return nil, fmt.Errorf("fsm business %v invalid", f.Business)
	}

	nextMap, ok := stateMap[f.CurrentState]
	if !ok {

		return nil, fmt.Errorf("fsm currentState %v invalid, business = %v", f.CurrentState, f.Business)
	}

	transitionEntry, ok := nextMap[f.NextState]
	if !ok {
		return nil, fmt.Errorf("fsm nextState %v invalid, business = %v", f.NextState, f.Business)
	}

	return transitionEntry.Handler(ctx, param)
}
