package tool

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"im/common/uniqueid"
)

// CreateQrCodeSavePath 生成一个二维码并保存到指定的路径
func CreateQrCodeSavePath(url string, fileData FileData) (string, error) {
	pngData, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return "", fmt.Errorf("二维码生成失败 : %v", err)
	}
	// 设置数据
	fileData.Data = pngData
	fileData.Prefix = uniqueid.SN_PREFIX_SKU_QRCODE
	fileData.FileType = "png"

	// 保存文件到本地
	return SaveFileToLocal(pngData, fileData)
}

func CreateShopQrCodeSavePath(url string, fileData FileData) (string, error) {
	pngData, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return "", fmt.Errorf("二维码生成失败 : %v", err)
	}
	// 设置数据
	fileData.Data = pngData
	fileData.Prefix = uniqueid.SN_PREFIX_SHOP_QRCODE
	fileData.FileType = "png"

	// 保存文件到本地
	return SaveFileToLocal(pngData, fileData)
}
