package main

import "fmt"

func findSmallest(list []int) int {
	smallest := list[0]
	smallestIndex := 0

	for i, v := range list {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}

	return smallestIndex
}

// time complexity O(n)
func selectionSort(list []int) []int {
	newArr := make([]int, len(list))

	for i := range list {
		smallestIndex := findSmallest(list)
		newArr[i] = list[smallestIndex]
		list = append(list[:smallestIndex], list[smallestIndex+1:]...) // pop value that update to newArr
	}
	return newArr
}

func main() {
	list := []int{5, 3, 6, 2, 10}
	fmt.Println("selection sort", selectionSort(list))
}
