package tool

import (
	"fmt"
	"strings"
	"time"
)

const (
	// 日期格式
	DATE_FORMAT = "2006-01-02"
	// 日期时间格式
	DATETIME_FORMAT = "2006-01-02 15:04:05"
	DateNumFormat   = "20060102150405" // 日期数字格式

	TimeFormat = "15:04:05" // 时间格式
)

// TimestampToDateNum 时间戳转日期数字格式
func TimestampToDateNum(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	t := time.Unix(timestamp, 0)
	return t.Format(DateNumFormat)
}

func Date2Timestamp(date string) int64 {
	if date == "" {
		return 0
	}
	t, _ := time.ParseInLocation(DATE_FORMAT, date, time.Local)
	return t.Unix()
}

func DateTime2Timestamp(date string) int64 {
	if date == "" {
		return 0
	}
	t, _ := time.ParseInLocation(DATETIME_FORMAT, date, time.Local)
	return t.Unix()
}

func DateToTimestamp(date string) int64 {
	if date == "" {
		return 0
	}
	t, _ := time.ParseInLocation(DATE_FORMAT, date, time.Local)
	return t.Unix() * 1000
}

func DateTimeToTimestamp(date string) int64 {
	if date == "" {
		return 0
	}
	t, _ := time.ParseInLocation(DATETIME_FORMAT, date, time.Local)
	return t.Unix() * 1000
}

func Timestamp2Date(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	t := time.Unix(timestamp, 0)
	return t.Format(DATETIME_FORMAT)
}

// Timestamp2Day 时间戳转日期
func Timestamp2Day(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	t := time.Unix(timestamp, 0)
	return t.Format(DATE_FORMAT)
}

// 毫秒的时间戳  转 日期
func TimestampToDate(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	t := time.Unix(timestamp/1000, 0)
	return t.Format(DATETIME_FORMAT)
}

// TimestampToDay 毫秒的时间戳  转 日期
func TimestampToDay(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	t := time.Unix(timestamp/1000, 0)
	return t.Format(DATE_FORMAT)
}

func GetCurrentDay() string {
	return time.Now().Format(DATE_FORMAT)
}

// 获取当前的时间 - 字符串
func GetCurrentDate() string {
	return time.Now().Format(DATETIME_FORMAT)
}

// 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}

func GetMaxAndMinDate(num int) (min, max time.Time) {
	now := time.Now()
	days := now.AddDate(0, 0, num)
	// 计算三天前的零点时间
	min = time.Date(days.Year(), days.Month(), days.Day(), 0, 0, 0, 0, days.Location())
	// 计算三天前的最大时间戳（23:59:59）
	max = time.Date(days.Year(), days.Month(), days.Day(), 23, 59, 59, 999999999, days.Location())
	return
}

// AddDaysToCurrentTime 返回当前时间加上指定天数后的时间戳（Unix 时间戳，单位为秒）
func AddDaysToCurrentTime(daysToAdd int) int64 {
	// 获取当前时间
	currentTime := time.Now()

	// 计算未来时间
	futureTime := currentTime.AddDate(0, 0, daysToAdd)

	// 将未来时间转换为 Unix 时间戳（秒）
	futureTimestamp := futureTime.Unix()

	return futureTimestamp
}

func AddMonthToCurrentTime(num int) int64 {
	// 获取当前时间
	currentTime := time.Now()

	// 计算未来时间
	futureTime := currentTime.AddDate(0, num, 0)

	// 将未来时间转换为 Unix 时间戳（秒）
	futureTimestamp := futureTime.Unix()

	return futureTimestamp
}

// DaysDifferenceByDate 计算时间戳与当前时间的天数差(按日期算 - 隔1天算1天)
func DaysDifferenceByDate(timestamp int64) int {
	givenTime := time.Unix(timestamp, 0)
	currentTime := time.Now()

	// 提取日期部分
	givenDate := time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), 0, 0, 0, 0, time.Local)
	currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)

	// 计算日期差
	days := int(currentDate.Sub(givenDate).Hours() / 24)
	return days
}

// DaysDifference 【天数为正数绝对值】计算时间戳与当前时间的天数差(按时间戳算 - 满24小时算1天)
func DaysDifference(givenTimestamp int64) int {
	// 获取当前时间戳
	currentTime := time.Now().Unix()

	// 将当前时间戳转换为天数
	currentDay := currentTime / 86400 // 一天的秒数为 86400

	// 将给定时间戳转换为天数
	givenDay := givenTimestamp / 86400

	// 计算天数差（绝对值）
	daysDifference := currentDay - givenDay
	if daysDifference < 0 {
		daysDifference = -daysDifference
	}

	return int(daysDifference)
}

