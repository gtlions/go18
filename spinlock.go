package go18

import (
	"runtime"
	"sync/atomic"
)

const maxBackOff = 32

type SpinLock struct {
	l uint32
}

// Lock 加锁
func (sl *SpinLock) Lock() {
	backoff := 1
	// 自旋尝试获取锁
	for !atomic.CompareAndSwapUint32((&sl.l), 0, 1) {
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}
		if backoff < maxBackOff {
			backoff <<= 1
		}
	}
}

// UnLock 解锁
func (sl *SpinLock) UnLock() {
	atomic.CompareAndSwapUint32((&sl.l), 1, 0)
}
