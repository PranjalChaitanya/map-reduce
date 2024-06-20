package main

import (
	"MapReduce"
	"bufio"
	"fmt"
	"os"
	"strconv"
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

//func main() {
//	// Call Combiner function
//	combinedArr := Combiner("input.txt")
//
//	// Print the combined array
//	for _, pair := range combinedArr {
//		fmt.Printf("(%s, %s)\n", pair[0], pair[1])
//	}
//}

func main() {
	buffer := &MapReduce.CircularBuffer{
		TotalSize:         222256,
		SizeAvailable:     222256,
		CurrentPoint:      0,
		CurrentCleanPoint: 0,
		BufferArray:       [222256]string(make([]string, 222256)),
		CurrentlyCleaning: false,
	}

	for i := 0; i < 490000; i++ {
		buffer.Add("Hello")
		buffer.Add("World")
	}

	combinedArr := Combiner("intermediate.txt")

	for _, pair := range combinedArr {
		fmt.Printf("(%s, %s)\n", pair[0], pair[1])
	}
}
