package main

import (
	"fmt"
)

func main() {
	numbers := []int{5, 3, 6, 2, 10, 11, 1, 90}

	BubbleSort(&numbers)
	fmt.Println(numbers)
}

func BubbleSort(numbers *[]int) {
	list := *numbers
	for i := 1; i < len(list); i++ {
		Swap(&list, i)
	}
}

func Swap(numbers *[]int, i int) {
	list := *numbers
	for j := 0; j < len(list)-i; j++ {
		if list[j] > list[j+1] {
			temp := list[j]
			list[j] = list[j+1]
			list[j+1] = temp
		}
	}
}
