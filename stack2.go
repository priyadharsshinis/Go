package main

import (
	"fmt"
)

type stack []int

func (s *stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	} 
	return false
}

func (s *stack) push(val int) {
	
	*s = append(*s,val)
	
}

func (s *stack) pop () {
	if s.isEmpty() {
		fmt.Println("stack underflow, nothing to return")
		
	}
	index := len(*s)-1
	result := (*s)[index]
	*s = (*s)[:index]
	fmt.Println(result)
}

func main() {
	s := stack{}
	s.push(10)
	s.pop()
	s.push(20)
	s.push(30)
	s.push(40)
	s.pop()
	s.pop()
	s.pop()
	
}
