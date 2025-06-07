package main

import (
	"fmt"
	"hedera/cp/sharedpkg"
	"slices"
	"strconv"
	"strings"
)

// Use stack first
type stack struct {
	data []int
}

func (s *stack) pop() {
	if len(s.data) > 0 {
		s.data = slices.Delete(s.data, len(s.data)-1, len(s.data))
	}
}

func (s *stack) insert(element int) {
	s.data = slices.Insert(s.data, len(s.data), element)
}

func newstack() *stack {
	return &stack{
		data: []int{},
	}
}

func main() {
	l := sharedpkg.Scanner()
	stack := newstack()
	for i, v := range l {
		num, _ := strconv.Atoi(v)
		stack.insert(num)

		if i == len(l)-1 {
			stack.pop()
		}
	}
	fmt.Print(strings.Trim(fmt.Sprintf("%v", stack.data), "[]"))
}
