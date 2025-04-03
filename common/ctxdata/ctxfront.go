package ctxdata

import (
	"context"
	"strconv"

	"github.com/goinggo/mapstructure"
	"google.golang.org/grpc/metadata"
)

var (
	FrontJwtUserId   = "frontJwtUserId"
	FrontJwtStoreId  = "frontJwtStoreId"
	FrontJwtMsId     = "frontJwtMsId"
	FrontJwtPhone    = "frontJwtPhone"
	FrontJwtUserType = "frontJwtUserType"
)

type claimFront struct {
	FrontJwtUserId   []string `json:"frontJwtUserId"`
	FrontJwtStoreId  []string `json:"frontJwtStoreId"`
	FrontJwtMsId     []string `json:"frontJwtMsId"`
	FrontJwtPhone    []string `json:"frontJwtPhone"`
	FrontJwtUserType []string `json:"frontJwtUserType"`
}

func GetFrontUidFromCtx(ctx context.Context) int64 {
	md, _ := metadata.FromOutgoingContext(ctx)
	var claim claimFront
	_ = mapstructure.Decode(md, &claim)

	parseInt, _ := strconv.ParseInt(claim.FrontJwtUserId[0], 10, 64)

	return parseInt
}

func GetFrontStoreIdFromCtx(ctx context.Context) int64 {
	md, _ := metadata.FromOutgoingContext(ctx)
	var claim claimFront
	_ = mapstructure.Decode(md, &claim)

	parseInt, _ := strconv.ParseInt(claim.FrontJwtStoreId[0], 10, 64)

	return parseInt
}

func GetFrontMsIdFromCtx(ctx context.Context) int64 {
	md, _ := metadata.FromOutgoingContext(ctx)
	var claim claimFront
	_ = mapstructure.Decode(md, &claim)

	parseInt, _ := strconv.ParseInt(claim.FrontJwtMsId[0], 10, 64)

	return parseInt
}

func GetFrontPhoneFromCtx(ctx context.Context) string {
	md, _ := metadata.FromOutgoingContext(ctx)
	return md.Get(FrontJwtPhone)[0]
}

func GetFrontUserTypeFromCtx(ctx context.Context) int32 {
	md, _ := metadata.FromOutgoingContext(ctx)
	var claim claimFront
	_ = mapstructure.Decode(md, &claim)

	parseInt, _ := strconv.ParseInt(claim.FrontJwtUserType[0], 10, 64)
	return int32(parseInt)
}

func FrontUpdateMetadata(ctx context.Context, data map[string]string) context.Context {
	// 从原始上下文中提取元数据
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	// 更新元数据
	for key, value := range data {
		md[key] = append(md[key], value)
	}

	// 返回更新后的上下文
	return metadata.NewOutgoingContext(ctx, md)
}
