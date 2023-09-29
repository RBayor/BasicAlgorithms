package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	numIterations := 1
	creationTimes := []float64{}
	sortTimes := []float64{}
	messyArr, err := readSortableArray()
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	for i := 0; i < numIterations; i++ {
		cpy := make([]int, len(messyArr))
		copy(cpy, messyArr)
		creationTime, sortTime := benchmarkBubbleSort(cpy)
		creationTimes = append(creationTimes, creationTime)
		sortTimes = append(sortTimes, sortTime)
	}

	avgCreationTime := calculateAverage(creationTimes)
	avgSortTime := calculateAverage(sortTimes)

	// Calculate average times in seconds
	avgCreationTimeSec := avgCreationTime / 1000
	avgSortTimeSec := avgSortTime / 1000
	crTimes := floatSliceToStringSecs(creationTimes)
	srTimes := floatSliceToStringSecs(sortTimes)

	table := map[string]interface{}{
		"Sort Type":              "Bubble Sort",
		"Array Size":             len(messyArr),
		"Iterations":             numIterations,
		"Avg Creation Time (ms)": fmt.Sprintf("%.4fms", avgCreationTime),
		"Avg Sort Time (ms)":     fmt.Sprintf("%.4fms", avgSortTime),
		"Avg Creation Time (s)":  fmt.Sprintf("%.4fs", avgCreationTimeSec),
		"Avg Sort Time (s)":      fmt.Sprintf("%.4fs", avgSortTimeSec),
		"Creation Times (s)":     crTimes,
		"Sort Times (s)":         srTimes,
	}
	printTable(table)
}

func measureTime(operation func()) float64 {
	start := time.Now()
	operation()
	end := time.Now()
	return end.Sub(start).Seconds() * 1000
}

func benchmarkBubbleSort(messyArr []int) (float64, float64) {
	size := len(messyArr)
	creationTime := measureTime(func() { generateRandomArray(size) })
	sortTime := measureTime(func() { bubbleSort(messyArr) })

	return creationTime, sortTime
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000000)
	}
	return arr
}

func readSortableArray() ([]int, error) {
	filePath := "/Users/rb/Documents/Projects/GeneralProjects/BasicAlgorithms/data.json"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var res []int
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func bubbleSort(arr []int) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		swapped := false
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func calculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func floatSliceToStringSecs(numbers []float64) string {
	result := ""
	for i, num := range numbers {
		result += fmt.Sprintf("%.4f", num/1000)
		if i < len(numbers)-1 {
			result += ", "
		}
	}
	return result
}

func printTable(table map[string]interface{}) {
	maxKeyLen := 0
	for key := range table {
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}
	}

	for key, value := range table {
		fmt.Printf("%-*s %v\n", maxKeyLen, key, value)
	}
}
