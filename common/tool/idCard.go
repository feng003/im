package tool

import (
	"strconv"
	"strings"
	"time"
	"unicode"
)

// NewIdCardProcessor 实例化封装身份证处理逻辑
func NewIdCardProcessor() *IdCardProcessor {
	return &IdCardProcessor{}
}

// IDCardInfo 包含从身份证或居住证中提取的信息
type IDCardInfo struct {
	BirthDate time.Time // 出生日期
	Age       int       // 年龄
	Gender    string    // 性别：男或女
	Type      string    // 证件类型：大陆身份证、港澳居住证、台湾居住证
}

// IdCardProcessor 封装身份证和居住证的处理逻辑
type IdCardProcessor struct{}

// Validate 校验身份证或居住证号码并提取信息
func (p *IdCardProcessor) Validate(id string) (bool, *IDCardInfo) {
	id = strings.ToUpper(id)
	if len(id) != 18 {
		return false, nil
	}

	info := &IDCardInfo{}
	info.Type = p.detectType(id)

	// 校验格式
	if !p.isValidFormat(id) {
		return false, nil
	}

	// 校验校验码
	if !p.verifyChecksum(id) {
		return false, nil
	}

	// 提取生日
	birthDate, err := p.parseBirthDate(id[6:14])
	if err != nil {
		return false, nil
	}
	info.BirthDate = birthDate

	// 计算年龄
	info.Age = p.calculateAge(birthDate)

	// 提取性别
	info.Gender = p.parseGender(id[16])

	return true, info
}

// 判断证件类型：大陆身份证、港澳居住证、台湾居住证
func (p *IdCardProcessor) detectType(id string) string {
	prefix := id[:2]
	switch prefix {
	case "81": // 香港
		return "港澳居民居住证"
	case "82": // 澳门
		return "港澳居民居住证"
	case "83": // 台湾
		return "台湾居民居住证"
	default:
		return "大陆身份证"
	}
}

// 判断身份证号码的格式是否正确（全部数字，最后一位可以是X）
func (p *IdCardProcessor) isValidFormat(id string) bool {
	for i, r := range id {
		if i == 17 && r == 'X' {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// 计算校验码
func (p *IdCardProcessor) calculateChecksum(id string) string {
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checksumMap := "10X98765432"
	sum := 0
	for i := 0; i < 17; i++ {
		num, _ := strconv.Atoi(string(id[i]))
		sum += num * weights[i]
	}
	return string(checksumMap[sum%11])
}

// 校验校验码
func (p *IdCardProcessor) verifyChecksum(id string) bool {
	return p.calculateChecksum(id) == string(id[17])
}

// 提取生日
func (p *IdCardProcessor) parseBirthDate(birth string) (time.Time, error) {
	return time.Parse("20060102", birth)
}

// 计算年龄
func (p *IdCardProcessor) calculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}

// 提取性别
func (p *IdCardProcessor) parseGender(genderChar byte) string {
	num, _ := strconv.Atoi(string(genderChar))
	if num%2 == 0 {
		return "女"
	}
	return "男"
}
