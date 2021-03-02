package main

import (
	"fmt"
)

func main() {
	input := []int{10, 7, 8, 9, 6,4, 3, 2, 1}
	result := mergeSort(input)
	fmt.Println(result)
}

func mergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	n := len(input) / 2
	l := mergeSort(input[:n])
	r := mergeSort(input[n:])
	return merge(l, r)
}

func merge(l []int, r []int) []int {

	//Slice: The size specifies the length. The capacity of the slice is
	//equal to its length. A second integer argument may be provided to
	//specify a different capacity; it must be no smaller than the
	//length. For example, make([]int, 0, 10) allocates an underlying array
	//of size 10 and returns a slice of length 0 and capacity 10 that is
	//backed by this underlying array.
	
	result := make([]int, 0, len(r)+len(l))
	for len(r) > 0 || len(l) > 0 {
		if len(r) == 0 {
			result := append(result, l...)
			return result
		} else if len(l) == 0 {
			result = append(result, r...)
			return result
		} else if l[0] <= r[0] {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}
	return result
}
