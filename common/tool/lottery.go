package tool

import (
	"context"
	"time"
)

type LotteryStrategy interface {
	Run() error
}

type TimeLotteryStrategy struct {
	*AnnounceLotteryLogic
	CurrentTime time.Time
}

type PeopleLotteryStrategy struct {
	*AnnounceLotteryLogic
	CurrentTime time.Time
}

type Winner struct {
	LotteryId int64
	UserId    int64
	PrizeId   int64
}

type AnnounceLotteryLogic struct {
}

// []*model.Prize
func (l *AnnounceLotteryLogic) DrawLottery(ctx context.Context, lotteryId int64, prizes []float64, participantor []int64) ([]Winner, error) {
	// 随机选择中奖者
	//rand.New(rand.NewSource(time.Now().UnixNano()))
	//
	//// 获取奖品总数 = 中奖人数
	//var WinnersNum int64
	//for _, p := range prizes {
	//	WinnersNum += p.Count
	//}

	//winners := make([]Winner, 0)
	//
	//records, err := l.svcCtx.ClockTaskRecordModel.GetClockTaskRecordByLotteryIdAndUserIds(lotteryId, participantor)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}
