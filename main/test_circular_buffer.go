package main

import "MapReduce"

func main() {
	buffer := &MapReduce.CircularBuffer{
		TotalSize:         72,
		SizeAvailable:     72,
		CurrentPoint:      0,
		CurrentCleanPoint: 0,
		BufferArray:       [72]string(make([]string, 72)),
		CurrentlyCleaning: false,
	}

	for i := 0; i < 1200; i++ {
		buffer.Add("Hello")
		buffer.Add("World")
	}
}
