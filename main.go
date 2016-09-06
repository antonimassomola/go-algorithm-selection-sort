package main

import (
	"fmt"
)

func findSmallest(arr ...int) (int, int) {
	smallest := arr[0]
	smallest_index := 0

	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallest_index = i
		}
	}

	return smallest_index, smallest
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func selectionSort(arr ...int) {
	newArr := []int{}
	for _, _ = range arr {
		index, value := findSmallest(arr...)
		newArr = append(newArr, value)
		if len(arr) > 1 {
			arr = removeIndex(arr, index)
		}
	}
	fmt.Println(newArr)
}

func main() {
	arr := []int{7,5,3,8,4,2,11,22,2122,0,2,3,88}
	selectionSort(arr...)
}
