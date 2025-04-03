package tool

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// ExcelData excel数据
type ExcelData struct {
	Sheet   string
	Title   []string
	RowData [][]interface{}
}

// SaveExcelToLocalWithSingleSheet 保存excel 支持单个工作表数据到本地并返回url链接地址提供下载
func SaveExcelToLocalWithSingleSheet(excelData ExcelData, fileData FileData) (string, error) {
	var excelDataList []ExcelData
	excelDataList = append(excelDataList, excelData)
	return SaveExcelToLocalWithMultiSheet(excelDataList, fileData)
}

// SaveExcelToLocalWithMultiSheet 保存excel 支持多个工作表数据到本地并返回url链接地址提供下载
func SaveExcelToLocalWithMultiSheet(excelDataList []ExcelData, fileData FileData) (string, error) {
	// 创建一个新的 Excel 文件
	file := excelize.NewFile()

	// 遍历每个工作表的数据
	for _, data := range excelDataList {
		sheetName := data.Sheet
		title := data.Title
		rowData := data.RowData

		// 添加工作表
		index, _ := file.NewSheet(sheetName)

		// 设置标题行
		for colIndex, title := range title {
			cellAddr := string(rune('A'+colIndex)) + "1"
			_ = file.SetCellValue(sheetName, cellAddr, title)
		}

		// 将数据写入工作表
		for rowIndex, row := range rowData {
			for colIndex, cell := range row {
				cellAddr := string(rune('A'+colIndex)) + fmt.Sprint(rowIndex+2)
				_ = file.SetCellValue(sheetName, cellAddr, cell)
			}
		}

		// 设置工作表为默认激活工作表
		if index > 1 {
			file.SetActiveSheet(index)
		}
	}

	// 删除默认的 "Sheet1" 工作表
	err := file.DeleteSheet("Sheet1")
	if err != nil {
		fmt.Println("删除默认的 'Sheet1' 工作表失败:", err)
	}

	// 保存文件为字节数据
	buffer, err := file.WriteToBuffer()
	if err != nil {
		return "", err
	}

	// 保存到本地并返回url链接地址
	return SaveFileToLocal(buffer.Bytes(), fileData)
}
