package tool

import (
	"net/url"
	"path"
	"regexp"
	"strings"
)

// JoinURL 拼接域名和路径
func JoinURL(domain, urlPath string) string {
	domain = strings.TrimSpace(domain)
	urlPath = strings.TrimSpace(urlPath)
	// 移除 domain 末尾的 /
	domain = strings.TrimRight(domain, "/")

	// 确保 urlPath 以 / 开头
	if !strings.HasPrefix(urlPath, "/") {
		urlPath = "/" + urlPath
	}

	// 拼接路径
	fullURL := domain + urlPath

	return fullURL
}

// IsValidURL 检查字符串是否是以 "http" 或 "https" 开头的有效域名链接地址
func IsValidURL(url string) bool {
	// 正则表达式，用于匹配以 http 或 https 开头的有效域名链接地址
	var re = regexp.MustCompile(`^https?://(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,6}(?:/[^/#?]*)?(?:\?[^#]*)?(?:#.*)?$`)
	return re.MatchString(url)
}

// GetFileNameByURL 从给定的 URL 中提取文件名
func GetFileNameByURL(link string) string {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return ""
	}
	return path.Base(parsedURL.Path)
}
