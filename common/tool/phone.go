package tool

import "regexp"

// IsValidPhoneNumber 检测11位手机号码格式是否正确且有效
func IsValidPhoneNumber(phone string) bool {
	pattern := regexp.MustCompile(`^1\d{10}$`)
	return pattern.MatchString(phone)
}
