package MapReduce

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CleanerThread struct {
	StartingPoint int
	AmountToClean int
	CurrentBuffer []string
}

func MergeSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := arr[:mid]
	right := arr[mid:]

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left, right []string) []string {
	var result []string
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if strings.Compare(left[i], right[j]) < 0 {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func Combiner(arr []string) [][2]string {
	var result [][2]string
	var count int

	for i := 0; i < len(arr); i++ {
		count++

		// If it's the last element or the next element is different
		if i == len(arr)-1 || arr[i] != arr[i+1] {
			result = append(result, [2]string{arr[i], fmt.Sprint(count)})
			count = 0
		}
	}

	return result
}

func (ct *CleanerThread) StartCleanerThread(buffer *CircularBuffer, id int) {
	fmt.Println(buffer)
	ct.StartingPoint = buffer.CurrentCleanPoint
	//for {
	//	ct.AmountToClean -= 1
	//	if buffer.BufferArray[ct.StartingPoint%buffer.TotalSize] == " " {
	//		ct.StartingPoint += 1
	//		break
	//	}
	//	ct.StartingPoint += 1
	//}

	tmp := ""
	for {
		if buffer.BufferArray[ct.StartingPoint%buffer.TotalSize] == " " {
			ct.CurrentBuffer = append(ct.CurrentBuffer, tmp)
			tmp = ""
		} else {
			tmp += buffer.BufferArray[ct.StartingPoint%buffer.TotalSize]
		}
		if ct.AmountToClean <= 0 && buffer.BufferArray[ct.StartingPoint%buffer.TotalSize] == " " {
			if tmp != "" && tmp != " " {
				ct.CurrentBuffer = append(ct.CurrentBuffer, tmp)
			}
			buffer.CallCompletionThreadClean(id, ct.StartingPoint+1)
			break
		}
		ct.AmountToClean -= 1
		ct.StartingPoint += 1
	}

	sorted := MergeSort(ct.CurrentBuffer)
	fmt.Println(sorted)
	combined := Combiner(sorted)

	for _, pair := range combined {
		if (len(pair[0])) > 0 {
			file, err := os.OpenFile("intermediate.txt", os.O_APPEND|os.O_WRONLY, 0600)
			defer file.Close()
			if err != nil {

			}
			num, err := strconv.Atoi(pair[1])
			fmt.Fprintf(file, "%s, %d\n", pair[0], num)
		}
	}
}

func (ct *CleanerThread) FlushRemainingBuffer(buffer *CircularBuffer) {
	for i := buffer.CurrentCleanPoint; i < buffer.CurrentPoint; i++ {
	}
}
