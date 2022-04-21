package main

import (
	"fmt"
	"time"
)

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
	a := []int{171, 153, 82, 124, 25, 107, 138, 184, 123, 99, 121, 37, 119, 113, 67, 113, 141, 181, 121, 107, 32, 34, 66, 138, 80, 115, 174, 89, 71, 83, 61, 17, 190, 176, 112, 172, 112, 169, 65, 86, 141, 62, 69, 156, 20, 182, 182, 175, 90, 16, 1, 6, 123, 196, 195, 142, 60, 141, 128, 140, 105, 74, 184, 64, 75, 50, 58, 22, 90, 62, 89, 26, 33, 89, 147, 161, 129, 8, 124, 15, 40, 68, 87, 174, 106, 161, 59, 154, 37, 53, 131, 64, 169, 88, 143, 84, 97, 83, 4, 76, 21, 147, 95, 47, 149, 129, 151, 73, 121, 63, 73, 53, 55, 79, 120, 158, 20, 146, 169, 18, 71, 34, 64, 142, 105, 15, 63, 178, 142, 55, 92, 49, 67, 150, 123, 0, 26, 6, 130, 40, 139, 75, 7, 178, 88, 86, 21, 39, 69, 32, 90, 44, 4, 157, 154, 52, 119, 82, 116, 186, 138, 107, 54, 155, 96, 157, 175, 186, 49, 113, 144, 165, 97, 109, 48, 105, 163, 71, 95, 89, 66, 143, 125, 82, 65, 95, 153, 90, 159, 29, 69, 144, 68, 195, 43, 139, 120, 143, 164, 165, 96, 55, 99, 161, 187, 72, 155, 62, 99, 33, 23, 129, 100, 164, 12, 146, 43, 172, 131, 2, 87, 188, 41, 88, 67, 50, 25, 177, 88, 38, 46, 58, 64, 185, 164, 107, 181, 107, 60, 45, 5, 100, 197, 195, 93, 134, 30, 99, 171, 8, 179, 78, 158, 91, 74, 67, 70, 43, 121, 73, 84, 162, 53, 159, 111, 177, 76, 29, 82, 184, 101, 83, 186, 75, 116, 166, 59, 35, 173, 3, 167, 63, 81, 9, 172, 50, 95, 174, 123, 37, 41, 101, 15, 187, 188, 168, 185, 144, 72, 63, 98, 84, 25, 75, 189, 44, 134, 146, 124, 194, 68, 57, 16, 66, 18, 50, 160, 129, 75, 91, 184, 30, 187, 20, 127, 55, 22, 36, 38, 97, 55, 122, 40, 163, 90, 192, 73, 43, 92, 168, 65, 37, 10, 35, 147, 145, 94, 105, 165, 192, 196, 99, 141, 33, 80, 133, 2, 30, 122, 140, 55, 54, 102, 9, 102, 95, 40, 157, 27, 33, 89, 188, 88, 147, 179, 129, 5, 18, 87, 105, 7, 121, 132, 95, 73, 39, 51, 52, 101, 94, 98, 115, 151, 154, 83, 93, 18, 185, 118, 154, 157, 22, 10, 174, 70, 145, 79, 59, 90, 98, 68, 49, 89, 141, 146, 130, 134, 123, 152, 30, 63, 10, 130, 0, 98, 182, 90, 92, 50, 110, 99, 120, 45, 74, 76, 22, 12, 155, 64, 99, 107, 105, 130, 14, 180, 187, 163, 162, 154, 37, 59, 64, 59, 100, 45, 47, 72, 167, 52, 64, 83, 17, 11, 33, 193, 132, 7, 189, 164, 169, 56, 37, 165, 184, 21, 108, 47, 50, 4, 105, 71, 25, 171, 32, 118, 97, 187, 99, 27, 37, 50, 2, 80, 119, 55, 24, 187, 15, 22, 102, 3, 85, 56, 49, 29, 161, 35, 88, 36, 142, 32, 54, 61, 64, 13, 76, 51, 88, 3, 156, 50, 0, 193, 125, 26, 106, 2, 48, 7, 150, 105, 65, 64, 27, 148, 35, 66, 144, 25, 19, 152, 143, 24, 123, 87, 53, 138, 199, 185, 187, 61, 198, 186, 133, 1, 57, 186, 16, 127, 24, 11, 130, 140, 16, 76, 93, 43, 114, 112, 4, 39, 187, 197, 73, 3, 140, 148, 21, 64, 13, 46, 99, 38, 35, 192, 196, 145, 15, 181, 61, 118, 61, 166, 138, 41, 56, 41, 162, 98, 54, 86, 57, 20, 2, 110, 59, 176, 124, 45, 86, 17, 154, 35, 44, 94, 102, 121, 156, 1, 11, 170, 116, 37, 181, 63, 192, 148, 69, 42, 92, 184, 18, 129, 131, 129, 78, 87, 132, 54, 192, 181, 122, 152, 39, 185, 118, 161, 95, 176, 62, 39, 66, 40, 21, 14, 73, 187, 143, 20, 101, 53, 177, 70, 118, 92, 1, 108, 49, 68, 132, 146, 115, 44, 29, 178, 33, 188, 98, 59, 72, 63, 26, 3, 117, 183, 88, 184, 58, 118, 97, 26, 147, 73, 170, 18, 23, 15, 149, 47, 64, 126, 105, 14, 10, 91, 59, 167, 6, 17, 119, 2, 26, 16, 119, 186, 27, 158, 3, 118, 59, 190, 26, 134, 132, 169, 12, 124, 85, 34, 8, 176, 137, 177, 3, 94, 75, 121, 158, 45, 173, 11, 89, 18, 50, 191, 191, 164, 195, 77, 7, 85, 158, 171, 59, 74, 11, 187, 51, 72, 194, 131, 171, 191, 100, 180, 41, 106, 108, 196, 51, 161, 99, 86, 87, 32, 174, 117, 59, 27, 35, 25, 82, 179, 126, 55, 128, 180, 2, 20, 51, 27, 119, 44, 106, 10, 62, 160, 199, 88, 0, 77, 128, 161, 46, 77, 14, 30, 82, 112, 63, 158, 24, 131, 39, 156, 1, 118, 152, 86, 147, 89, 176, 118, 85, 21, 102, 188, 95, 59, 66, 48, 46, 174, 112, 35, 95, 107, 21, 108, 169, 128, 86, 16, 56, 99, 47, 122, 146, 90, 53, 130, 148, 7, 38, 83, 51, 198, 48, 110, 47, 94, 75, 131, 184, 20, 107, 177, 191, 128, 169, 70, 146, 16, 175, 164, 185, 107, 81, 62, 107, 54, 145, 10, 63, 198, 77, 173, 74, 82, 158, 198, 32, 64, 132, 16, 19, 79, 168, 1, 40, 54, 150, 189, 119, 188, 134, 35, 163, 135, 89, 186, 132, 53, 151, 149, 104, 108, 3, 187, 183, 165, 14, 49, 3, 12, 156, 163, 23, 3, 160, 145, 88, 196, 19, 78, 17, 176, 144, 110, 67, 31, 72, 45, 152, 123, 21, 187, 132, 54, 71, 123, 125, 51, 101, 75, 160, 61, 9, 109, 47, 143, 119, 116, 25, 155, 168, 105, 185, 138, 76, 148, 74, 82, 176, 35, 20, 28, 16, 97, 179, 47, 175, 33, 162, 75, 79, 45, 186, 85, 58, 143, 129, 17, 113, 48, 91, 123, 94, 79, 170}
	start := time.Now()
	sorted := mergeSort(a)
	duration := time.Since(start)
	fmt.Println("mergeSort:", duration)
	fmt.Println(sorted)
}
