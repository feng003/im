package ctxdata

import (
	"context"
	"strconv"

	"github.com/goinggo/mapstructure"
	"google.golang.org/grpc/metadata"
)

var (
	CmsJwtUserId   = "cmsJwtUserId"
	CmsJwtStoreId  = "cmsJwtStoreId"
	CmsJwtMsId     = "cmsJwtMsId"
	CmsJwtPhone    = "cmsJwtPhone"
	CmsJwtUserType = "cmsJwtUserType"
)

type claimCms struct {
	CmsJwtUserId   []string `json:"cmsJwtUserId"`
	CmsJwtStoreId  []string `json:"cmsJwtStoreId"`
	CmsJwtMsId     []string `json:"cmsJwtMsId"`
	CmsJwtPhone    []string `json:"cmsJwtPhone"`
	CmsJwtUserType []string `json:"cmsJwtUserType"`
}

func GetCmsUidFromCtx(ctx context.Context) int64 {
	md, _ := metadata.FromOutgoingContext(ctx)
	var claim claimCms
	_ = mapstructure.Decode(md, &claim)

	parseInt, _ := strconv.ParseInt(claim.CmsJwtUserId[0], 10, 64)

	return parseInt
}

func GetCmsMsIdFromCtx(ctx context.Context) int64 {
	md, _ := metadata.FromOutgoingContext(ctx)
	var claim claimCms
	_ = mapstructure.Decode(md, &claim)

	parseInt, _ := strconv.ParseInt(claim.CmsJwtMsId[0], 10, 64)

	return parseInt
}
