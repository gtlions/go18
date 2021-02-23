package gos10i

import (
	"fmt"
	"strings"
	"time"
)

// XMonthDayList 当前月份日期列表
//
// when 时间
//
func XMonthDayList(when time.Time) (dayList []string) {
	for i := 0; i < int((XMonthLast(when).Sub(XMonthFirst(when)).Hours()/24)+1); i++ {
		dayList = append(dayList, XMonthFirst(when).AddDate(0, 0, i).Format("2006-01-02"))
	}
	return

}

// XMonthDayCnt 当前月份天数
//
// when 时间
//
func XMonthDayCnt(when time.Time) int {
	return int((XMonthLast(when).Sub(XMonthFirst(when)).Hours() / 24) + 1)
}

// XOrderNoFromNow 基于当前时间的字符串，可作为订单号使用
//
// 格式 200601021504050000
//
func XOrderNoFromNow() string {
	ct := time.Now().Format("20060102150405.0000")
	ct = strings.Replace(ct, ".", "", -1)
	return ct
}

// XChinaMonth 返回日期的『n月』
//
// day 日期
//
func XChinaMonth(day time.Time) string {
	switch day.Month() {
	case time.January:
		return fmt.Sprintf("%s月", "一")
	case time.February:
		return fmt.Sprintf("%s月", "二")
	case time.March:
		return fmt.Sprintf("%s月", "三")
	case time.April:
		return fmt.Sprintf("%s月", "四")
	case time.May:
		return fmt.Sprintf("%s月", "五")
	case time.June:
		return fmt.Sprintf("%s月", "六")
	case time.July:
		return fmt.Sprintf("%s月", "七")
	case time.August:
		return fmt.Sprintf("%s月", "八")
	case time.September:
		return fmt.Sprintf("%s月", "九")
	case time.October:
		return fmt.Sprintf("%s月", "十")
	case time.November:
		return fmt.Sprintf("%s月", "十一")
	case time.December:
		return fmt.Sprintf("%s月", "十二")
	default:
		return "错误时间格式"
	}
}

// XChinaWeekday 返回日期的『周n』或『星期n』
//
// 每周的第一天为周一
//
// prefix 前缀,周|星期
//
// day 日期
//
func XChinaWeekday(prefix string, day time.Time) string {
	switch day.Weekday() {
	case time.Sunday:
		return fmt.Sprintf("%s%s", prefix, "日")
	case time.Monday:
		return fmt.Sprintf("%s%s", prefix, "一")
	case time.Tuesday:
		return fmt.Sprintf("%s%s", prefix, "二")
	case time.Wednesday:
		return fmt.Sprintf("%s%s", prefix, "三")
	case time.Thursday:
		return fmt.Sprintf("%s%s", prefix, "四")
	case time.Friday:
		return fmt.Sprintf("%s%s", prefix, "五")
	case time.Saturday:
		return fmt.Sprintf("%s%s", prefix, "六")
	default:
		return "错误时间格式"
	}
}

// XWeekdayInt 返回日期的本周的第『n』天,1-7
//
// 每周的第一天为周一
//
// t 时间
//
func XWeekdayInt(t time.Time) int {
	switch t.Weekday() {
	case time.Sunday:
		return 7
	case time.Monday:
		return 1
	case time.Tuesday:
		return 2
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 5
	case time.Saturday:
		return 6
	default:
		return 0
	}
}

// XDayLast235959 返回天的最后时间23:59:59.999999999
//
// t 时间
//
func XDayLast235959(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// XWeekFirst 返回周的第一天的时间
//
// t 时间
//
func XWeekFirst(t time.Time) time.Time {
	return XZeroTime(t.AddDate(0, 0, -(XWeekdayInt(t) - 1)))
}

// XWeekLast 返回周的最后一天的时间
//
// t 时间
//
func XWeekLast(t time.Time) time.Time {
	return XWeekFirst(t).AddDate(0, 0, 6)
}

// XWeekLast235959 返回周的最后一天的时间23:59:59.999999999
//
// t 时间
//
func XWeekLast235959(t time.Time) time.Time {
	d := XWeekFirst(t).AddDate(0, 0, 6)
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 999999999, d.Location())
}

// XMonthFirst 返回月的第一天的时间
//
// t 时间
//
func XMonthFirst(t time.Time) time.Time {
	return XZeroTime(t.AddDate(0, 0, -t.Day()+1))
}

// XMonthLast 返回月的最后一天的时间
//
// t 时间
//
func XMonthLast(t time.Time) time.Time {
	return XMonthFirst(t).AddDate(0, 1, -1)
}

// XMonthLast235959 返回月的最后一天的时间23:59:59.999999999
//
// t 时间
//
func XMonthLast235959(t time.Time) time.Time {
	d := XMonthFirst(t).AddDate(0, 1, -1)
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 999999999, d.Location())
}

// XYearFirst 返回年的第一天的时间
//
// t 时间
//
func XYearFirst(t time.Time) time.Time {
	t1, _ := time.Parse("2006-01-02", t.Format("2006")+"-01-01")
	return XZeroTime(t1)
}

// XYearLast 返回年的最后一天的时间
//
// t 时间
//
func XYearLast(t time.Time) time.Time {
	return XYearFirst(t).AddDate(1, 0, -1)
}

// XYearLast235959 返回年的最后一天的时间23:59:59.999999999
//
// t 时间
//
func XYearLast235959(t time.Time) time.Time {
	d := XYearFirst(t).AddDate(1, 0, -1)
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 999999999, d.Location())
}

// XZeroTime 0点时间
//
// t 时间
//
func XZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
