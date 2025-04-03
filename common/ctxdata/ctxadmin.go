package ctxdata

import (
	"context"
	"strconv"
	"strings"

	"github.com/goinggo/mapstructure"
	"google.golang.org/grpc/metadata"
)

var (
	CtxKeyJwtMsId    = "jwtMsId"
	CtxKeyJwtUserId  = "jwtUserId"
	CtxKeyJwtStaffId = "jwtStaffId"
	CtxKeyJwtStoreId = "jwtStoreId"
	CtxKeyJwtShopId  = "jwtShopId"
	CtxIsSupperAdmin = "jwtIsSupperAdmin"
	CtxUserType      = "jwtUserType"
	CtxUserRole      = "jwtUserRole"
	CtxUserUuid      = "jwtUserUuid"
	CtxUserIsMask    = "jwtUserIsMask"
	CtxReferralCode  = "jwtReferralCode"
)

type claimRes struct {
	JwtMsId          []string `json:"jwtMsId"`
	JwtUserId        []string `json:"jwtUserId"`
	JwtStaffId       []string `json:"jwtStaffId"`
	JwtStoreId       []string `json:"jwtStoreId"`
	JwtShopId        []string `json:"jwtShopId"`
	JwtIsSupperAdmin []string `json:"jwtIsSupperAdmin"`
	JwtUserType      []string `json:"jwtUserType"`
	JwtUserRole      []string `json:"jwtUserRole"`
	JwtUserUuid      []string `json:"jwtUserUuid"`
	JwtUserIsMask    []string `json:"jwtUserIsMask"`
	JwtReferralCode  []string `json:"jwtReferralCode"`
}

// GetUidFromCtx 获取用户编号
func GetUidFromCtx(ctx context.Context) int64 {
	md, _ := metadata.FromOutgoingContext(ctx)

	var claim claimRes
	_ = mapstructure.Decode(md, &claim)
	// 如果没有返回 0
	if len(claim.JwtUserId) == 0 {
		return 0
	}

	parseInt, _ := strconv.ParseInt(claim.JwtUserId[0], 10, 64)
	return parseInt
}

// GetStaffIdFromCtx
func GetStaffIdFromCtx(ctx context.Context) string {
	md, _ := metadata.FromOutgoingContext(ctx)

	return md.Get(CtxKeyJwtStaffId)[0]
}
func GetUuidFromCtx(ctx context.Context) string {
	md, _ := metadata.FromOutgoingContext(ctx)

	return md.Get(CtxUserUuid)[0]
}

// GetStoreIdFromCtx 获取门店编号
func GetStoreIdFromCtx(ctx context.Context) string {
	md, _ := metadata.FromOutgoingContext(ctx)

	return md.Get(CtxKeyJwtStoreId)[0]
}

func GetShopIdFromCtx(ctx context.Context) int64 {

	md, _ := metadata.FromOutgoingContext(ctx)

	var claim claimRes
	_ = mapstructure.Decode(md, &claim)
	// 如果没有返回 0
	if len(claim.JwtShopId) == 0 {
		return 0
	}

	parseInt, _ := strconv.ParseInt(claim.JwtShopId[0], 10, 64)
	return parseInt
}

// GetStoreIdFromCtxToInt64 获取门店编号转为int64
func GetStoreIdFromCtxToInt64(ctx context.Context) int64 {
	storeIdStr := strings.TrimSpace(GetStoreIdFromCtx(ctx))
	if storeIdStr == "" {
		return 0
	}
	storeId, err := strconv.ParseInt(storeIdStr, 10, 64)
	if err != nil {
		return 0
	}
	return storeId
}

func GetReferralCodeFromCtx(ctx context.Context) string {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return ""
	}
	var claim claimRes
	_ = mapstructure.Decode(md, &claim)
	return claim.JwtReferralCode[0]
}

// GetUserRoleFromCtx 获取用户角色
func GetUserRoleFromCtx(ctx context.Context) string {
	md, _ := metadata.FromOutgoingContext(ctx)

	return md.Get(CtxUserRole)[0]
}

// GetMsIdFromCtx 获取系统编号
func GetMsIdFromCtx(ctx context.Context) int32 {
	md, _ := metadata.FromOutgoingContext(ctx)

	var claim claimRes
	_ = mapstructure.Decode(md, &claim)
	parseInt, _ := strconv.ParseInt(claim.JwtMsId[0], 10, 64)
	return int32(parseInt)
}

func GetUserIsMaskFromCtx(ctx context.Context) int32 {
	md, _ := metadata.FromOutgoingContext(ctx)

	var claim claimRes
	_ = mapstructure.Decode(md, &claim)
	parseInt, _ := strconv.ParseInt(claim.JwtUserIsMask[0], 10, 64)
	return int32(parseInt)
}

// GetUserTypeFromCtx 获取用户类别
func GetUserTypeFromCtx(ctx context.Context) int32 {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return 0
	}
	var claim claimRes
	_ = mapstructure.Decode(md, &claim)

	parseInt, _ := strconv.ParseInt(claim.JwtUserType[0], 10, 64)
	return int32(parseInt)
}

// GetIsSupperAdminFromCtx 获取用户是否为超级管理员
func GetIsSupperAdminFromCtx(ctx context.Context) int32 {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return 0
	}
	var claim claimRes
	_ = mapstructure.Decode(md, &claim)

	parseInt, _ := strconv.ParseInt(claim.JwtIsSupperAdmin[0], 10, 64)
	return int32(parseInt)
}

func UpdateMetadata(ctx context.Context, data map[string]string) context.Context {
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
