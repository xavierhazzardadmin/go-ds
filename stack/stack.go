package main

import (
	"fmt"
	"reflect"
)

type Stack struct {
	items []any
}

func (s *Stack) Peek() any {
	length := len(s.items)
	items := s.items[length-1]
	return items
}

func (s *Stack) Push(item any) {
	if len(s.items) == 0 {
		s.items = append(s.items, item)
	} else if reflect.TypeOf(s.items[0]) != reflect.TypeOf(item) {
		fmt.Printf("Item isn't the same type as slice.\n")
	} else {
		s.items = append(s.items, item)
	}
}

func (s *Stack) PushMany(items ...any) {
	for _, item := range items {
		if reflect.TypeOf(s.items[0]) != reflect.TypeOf(item) {
			fmt.Printf("Item isn't the same type as slice.\n")
			return
		}
	}

	if len(items) == 1 {
		fmt.Printf("Only one Item in array, use Push instead.\n")
		return
	}

	s.items = append(s.items, items...)
}

func (s *Stack) Pop() any {
	if s.IsEmpty() {
		fmt.Printf("Error: Stack is empty.\n")
		return nil
	}
	length := len(s.items)
	popped := s.items[length-1]
	s.items = s.items[0 : len(s.items)-2]

	return popped

}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Stack is working!")
	s1 := "Taylor Swift"
	s2 := "Claire Rosinkranz"
	s3 := "Olivia Rodrigo"
	s4 := "Sabrina Carpenter"

	stack := &Stack{}

	stack.Push(s1)
	stack.Push(s2)
	stack.Push(18)
	stack.Push(s3)
	stack.Push(s4)

	fmt.Println("Last item: ", stack.Peek())

	stack.Pop()

	fmt.Println("Stack: ", stack.items)

}
