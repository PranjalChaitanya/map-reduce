package main

import (
	"MapReduce"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Combiner(filename string) [][2]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var result [][2]string
	counts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		counts[line]++
	}

	for line, count := range counts {
		result = append(result, [2]string{line, strconv.Itoa(count)})
	}

	return result
}

func main() {
	buffer := &MapReduce.CircularBuffer{
		TotalSize:         72,
		SizeAvailable:     72,
		CurrentPoint:      0,
		CurrentCleanPoint: 0,
		BufferArray:       [72]string(make([]string, 72)),
		CurrentlyCleaning: false,
	}

	for i := 0; i < 400; i++ {
		buffer.Add("Hello")
		buffer.Add("World")
	}

	time.Sleep(5 * time.Second)

	combinedArr := Combiner("intermediate.txt")

	for _, pair := range combinedArr {
		fmt.Printf("(%s, %s)\n", pair[0], pair[1])
	}

	buffer.FlushRemainingBuffer()
}
