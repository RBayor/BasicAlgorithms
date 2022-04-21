package main

import (
	"fmt"
	"time"
)

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		currentVal := arr[i]
		j := i - 1

		for j >= 0 && currentVal < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = currentVal
	}

	return arr
}

func main() {
	a := []int{183, 121, 76, 32, 159, 52, 160, 102, 114, 105, 145, 180, 32, 93, 10, 67, 73, 188, 180, 158, 52, 24, 72, 33, 19, 147, 64, 100, 115, 133, 178, 69, 129, 107, 77, 152, 87, 29, 78, 28, 120, 100, 110, 149, 180, 189, 69, 12, 123, 179, 20, 81, 3, 157, 90, 78, 44, 116, 135, 188, 33, 102, 97, 49, 30, 197, 75, 5, 66, 168, 69, 58, 185, 18, 17, 49, 54, 7, 40, 154, 60, 104, 17, 7, 63, 97, 119, 25, 64, 62, 196, 25, 166, 157, 193, 46, 24, 95, 31, 35}

	start := time.Now()
	sorted := insertionSort(a)
	//insertionSort(a)
	duration := time.Since(start)

	fmt.Println(sorted)
	fmt.Println("insertionSort:", duration)
}