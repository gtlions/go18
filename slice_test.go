package gos10i

import "testing"

func TestXInSlice(t *testing.T) {
	t.Log(XIsInSlice("a", []string{"a", "b"}))
	t.Log(XIsInSlice("c", []string{"e", "e"}))
	t.Log(XIsInSlice(1, []int{1, 2}))
	t.Log(XIsInSlice(3, []int{4, 5}))
	t.Log(XIsInSlice("a", []int{1, 2}))
}
