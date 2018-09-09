package main

import "fmt"

func main() {
	arr := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	target := 33

	fmt.Println(BinarySearch(arr, target))
}

func BinarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (left + right) / 2
		potentialMatch := array[mid]
		if target == potentialMatch {
			return mid
		} else if target < potentialMatch {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
