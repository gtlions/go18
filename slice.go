package go18

import (
	"fmt"
	"reflect"
)

// XIsInSlice 判断元素item是否在切片内
//
// item 元素
//
// parent 切片对象
func XIsInSlice(item interface{}, parent interface{}) (isExist bool, err error) {
	t := reflect.TypeOf(parent)
	v := reflect.ValueOf(parent)
	if t.Kind() != reflect.Slice {
		return isExist, fmt.Errorf("%s", "parent参数错误,应该传入slice")
	}
	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(item, v.Index(i).Interface()) {
			return true, nil
		}
	}
	return
}
