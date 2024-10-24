package clock

import (
	"sync"
)

type LamportClock struct {
	clock uint64
	mutex sync.Mutex
}

func (c *LamportClock) Tick() uint64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clock += 1
	return c.clock
}

func (c *LamportClock) Advance(new_time uint64) uint64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clock = max(c.clock, new_time) + 1
	return c.clock
}
