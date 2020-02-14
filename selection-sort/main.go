package main

import "fmt"

func main() {
	elements := []int{5, 3, 6, 2, 10}
	fmt.Println(elements)
	selectionSort(elements)
}

func findSmallest(elements []int) int {
	smallest := elements[0]
	smallestIndex := 0

	for i := 1; i < len(elements); i++ {
		if elements[i] < smallest {
			smallest = elements[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(elements []int) {
	var newElements []int
	for i := 0; i < len(elements); i++ {
		smallest := findSmallest(elements)
		newElements = append(newElements, elements[smallest])
		removeElement(elements, smallest)
	}
	fmt.Println(newElements)
}

func removeElement(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
