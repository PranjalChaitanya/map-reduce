package main

import (
	"MapReduce"
)

func main() {
	buffer := &MapReduce.CircularBuffer{
		TotalSize:         22222256,
		SizeAvailable:     22222256,
		CurrentPoint:      0,
		CurrentCleanPoint: 0,
		BufferArray:       [22222256]string(make([]string, 22222256)),
		CurrentlyCleaning: false,
	}

	for i := 0; i < 49000000; i++ {
		buffer.Add("Hello")
		buffer.Add("World")
	}
}
