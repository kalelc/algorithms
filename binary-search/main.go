package main

import "fmt"

func main() {
	elements := []int{1, 3, 5, 7, 9}
	value := 3
	indexPosition := binarySearch(elements, value)

	if indexPosition == -1 {
		fmt.Println("element does not exist")
	} else {
		fmt.Printf("index position %d\n", indexPosition)
	}
}

func binarySearch(elements []int, value int) int {
	low := 0
	high := len(elements) - 1
	for i := 0; low <= high; i++ {
		mid := low + high
		guess := elements[mid]
		if guess == value {
			return mid
		}
		if guess > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
