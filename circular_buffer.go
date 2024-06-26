package MapReduce

import (
	"fmt"
	"sync"
)

type ThreadCleanerStatus int

const (
	NOT_CLEANING ThreadCleanerStatus = iota
	CLEANING
	DONE_CLEANING
)

type CircularBuffer struct {
	mu                sync.Mutex
	TotalSize         int
	SizeAvailable     int
	CurrentPoint      int
	CurrentCleanPoint int
	BufferArray       [77]string
	CurrentlyCleaning bool
	ThreadCleaner     [1]int
}

func (c *CircularBuffer) Add(value string) {
	c.SizeAvailable = c.TotalSize - (c.CurrentPoint - c.CurrentCleanPoint)
	if c.CurrentPoint == 77 {
		fmt.Println("77 PRINT STATS")
		fmt.Println(len(value))
		fmt.Println(c.SizeAvailable)
		fmt.Println(c.CurrentlyCleaning)
	}
	if (len(value)+1) > c.SizeAvailable && c.CurrentlyCleaning == false {
		fmt.Println("PERFORMING FULL CLEAN")
		fmt.Println(c.CurrentCleanPoint)
		fmt.Println(c.CurrentPoint)
		fmt.Println(c.SizeAvailable)
		go c.Clean()
		for len(value) >= c.SizeAvailable {
			// Infinite loop till the go routine cleans up
			c.SizeAvailable = c.TotalSize - (c.CurrentPoint - c.CurrentCleanPoint)
		}
	}

	//casted := int(0.8 * float64(c.TotalSize))
	//if (c.TotalSize-c.SizeAvailable) > casted && c.CurrentlyCleaning == false {
	//	fmt.Println("PERFORMING PARTIAL CLEAN")
	//	go c.Clean()
	//}

	for _, char := range value {
		c.BufferArray[c.CurrentPoint%c.TotalSize] = string(char)
		c.CurrentPoint += +1
		c.mu.Lock()
		c.SizeAvailable -= 1
		c.mu.Unlock()
	}

	c.BufferArray[c.CurrentPoint%c.TotalSize] = " "
	c.CurrentPoint += 1
	c.mu.Lock()
	c.SizeAvailable -= 1
	c.mu.Unlock()
}

func (c *CircularBuffer) Clean() {
	c.CurrentlyCleaning = true
	c.ThreadCleaner = [1]int{-1}
	cleaner1 := CleanerThread{c.CurrentCleanPoint, int(0.8 * float64(c.CurrentPoint-c.CurrentCleanPoint)), make([]string, 12)}
	//cleaner2 := CleanerThread{c.CurrentCleanPoint + 1000, 1000, make([]string, 1000)}
	//cleaner3 := CleanerThread{c.CurrentCleanPoint + 2000, 1000, make([]string, 1000)}

	go cleaner1.StartCleanerThread(c, 0)
	//go cleaner2.StartCleanerThread(c, 1)
	//go cleaner3.StartCleanerThread(c, 2)
}

func (c *CircularBuffer) CallCompletionThreadClean(id int, end int) {
	fmt.Println("COMPLETION THREAD IS CALLED")
	c.mu.Lock()
	defer c.mu.Unlock()
	//c.ThreadCleaner[id] = end
	//for i := 0; i < len(c.ThreadCleaner); i++ {
	//	if c.ThreadCleaner[i] == -1 {
	//		return
	//	}
	//	//// If the new clean point is greaetr than the current you give that much space
	//	//if c.ThreadCleaner[i] > c.CurrentCleanPoint {
	//	//	c.SizeAvailable += c.ThreadCleaner[i] - c.CurrentCleanPoint
	//	//}
	//	c.CurrentCleanPoint = end
	//}

	fmt.Println("COMPLETION THREAD DATA")
	fmt.Println(c.CurrentPoint)
	fmt.Println(end)

	c.CurrentCleanPoint = end
	c.ThreadCleaner = [1]int{-1}
	c.CurrentlyCleaning = false
	c.SizeAvailable = c.TotalSize - (c.CurrentPoint - c.CurrentCleanPoint)

	fmt.Println(c.SizeAvailable)

}

func (c *CircularBuffer) FlushRemainingBuffer() {
	//fmt.Println(c.CurrentCleanPoint)
	//fmt.Println(c.CurrentPoint)
	//fmt.Println(c.SizeAvailable)
	//fmt.Println("Flushing remaining buffer")
	//tmp := ""
	//
	//for i := c.CurrentCleanPoint; i < c.CurrentPoint; i++ {
	//	if c.BufferArray[i%c.TotalSize] == " " {
	//
	//	}
	//	tmp += string(c.BufferArray[i*c.TotalSize])
	//}
	cleaner1 := CleanerThread{c.CurrentCleanPoint, int(0.8 * float64(c.CurrentPoint-c.CurrentCleanPoint)), make([]string, 12)}
	cleaner1.FlushRemainingBuffer(c)
}
