package main

import (
	"fmt"
)

func findSmallest(arr ...int) (int, int) {
	smallest := arr[0]
	smallestIndex := 0

	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}

	return smallestIndex, smallest
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func selectionSort(arr ...int) {
	newArr := []int{}
	for _, _ = range arr {
		i, v := findSmallest(arr...)
		newArr = append(newArr, v)
		if len(arr) > 1 {
			arr = removeIndex(arr, i)
		}
	}
	fmt.Println("Ordered array:", newArr)
}

func main() {
	arr := []int{7, 5, 3, 8, 4, 2, 11, 22, 2122, 0, 2, 3, 88}
	selectionSort(arr...)
}
