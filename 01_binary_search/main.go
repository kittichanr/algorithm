package main

import "fmt"

// list should be a sort list EX: [1,2,3,...,100]
// complexity time is O(log n)
func BinarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if item == guess {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	list := []int{1, 3, 5, 7, 9}
	index := BinarySearch(list, 3)
	fmt.Println("index", index)
}
