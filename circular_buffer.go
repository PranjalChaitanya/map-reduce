package MapReduce

import (
	"fmt"
	"sync"
)

type CircularBuffer struct {
	mu                sync.Mutex
	TotalSize         int
	SizeAvailable     int
	CurrentPoint      int
	CurrentCleanPoint int
	BufferArray       [22222256]string
	CurrentlyCleaning bool
}

func (c *CircularBuffer) Add(value string) {
	if len(value) > c.SizeAvailable {
		fmt.Println("PERFORMING FULL CLEAN")
		go c.Clean()
		for len(value) >= c.SizeAvailable {
			// Infinite loop till the go routine cleans up
		}
	}

	casted := int(0.8 * float64(c.TotalSize))
	if (c.TotalSize - c.SizeAvailable) > casted {
		fmt.Println("PERFORMING PARTIAL CLEAN")
		go c.Clean()
	}

	for _, char := range value {
		c.BufferArray[c.CurrentPoint%c.TotalSize] = string(char)
		c.CurrentPoint += +1
		c.mu.Lock()
		c.SizeAvailable -= 1
		c.mu.Unlock()
	}

	c.BufferArray[c.CurrentPoint%c.TotalSize] = " "
	c.CurrentPoint += len(value) + 1
	c.mu.Lock()
	c.SizeAvailable -= 1
	c.mu.Unlock()
}

func (c *CircularBuffer) Clean() {
	charactersToClean := int(0.8 * float64(c.TotalSize))

	for charactersToClean >= 0 {
		c.CurrentCleanPoint += 1
		c.mu.Lock()
		c.SizeAvailable += 1
		c.mu.Unlock()
		charactersToClean -= 1
	}
}
