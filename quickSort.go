package main

import (
	"fmt"
)

func main() {
	input := []int{10,20,3,4,90,-1}
	quicksort(input)
	fmt.Println(input)
}

func quicksort(input []int) {
	if len(input) <= 1 {
		return
	}
	partitionElement := partition(input)
	quicksort(input[0:partitionElement])
	quicksort(input[partitionElement+1 : ])
}

func partition(input []int)int {
	pivot := input[0]
	i, j := 0, len(input)-1
	
	//infinte for loop - go thur infinite for loop example for doubt

	for{
		//keep on moving the i till we find an element less than or equal pivot to swap
		for i < len(input) && input[i] <= pivot {
			i++
		} 
		//keep on moving j till we find an element greater than pivot to swap
		for input[j] >= pivot {
			j--
		}
		
		// this is the condition to break the infinite for loop
		// we break when start crosses end 
		//we break and then work on the divided array later same way
		if i >= j {
			// we swap the first element with the element at j , so that after swapping that element place will be fixed 
			input[0], input[j] = input[j], input[0]
			return j
		}
		
		// this is for normal swapping
		input[i], input[j] = input[j], input[i]
	}
}
