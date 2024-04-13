package main

import (
	"fmt"
	"slices"
)

// run-time complexity = Best case O(n log(n))
// Average case O(n log(n))
// Worst case O(n^2) if already sorted
func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	var less []int
	var greater []int

	mid := len(list) / 2
	pivot := list[mid]

	for i, v := range list {
		if i != mid {
			if v <= pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}
	}

	concat := slices.Concat(quickSort(less), []int{pivot}, quickSort(greater))

	return concat
}

func quickSortPivotFirst(list []int) []int {
	if len(list) < 2 {
		return list
	}

	var less []int
	var greater []int

	pivot := list[0]

	for _, v := range list[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	concat := slices.Concat(quickSort(less), []int{pivot}, quickSort(greater))

	return concat
}

func main() {
	list := []int{10, 5, 2, 3, 1}

	fmt.Println(quickSort(list))
	fmt.Println(quickSortPivotFirst(list))
}
