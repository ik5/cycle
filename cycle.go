package cycle

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// CyclingInterface holds interface for what to expect from cycling
type CyclingInterface interface {
	Cycle() interface{}
	String() string
	Value() interface{}
}

// IntCycle holds internal values for doing cycling of values
type IntCycle struct {
	minValue int64
	maxValue int64
	value    int64

	rwMutex *sync.RWMutex
}

// UintCycle holds internal values for doing cycling of values
type UintCycle struct {
	minValue uint64
	maxValue uint64
	value    uint64

	rwMutex *sync.RWMutex
}

// InitIntCycle create a Cycle struct to start cycling
func InitIntCycle(minValue, startAt, endAt int64) *IntCycle {
	return &IntCycle{
		minValue: minValue,
		maxValue: endAt,
		value:    startAt,
		rwMutex:  &sync.RWMutex{},
	}
}

// Cycle do cycling and return the current new value
func (c *IntCycle) Cycle() int64 {
	defer c.rwMutex.Unlock()
	c.rwMutex.Lock()
	if c.value == c.maxValue {
		atomic.StoreInt64(&c.value, c.minValue)
		return c.value
	}

	value := atomic.AddInt64(&c.value, 1)
	if value > c.maxValue {
		atomic.StoreInt64(&c.value, c.minValue)
		return c.value
	}
	return c.value
}

// Value return the value of the current cycle without changing it's value
func (c *IntCycle) Value() int64 {
	defer c.rwMutex.RUnlock()
	c.rwMutex.RLock()
	return c.value
}

func (c *IntCycle) String() string {
	defer c.rwMutex.RUnlock()
	c.rwMutex.RLock()
	return fmt.Sprintf("%d", c.value)
}
