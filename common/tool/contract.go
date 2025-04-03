package tool

import (
	"fmt"
	"im/common/uniqueid"
)

// DownloadAndSaveContractFile 下载并保存pdf合同文件
func DownloadAndSaveContractFile(url string, fileData FileData) (string, error) {
	// 下载文件
	pdf, err := DownloadFile(url)
	if err != nil {
		return "", fmt.Errorf("下载并保存pdf合同文件 - 下载文件失败 : %v", err)
	}
	// 保存文件
	return SaveContractFile(pdf, fileData)
}

// SaveContractFile 保存pdf合同文件
func SaveContractFile(pdf []byte, fileData FileData) (string, error) {
	// 设置数据
	fileData.Data = pdf
	fileData.Prefix = uniqueid.SnPrefixAntChainSign
	fileData.FileType = "pdf"

	// 保存文件到本地
	return SaveFileToLocal(pdf, fileData)
}
