package tool

import (
	"net"
	"net/http"
	"strings"
)

func GetClientIP(r *http.Request) string {
	// 从X-Forwarded-For头部获取客户端IP地址
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		ips := strings.Split(forwarded, ",")
		for i := len(ips) - 1; i >= 0; i-- {
			// 去除空格
			ip := strings.TrimSpace(ips[i])
			if ip != "" {
				return ip
			}
		}
	}

	// 如果X-Forwarded-For头部不存在，则直接从RemoteAddr中获取
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	return ip
}

// GetClientIpV2 获取客户端的真实IP地址，不包含端口号
func GetClientIpV2(r *http.Request) string {
	// 获取 X-Forwarded-For 头
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// X-Forwarded-For 可能包含多个IP地址，取第一个
		ips := strings.Split(xForwardedFor, ",")
		if len(ips) > 0 {
			ip := strings.TrimSpace(ips[0])
			return stripPort(ip)
		}
	}

	// 获取 X-Real-IP 头
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return stripPort(xRealIP)
	}

	// 返回 RemoteAddr
	return stripPort(r.RemoteAddr)
}

// stripPort 移除IP地址中的端口号
func stripPort(address string) string {
	ip, _, err := net.SplitHostPort(address)
	if err != nil {
		// 如果 SplitHostPort 解析失败，说明地址中没有端口号，直接返回原始地址
		return address
	}
	return ip
}
