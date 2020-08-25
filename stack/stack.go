package stack

import "algos/linkedlist"

type list interface {
	Append(value interface{})
	Remove(index int) (interface{}, error)
	Size() int
}

type stack struct {
	list list
}

func NewStack() *stack {
	return &stack{list: linkedlist.NewSinglyLinkedList()}
}

func (s *stack) Push(value interface{}) {
	s.list.Append(value)
}

func (s *stack) Pop() (interface{}, error) {
	return s.list.Remove(0)
}

func (s *stack) Size() int {
	return s.list.Size()
}
