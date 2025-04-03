package tool

import (
	"math"
	"regexp"
	"strconv"
)

// AmountConvertChinese 将金额数字转化为中文的大写汉字
func AmountConvertChinese(money float64, pRound bool) string {
	var NumberUpper = []string{"壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖", "零"}
	var Unit = []string{"分", "角", "圆", "拾", "佰", "仟", "万", "拾", "佰", "仟", "亿", "拾", "佰", "仟"}
	var regex = [][]string{
		{"零拾", "零"}, {"零佰", "零"}, {"零仟", "零"}, {"零零零", "零"}, {"零零", "零"},
		{"零角零分", "整"}, {"零分", "整"}, {"零角", "零"}, {"零亿零万零元", "亿元"},
		{"亿零万零元", "亿元"}, {"零亿零万", "亿"}, {"零万零元", "万元"}, {"万零元", "万元"},
		{"零亿", "亿"}, {"零万", "万"}, {"拾零圆", "拾元"}, {"零圆", "元"}, {"零零", "零"}}
	str, DigitUpper, UnitLen, round := "", "", 0, 0
	if money == 0 {
		return "零"
	}
	if money < 0 {
		str = "负"
		money = math.Abs(money)
	}
	if pRound {
		round = 2
	} else {
		round = 1
	}
	Digit_byte := []byte(strconv.FormatFloat(money, 'f', round+1, 64)) //注意币种四舍五入
	UnitLen = len(Digit_byte) - round

	for _, v := range Digit_byte {
		if UnitLen >= 1 && v != 46 {
			s, _ := strconv.ParseInt(string(v), 10, 0)
			if s != 0 {
				DigitUpper = NumberUpper[s-1]

			} else {
				DigitUpper = "零"
			}
			str = str + DigitUpper + Unit[UnitLen-1]
			UnitLen = UnitLen - 1
		}
	}
	for i, _ := range regex {
		reg := regexp.MustCompile(regex[i][0])
		str = reg.ReplaceAllString(str, regex[i][1])
	}
	if string(str[0:3]) == "元" {
		str = string(str[3:len(str)])
	}
	if string(str[0:3]) == "零" {
		str = string(str[3:len(str)])
	}
	return str
}

// CheckAmount 检测金额是否为有效的2位小数
func CheckAmount(amount float64) bool {
	rounded := math.Round(amount * 100)
	return rounded == amount*100
}

// ConvertYuanToFen 把结算金额元转为分
func ConvertYuanToFen(amount float64) int64 {
	return int64(math.Round(amount * 100))
}

// ConvertFenToYuan 把结算金额分转为元
func ConvertFenToYuan(fen int64) float64 {
	return float64(fen) / 100.0
}

// IsValidMoneyFloat 判断是否是小数点2位金额
func IsValidMoneyFloat(num float64) bool {
	// 将数值乘以100，然后取整，再除以100，比较是否与原数相等
	rounded := math.Round(num*100) / 100
	return num == rounded
}

// FloatForMoney 设置金额数值 - 保留2位小数四舍五入
func FloatForMoney(value float64) float64 {
	return RoundToTwoDecimalPlaces(value)
}

// FloatForMoney4 设置金额数值 - 保留4位小数四舍五入
func FloatForMoney4(value float64) float64 {
	return RoundToFourDecimalPlaces(value)
}

// RoundToTwoDecimalPlaces 保留2位小数四舍五入
func RoundToTwoDecimalPlaces(value float64) float64 {
	return math.Round(value*100) / 100
}

// RoundToFourDecimalPlaces 保留4位小数四舍五入
func RoundToFourDecimalPlaces(value float64) float64 {
	return math.Round(value*10000) / 10000
}

// TruncateToTwoDecimalPlaces 截断保留两位小数
func TruncateToTwoDecimalPlaces(num float64) float64 {
	return math.Trunc(num*100) / 100
}
