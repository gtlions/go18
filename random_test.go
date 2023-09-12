package go18

import "testing"

func TestXRandomInt(t *testing.T) {
	t.Log(XRandomInt(12))
}

func TestXRandomString(t *testing.T) {
	t.Log(XRandomString(12))
}

func TestXRandomIntRange(t *testing.T) {
	t.Log(XRandomIntRange(10, 20))
}
