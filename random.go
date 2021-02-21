package gos10i

import (
	"math/rand"
	"time"
)

// XRandomInt 生成指定长度的随机数字字符串
//
// length 字符串长度
//
func XRandomInt(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// XRandomString 生成指定长度的随机字符串
//
// length 字符串长度
//
func XRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// XRandomIntRange 生成指定区间的整数随机数
//
// min 最小值
//
// max 最大值
//
func XRandomIntRange(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}
