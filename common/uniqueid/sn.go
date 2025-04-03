package uniqueid

import (
	"fmt"
	"time"
)

// 生成sn单号
type SnPrefix string

const (
	SN_PREFIX_ORDER          SnPrefix = "YS"  //订单前缀
	SN_PREFIX_THIRD_PAYMENT  SnPrefix = "PMT" //第三方支付流水记录前缀
	SN_PREFIX_UPLOAD_PICTURE SnPrefix = "YS"  //图片上传记录前缀
	SN_PREFIX_SHOP_QRCODE    SnPrefix = "SHOP"
	SN_PREFIX_SKU_QRCODE     SnPrefix = "SKU" //sku二维码前缀
	SN_PREFIX_CMS            SnPrefix = "CMS" //cms职位前缀
	SnPrefixBalance          SnPrefix = "BAL" // Balance余额前缀
	SnPrefixAntChainOrder    SnPrefix = "ANT" // 区块链订单前缀
	SnPrefixAntChainRecharge SnPrefix = "ACR" // 区块链预付费充值充值前缀
	SnPrefixAntChainSign     SnPrefix = "ACS" // 区块链用户签约合同前缀
)

// 生成单号
func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20060102150405"), Krand(8, KC_RAND_KIND_NUM))
}

func GenUUid() string {
	return time.Now().Format("20060102150405") + Krand(8, KC_RAND_KIND_NUM)
}

func GenerateStringSlice(start, end int) []string {
	slice := make([]string, 0, end-start+1)
	for i := start; i <= end; i++ {
		slice = append(slice, fmt.Sprintf("%d", i))
	}
	return slice
}
