package main

import (
	"fmt"
	"time"
)

// Implementation of the Merge Sort algorithm.
// It takes an array of integers and sorts it in ascending order.
// It uses the merge function to merge two sorted arrays.
// It returns the sorted array.
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	return result
}

func main() {
	a := []int{183, 121, 76, 32, 159, 52, 160, 102, 114, 105, 145, 180, 32, 93, 10, 67, 73,
		188, 180, 158, 52, 24, 72, 33, 19, 147, 64, 100, 115, 133, 178, 69, 129, 107,
		77, 152, 87, 29, 78, 28, 120, 100, 110, 149, 180, 189, 69, 12, 123, 179, 20,
		81, 3, 157, 90, 78, 44, 116, 135, 188, 33, 102, 97, 49, 30, 197, 75, 5, 66,
		168, 69, 58, 185, 18, 17, 49, 54, 7, 40, 154, 60, 104, 17, 7, 63, 97, 119, 25,
		64, 62, 196, 25, 166, 157, 193, 46, 24, 95, 31, 35}
	start := time.Now()
	sorted := mergeSort(a)
	duration := time.Since(start)
	fmt.Println("mergeSort:", duration)
	fmt.Println(sorted)
}