// CheckIfOverdue 检查是否逾期并返回逾期天数【基于当前时间0点】
func CheckIfOverdue(plannedTime int64) (bool, int) {
	// 将计划扣款时间戳转换为 `time.Time` 对象
	plannedDate := time.Unix(plannedTime, 0)

	// 计算计划日期的次日零点
	nextDay := time.Date(plannedDate.Year(), plannedDate.Month(), plannedDate.Day()+1, 0, 0, 0, 0, plannedDate.Location())

	// 获取当前时间并计算当前零点时间
	now := time.Now()
	nowDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 如果当前时间小于次日零点，表示未逾期
	if nowDate.Before(nextDay) {
		return false, 0
	}

	// 计算逾期天数
	overdueDays := int(nowDate.Sub(nextDay).Hours()/24) + 1
	return true, overdueDays
}

// DaysDifferenceHasNegative 【天数有正数和负数】计算时间戳与当前时间的天数差(按时间戳算 - 满24小时算1天)
func DaysDifferenceHasNegative(givenTimestamp int64) int {
	// 获取当前时间戳
	currentTime := time.Now().Unix()

	// 将当前时间戳转换为天数
	currentDay := currentTime / 86400 // 一天的秒数为 86400

	// 将给定时间戳转换为天数
	givenDay := givenTimestamp / 86400

	// 计算天数差（绝对值）
	daysDifference := currentDay - givenDay
	return int(daysDifference)
}

// GetTimeAfterDays 传入天数，返回天数后的时间
func GetTimeAfterDays(days int, calTime time.Time) time.Time {
	futureTime := calTime.AddDate(0, 0, days)
	return futureTime
}

// GetTimeAfterWeeks 传入周数，返回周数后的时间
func GetTimeAfterWeeks(weeks int, calTime time.Time) time.Time {
	futureTime := calTime.AddDate(0, 0, weeks*7)
	return futureTime
}

// GetTimeAfterMonths 传入月数，返回一个月后的时间
func GetTimeAfterMonths(months int, calTime time.Time) time.Time {
	futureTime := calTime.AddDate(0, months, 0)
	return futureTime
}

// GetTodayTimestamp 获取今天开始和结束的时间戳
func GetTodayTimestamp() (int64, int64) {
	now := time.Now()
	// 获取今天的日期
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return startOfDay.Unix(), endOfDay.Unix()
}

// GetStartOfDayTimestamp 获取今天开始的时间戳（00:00:00）
func GetStartOfDayTimestamp() int64 {
	now := time.Now()
	// 获取今天的日期
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return startOfDay.Unix()
}

// GetEndOfDayTimestamp 获取今天结束的时间戳（23:59:59）
func GetEndOfDayTimestamp() int64 {
	now := time.Now()
	// 获取今天的日期
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return endOfDay.Unix()
}

// GetStartOfTimeTimestamp 获取 Time 开始的时间戳（00:00:00）
func GetStartOfTimeTimestamp(t time.Time) int64 {
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return startOfDay.Unix()
}

// GetEndOfTimeTimestamp 获取 Time 结束的时间戳（23:59:59）
func GetEndOfTimeTimestamp(t time.Time) int64 {
	endOfDay := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return endOfDay.Unix()
}

// GetTimeTimestamp 获取 Time 开始和结束的时间戳
func GetTimeTimestamp(t time.Time) (int64, int64) {
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endOfDay := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return startOfDay.Unix(), endOfDay.Unix()
}

// Timestamp2Iso8601Datetime 时间戳转ISO 8601 标准日期
func Timestamp2Iso8601Datetime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

// Iso8601Datetime2timestamp ISO 8601 标准日期转时间戳
func Iso8601Datetime2timestamp(dateStr string) int64 {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" {
		return 0
	}
	// 解析日期字符串
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return 0
	}
	// 将时间对象转换为 Unix 时间戳（秒）
	unixTimestamp := t.Unix()
	return unixTimestamp
}

var (
	uptime = time.Now()
)

// AddDate returns the time corresponding to adding the given number of years, months, and days to t. For example, AddDate(-1, 2, 3) applied to January 1, 2011 returns March 4, 2010.
//
// However, this version of AddDate will examine the date of month. For example, AddDate(0, 1, 0) applied to January 31, 2011 returns Feburary 28, 2011 instead of March 3, 2011.
func AddDate(t time.Time, years, months, days int) time.Time {
	// limit month
	if months >= 12 || months <= 12 {
		years += months / 12
		months = months % 12
	}

	// get datetime parts
	ye := t.Year()
	mo := t.Month()
	da := t.Day()
	hr := t.Hour()
	mi := t.Minute()
	se := t.Second()
	ns := t.Nanosecond()
	lo := t.Location()

	// years
	ye += years

	// months
	mo += time.Month(months)
	if mo > 12 {
		mo -= 12
		ye++
	} else if mo < 1 {
		mo += 12
		ye--
	}

	// after adding month, we should adjust day of month value
	if da <= 28 {
		// nothing to change
	} else if da == 29 {
		if mo == 2 {
			if !isLeapYear(ye) {
				da = 28
			}
		}
	} else if da == 30 {
		if mo == 2 {
			if isLeapYear(ye) {
				da = 29
			} else {
				da = 28
			}
		}
	} else if da == 31 {
		switch mo {
		case 2:
			if isLeapYear(ye) {
				da = 29
			} else {
				da = 28
			}
		case 1, 3, 5, 7, 8, 10, 12:
			da = 31
		case 4, 6, 9, 11:
			da = 30
		}
	}
	da += days

	return time.Date(ye, mo, da, hr, mi, se, ns, lo)
}

