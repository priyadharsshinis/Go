package main

import (
	"fmt"
)

var stack []int

func isEmpty() bool {
	if len(stack) == 0 {
		return true
	} 
	return false
}

func push(val int) {
	
	stack = append(stack,val)
	
}

func pop () {
	if isEmpty() {
		fmt.Println("stack underflow, nothing to return")
		
	}
	index := len(stack)-1
	result := (stack)[index]
	stack = stack[:index]
	fmt.Println(result)
}

func main() {
	
	push(10)
	pop()
	push(20)
	push(30)
	push(40)
	pop()
	pop()
	pop()
	
}
