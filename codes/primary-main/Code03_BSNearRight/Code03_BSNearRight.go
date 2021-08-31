package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func nearestIndex(arr []int, value int) int {
	L := 0
	R := len(arr) - 1
	index := -1
	for L <= R {
		mid := L + (R-L)>>1
		if arr[mid] <= value {
			index = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return index
}

// for test
func test(arr []int, value int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] <= value {
			return i
		}
	}
	return -1
}

func generateRandomArray(maxSize int, maxValue int) []int {
	arr := make([]int, (int)(float32(maxSize+1)*rand.Float32()))
	for i := 0; i < len(arr); i++ {
		arr[i] = int(float32(maxValue)*rand.Float32()) - int(float32(maxValue)*rand.Float32())
	}
	return arr
}

func printArray(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func main() {
	testTime := 500000
	maxSize := 10
	maxValue := 100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		sort.Sort(sort.IntSlice(arr))
		value := int(float32(maxValue+1)*rand.Float32()) - int(float32(maxValue)*rand.Float32())
		if test(arr, value) != nearestIndex(arr, value) {
			printArray(arr)
			fmt.Println(value)
			fmt.Println(test(arr, value))
			fmt.Println(nearestIndex(arr, value))
			succeed = false
			break
		}
	}
	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}
}
