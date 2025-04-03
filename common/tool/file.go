package tool

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"im/common/uniqueid"
)

const (
	ProjectSku                       = "api"             // 项目 - api
	ModulesOrderExcel                = "order_excel"     // 项目 - 订单excel
	ModulesFinancialExcel            = "financial_excel" // 项目 - 收支明细excel
	ModulesOverdueExcel              = "overdue_excel"   // 项目 - 逾期明细excel
	ModulesPlanListExcel             = "plan_list_excel" // 项目 - 扣款计划
	ModuleSku                        = "sku"
	ModuleSkuCustom                  = "sku_custom"
	ModuleStoreShop                  = "store_shop"                   // 模块 - 商品
	ModulesHeePayOrderExcel          = "hee_pay_order_excel"          // 项目 - 汇付宝订单excel
	ModulesHeePayFinancialExcel      = "hee_pay_financial_excel"      // 项目 - 汇付宝收支明细excel
	ModulesHeePayOverdueExcel        = "hee_pay_overdue_excel"        // 项目 - 汇付宝逾期明细excel
	ModulesAntSignContract           = "ant_sign_contract"            // 项目 - 区块链签约合同
	ModulesAntOrderExcel             = "ant_order_excel"              // 项目 - 蚂蚁区块链订单excel
	ModulesAntFinancialExcel         = "ant_financial_excel"          // 项目 - 蚂蚁区块链收支明细excel
	ModulesAntOverdueExcel           = "ant_overdue_excel"            // 项目 - 蚂蚁区块链逾期明细excel
	ModulesAntPlanListExcel          = "ant_plan_list_excel"          // 项目 - 蚂蚁区块链扣款计划
	ModulesAntReceivableExcel        = "ant_receivable_excel"         // 项目 - 蚂蚁区块链待收款excel
	ModulesCommerceReceivableExcel   = "commerce_receivable_excel"    // 项目 - 安心付待收款excel
	ModulesCommerceComplainListExcel = "commerce_complain_list_excel" // 项目 - 安心付投诉列表
	ModulesPlatformComplainListExcel = "platform_complain_list_excel" // 项目 - 平台投诉列表
	ModulesAiSignContract            = "ai_sign_contract"             // 项目 - 爱签签约合同

	FileTypeXlsx = "xlsx" // excel类型
	FileTypePdf  = "pdf"  // pdf类型
)

// FileData 文件数据
type FileData struct {
	Data     []byte            // 文件数据
	Domain   string            // 域名
	Path     string            // 路径
	Prefix   uniqueid.SnPrefix // 前缀
	Project  string            // 项目
	Module   string            // 模块
	FileType string            // 文件类型 -> png
	FileName string            // 指定存储文件名 - 不指定则随机生成
}

// 设置本地文件路径
func setLocalFilePath(fileData FileData) (relativePath, fileDir, fileName string) {
	// 文件路径 : 路径/statics/项目/模块
	relativePath = fmt.Sprintf("/statics/%s/%s", fileData.Project, fileData.Module)
	fileDir = filepath.Join(fileData.Path, relativePath)
	// 判断是否指定文件名
	setFileName := strings.TrimSpace(fileData.FileName)
	if setFileName != "" {
		fileName = setFileName + "." + fileData.FileType
	} else {
		fileName = uniqueid.GenSn(fileData.Prefix) + "." + fileData.FileType
	}
	return
}

// SaveFileToLocal 保存文件到本地
func SaveFileToLocal(data []byte, fileData FileData) (string, error) {
	// 设置本地文件路径
	relativePath, fileDir, fileName := setLocalFilePath(fileData)

	// 如果目录不存在就创建
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err = os.MkdirAll(fileDir, 0755)
		if err != nil {
			return "", fmt.Errorf("文件夹创建失败 : %v", err)
		}
	}

	// 保存文件
	filePath := filepath.Join(fileDir, fileName)
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", fmt.Errorf("文件创建失败 : %v", err)
	}

	// 返回文件路径拼接的url地址
	fileRelativePath := filepath.Join(relativePath, fileName)
	url := fmt.Sprintf("%s%s", fileData.Domain, fileRelativePath)
	return url, nil
}

// ReadLocalFileUrl 读取本地文件的下载地址
func ReadLocalFileUrl(fileData FileData) (exist bool, url string, err error) {
	// 设置本地文件路径
	relativePath, fileDir, fileName := setLocalFilePath(fileData)
	filePath := filepath.Join(fileDir, fileName)
	// 判断文件是否存在
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, "", nil
	}
	if err != nil {
		return false, "", fmt.Errorf("文件探测状态失败 : %v", err)
	}
	// 返回文件路径拼接的url地址
	fileRelativePath := filepath.Join(relativePath, fileName)
	url = fmt.Sprintf("%s%s", fileData.Domain, fileRelativePath)
	return true, url, nil
}

// ReadLocalFile 读取本地保存的文件
func ReadLocalFile(fileData FileData) (exist bool, data []byte, err error) {
	// 设置本地文件路径
	_, fileDir, fileName := setLocalFilePath(fileData)
	filePath := filepath.Join(fileDir, fileName)
	// 判断文件是否存在
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil, nil
	}
	if err != nil {
		return false, nil, fmt.Errorf("文件探测状态失败 : %v", err)
	}
	// 读取文件
	data, err = os.ReadFile(filePath)
	if err != nil {
		return false, nil, fmt.Errorf("文件读取失败 : %v", err)
	}

	return true, data, nil
}

// ReadFile 读取文件
func ReadFile(filepath string) ([]byte, error) {
	// 检查文件是否存在
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("文件不存在: %s", filepath)
	}

	// 打开文件
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func GetFileInfo(filePath string) (fileName, fileExt string, reader io.Reader, err error) {
	// 解析 URL 以获取文件名
	parsedURL, err := url.Parse(filePath)
	if err != nil {
		return "", "", nil, fmt.Errorf("error parsing URL: %w", err)
	}

	fileName = path.Base(parsedURL.Path)
	fileExt = strings.TrimPrefix(path.Ext(fileName), ".")
	// 发起 HTTP 请求以获取文件
	resp, err := http.Get(filePath)
	if err != nil {
		return "", "", nil, fmt.Errorf("error fetching image: %w", err)
	}

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return "", "", nil, fmt.Errorf("bad status: %s", resp.Status)
	}
	reader = resp.Body

	//contentType = resp.Header.Get("Content-Type")
	return
}

// DownloadFile 下载文件
func DownloadFile(url string) ([]byte, error) {
	url = strings.TrimSpace(url)
	// 判空
	if url == "" {
		return nil, fmt.Errorf("无效文件的地址")
	}
	// 发送 HTTP GET 请求获取文件
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("无法下载：%v", err)
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("下载失败，状态码：%d", resp.StatusCode)
	}
	// 读取文件内容到字节切片
	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("下载失败，状态码：%d", resp.StatusCode)
	}

	// 保存文件到本地
	return fileBytes, nil
}
