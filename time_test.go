package gos10i

import (
	"testing"
	"time"
)

func TestXMonthDayList(t *testing.T) {
	when := time.Now()
	t.Log(XMonthDayList(when))
}

func TestXMonthDayCnt(t *testing.T) {
	when := time.Now()
	t.Log(XMonthDayCnt(when))
}
func TestXXNowString(t *testing.T) {
	t.Log(XOrderNoFromNow())
}
func TestXChinaWeekday(t *testing.T) {
	t.Log(XChinaWeekday("星期", time.Now()))
	t.Log(XChinaWeekday("周", time.Now()))
}

func TestXChinaMonth(t *testing.T) {
	t.Log(XChinaMonth(time.Now()))
}

func TestXXWeekdayInt(t *testing.T) {
	t.Log(XWeekdayInt(time.Now().AddDate(0, 0, -5)))
	t.Log(XWeekdayInt(time.Now().AddDate(0, 0, -4)))
	t.Log(XWeekdayInt(time.Now().AddDate(0, 0, -3)))
	t.Log(XWeekdayInt(time.Now().AddDate(0, 0, -2)))
	t.Log(XWeekdayInt(time.Now().AddDate(0, 0, -1)))
	t.Log(XWeekdayInt(time.Now().AddDate(0, 0, 0)))
	t.Log(XWeekdayInt(time.Now().AddDate(0, 0, 1)))
}

func TestXDayLast235959(t *testing.T) {
	t.Log(XDayLast235959(time.Now().AddDate(0, 0, -5)))
	t.Log(XDayLast235959(time.Now().AddDate(0, 0, -4)))
	t.Log(XDayLast235959(time.Now()))
}

func TestXWeekFirst(t *testing.T) {
	t.Log(XWeekFirst(time.Now().AddDate(0, 0, -5)))
	t.Log(XWeekFirst(time.Now().AddDate(0, 0, -4)))
	t.Log(XWeekFirst(time.Now().AddDate(0, 0, -3)))
	t.Log(XWeekFirst(time.Now().AddDate(0, 0, -2)))
	t.Log(XWeekFirst(time.Now().AddDate(0, 0, -1)))
	t.Log(XWeekFirst(time.Now().AddDate(0, 0, 0)))
	t.Log(XWeekFirst(time.Now().AddDate(0, 0, 1)))
}

func TestXWeekLast(t *testing.T) {
	t.Log(XWeekLast(time.Now().AddDate(0, 0, -5)))
	t.Log(XWeekLast(time.Now().AddDate(0, 0, -4)))
	t.Log(XWeekLast(time.Now().AddDate(0, 0, -3)))
	t.Log(XWeekLast(time.Now().AddDate(0, 0, -2)))
	t.Log(XWeekLast(time.Now().AddDate(0, 0, -1)))
	t.Log(XWeekLast(time.Now().AddDate(0, 0, 0)))
	t.Log(XWeekLast(time.Now().AddDate(0, 0, 1)))
}

func TestXWeekLast235959(t *testing.T) {
	t.Log(XWeekLast235959(time.Now().AddDate(0, 0, -5)))
	t.Log(XWeekLast235959(time.Now().AddDate(0, 0, -4)))
	t.Log(XWeekLast235959(time.Now().AddDate(0, 0, -3)))
	t.Log(XWeekLast235959(time.Now().AddDate(0, 0, -2)))
	t.Log(XWeekLast235959(time.Now().AddDate(0, 0, -1)))
	t.Log(XWeekLast235959(time.Now().AddDate(0, 0, 0)))
	t.Log(XWeekLast235959(time.Now().AddDate(0, 0, 1)))
}

func TestXMonthFirst(t *testing.T) {
	t.Log(XMonthFirst(time.Now().AddDate(0, 0, -1)))
	t.Log(XMonthFirst(time.Now().AddDate(0, 0, 0)))
	t.Log(XMonthFirst(time.Now().AddDate(0, 0, 1)))
}

func TestXMonthLast(t *testing.T) {
	t.Log(XMonthLast(time.Now().AddDate(0, 0, -1)))
	t.Log(XMonthLast(time.Now().AddDate(0, 0, 0)))
	t.Log(XMonthLast(time.Now().AddDate(0, 0, 1)))
}

func TestXMonthLast235959(t *testing.T) {
	t.Log(XMonthLast235959(time.Now().AddDate(0, 0, -1)))
	t.Log(XMonthLast235959(time.Now().AddDate(0, 0, 0)))
	t.Log(XMonthLast235959(time.Now().AddDate(0, 0, 1)))
}

func TestXYearFirst(t *testing.T) {
	t.Log(XYearFirst(time.Now().AddDate(0, 0, -1)))
	t.Log(XYearFirst(time.Now().AddDate(0, 0, 0)))
	t.Log(XYearFirst(time.Now().AddDate(0, 0, 1)))
}

func TestXYearLast(t *testing.T) {
	t.Log(XYearLast(time.Now().AddDate(0, 0, -1)))
	t.Log(XYearLast(time.Now().AddDate(0, 0, 0)))
	t.Log(XYearLast(time.Now().AddDate(0, 0, 1)))
}

func TestXYearLast235959(t *testing.T) {
	t.Log(XYearLast235959(time.Now().AddDate(0, 0, -1)))
	t.Log(XYearLast235959(time.Now().AddDate(0, 0, 0)))
	t.Log(XYearLast235959(time.Now().AddDate(0, 0, 1)))
}
