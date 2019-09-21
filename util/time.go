package util

import (
	"fmt"
	"time"
)

// GetTime 获取当前时间戳
func GetTime() int64 {
	return time.Now().Unix()
}

// GetMicrotime 获取微秒时间
func GetMicrotime() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// GetMillsTime 获取当前微秒时间
// 名称拼写错误了，不知道以前有哪些地方调用，以后的项目中废弃
func GetMillsTime() int64 {
	return time.Now().UnixNano() / 1000000
}

// GetMilliTime 获取当前毫秒时间
func GetMilliTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetTimestamp 获取当前格式化时间
func GetTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FormatUnixTime 将时间戳格式化
func FormatUnixTime(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("2006-01-02 15:04:05")
}

// GetYMD 获取当前年月日
func GetYMD() string {
	return time.Now().Format("20060102")
}

// GetChinaWeekDay 获取中国的星期几
func GetChinaWeekDay() int {
	weekDay := time.Now().Weekday()
	if weekDay == time.Sunday {
		return 7
	}
	return int(weekDay)
}

// GetTodayStartTime 获取当天开始时间
func GetTodayStartTime() int64 {
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02", timeStr)
	return t.Unix()
}

// GetChinaWeekStartTime 获取中国的本周一开始时间
func GetChinaWeekStartTime() int64 {
	// 获取今天是周几
	weekday := GetChinaWeekDay()
	// 获取当天开始
	startTime := GetTodayStartTime()

	return startTime - int64((weekday-1)*86400)
}

// MonthDayNum t 所在时间的月份总天数
func MonthDayNum(t time.Time) int {
	isLeapYear := isLeap(t.Year())

	month := t.Month()
	switch month {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.February:
		if isLeapYear {
			return 29
		}

		return 28
	default:
		return 30
	}
}

func TimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)
	hours := diff.Hours()
	if hours < 1.0 {
		return fmt.Sprintf("约 %.0f 分钟前", diff.Minutes())
	}

	if hours < 24.0 {
		return fmt.Sprintf("约 %.0f 小时前", hours)
	}

	if hours < 72.0 {
		return fmt.Sprintf("约 %.0f 天前", hours/24.0)
	}

	// 同一年，不用年份
	if now.Year() == t.Year() {
		return t.Format("01-02 15:04")
	}

	return t.Format("2006-01-02")
}

// 是否闰年
func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
