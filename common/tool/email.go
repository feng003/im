package tool

import "regexp"

// IsEmail 校验字符串是否是邮箱，返回bool
func IsEmail(email string) bool {
	// 正则表达式用于校验邮箱
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
