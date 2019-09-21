package util

import (
	"sync/atomic"
)

type Counter interface {
	Count() int64
	Inc(int64)
	Dec(int64)
	Snapshot() Counter
	Clear()
}

func NewCounter() Counter {
	return &StandardCounter{
		count: 0,
	}
}

type StandardCounter struct {
	count int64
}

func (c *StandardCounter) Count() int64 {
	return atomic.LoadInt64(&c.count)
}

func (c *StandardCounter) Inc(count int64) {
	atomic.AddInt64(&c.count, count)
}

func (c *StandardCounter) Dec(count int64) {
	atomic.AddInt64(&c.count, -count)
}

func (c *StandardCounter) Snapshot() Counter {
	tmp := &StandardCounter{
		count: atomic.LoadInt64(&c.count),
	}
	return tmp
}

func (c *StandardCounter) Clear() {
	atomic.StoreInt64(&c.count, 0)
}
