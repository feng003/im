package tool

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// DownloadImage 下载图片并返回一个 io.Reader 以及关闭函数
func DownloadImage(url string) (io.Reader, func() error, error) {
	// 发送 HTTP GET 请求获取图片
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("无法下载图片：%v", err)
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, nil, fmt.Errorf("下载失败，状态码：%d", resp.StatusCode)
	}

	// 返回响应体和关闭函数
	return resp.Body, resp.Body.Close, nil
}

// DownloadImageWithBytes 下载图片返回bytes
func DownloadImageWithBytes(url string) ([]byte, error) {
	url = strings.TrimSpace(url)
	// 判空
	if url == "" {
		return nil, fmt.Errorf("无效的图片地址")
	}
	// 发送 HTTP GET 请求获取图片
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("无法下载图片：%v", err)
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("下载失败，状态码：%d", resp.StatusCode)
	}
	// 读取图像文件内容到字节切片
	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("下载失败，状态码：%d", resp.StatusCode)
	}

	return imageBytes, nil
}