// AddDate 加日期，支持跨年、闰年以及月份和天数调整
//func AddDate2(t time.Time, years, months, days int) time.Time {
//	// 1. 先处理月份和年份调整
//	// 计算总的年份和月份
//	totalMonths := int(t.Month()) - 1 + months // 月份从 0 开始计算
//	years += totalMonths / 12                  // 超过 12 个月的部分加到年份
//	months = totalMonths % 12                  // 剩余的月份
//	if months < 0 {                            // 如果月份是负数，调整为上一年
//		years--
//		months += 12
//	}
//
//	// 2. 调整年份
//	newYear := t.Year() + years
//	newMonth := time.Month(months + 1) // 转换回 Go 的月份（从 1 开始）
//
//	// 3. 处理天数调整，确保日期有效
//	newTime := time.Date(newYear, newMonth, t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
//
//	// 如果初始日期是 31 号，但目标月份没有 31 号，则会自动调整为该月的最后一天
//	if newTime.Day() < t.Day() {
//		newTime = newTime.AddDate(0, 0, -newTime.Day()+t.Day())
//	}
//
//	// 4. 最后加上天数
//	newTime = newTime.AddDate(0, 0, days)
//
//	return newTime
//}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

// TimeDifference 时间差
type TimeDifference struct {
	Years   int
	Months  int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

// CalculateTimeDiffString 计算2个time的时间差返回字符串, 若t1大于t2，返回空字符串
// format = "2006-01-02 15:04:05"
func CalculateTimeDiffString(t1, t2 time.Time, format string) (bool, string) {
	isAfter, timeDifference := CalculateTimeDifference(t1, t2)
	if isAfter {
		return isAfter, ""
	}
	// 替换格式
	format = strings.ReplaceAll(format, "2006", "year")
	format = strings.ReplaceAll(format, "01", "month")
	format = strings.ReplaceAll(format, "02", "day")
	format = strings.ReplaceAll(format, "15", "hour")
	format = strings.ReplaceAll(format, "04", "minute")
	format = strings.ReplaceAll(format, "05", "second")
	// 替换时间
	format = strings.ReplaceAll(format, "year", fmt.Sprintf("%04d", timeDifference.Years))
	format = strings.ReplaceAll(format, "month", fmt.Sprintf("%02d", timeDifference.Months))
	format = strings.ReplaceAll(format, "day", fmt.Sprintf("%02d", timeDifference.Days))
	format = strings.ReplaceAll(format, "hour", fmt.Sprintf("%02d", timeDifference.Hours))
	format = strings.ReplaceAll(format, "minute", fmt.Sprintf("%02d", timeDifference.Minutes))
	format = strings.ReplaceAll(format, "second", fmt.Sprintf("%02d", timeDifference.Seconds))
	return isAfter, format
}

// CalculateTimeDifference 计算2个time的时间差, 若t1大于t2，返回nil
func CalculateTimeDifference(t1, t2 time.Time) (isAfter bool, timeDifference *TimeDifference) {
	// 判断2个那个大
	isAfter = t1.After(t2)
	if isAfter {
		return
	}

	// 计算天数差
	days := int(t2.Sub(t1).Hours() / 24)

	// 计算年数
	years := t2.Year() - t1.Year()

	// 处理月份差异
	months := int(t2.Month()) - int(t1.Month())
	if months < 0 {
		months += 12
		years--
	}

	// 计算剩余天数
	// 得到同一个年份和月份下的两个日期天数差
	daysInMonth := daysInMonth(t1.Year(), t1.Month())
	daysLeft := days - months*daysInMonth

	// 计算小时、分钟、秒差
	hours := int(t2.Sub(t1).Hours())
	minutes := int(t2.Sub(t1).Minutes()) % 60
	seconds := int(t2.Sub(t1).Seconds()) % 60

	// 设置返回值
	timeDifference = &TimeDifference{
		Years:   years,
		Months:  months,
		Days:    daysLeft,
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}
	return
}

// 获取月份天数
func daysInMonth(year int, month time.Month) int {
	switch month {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return 29 // 闰年
		}
		return 28
	}
	return 0
}
