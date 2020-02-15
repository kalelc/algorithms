package main

import "fmt"

func main() {
	elements := []int{10, 5, 2, 3, 14, 11, 8}
	fmt.Println(quickSort(elements))
}

func quickSort(elements []int) []int {
	var less []int
	var greater []int
	var result []int
	var pivot int

	if len(elements) < 2 {
		result = elements
	} else {
		pivot = elements[0]

		for _, num := range elements[1:] {
			if num <= pivot {
				less = append(less, num)
			}
		}

		for _, num := range elements[1:] {
			if num > pivot {
				greater = append(greater, num)
			}
		}

		result = append(result, quickSort(less)...)
		result = append(result, pivot)
		result = append(result, quickSort(greater)...)
	}
	return result
}
