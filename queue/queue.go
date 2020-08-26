package queue

import "algos/linkedlist"

type list interface {
	Append(value interface{})
	Remove(index int) (interface{}, error)
	Size() int
}

type queue struct {
	list list
}

func NewQueue() *queue {
	return &queue{list: linkedlist.NewSinglyLinkedList()}
}

func (s *queue) Push(value interface{}) {
	s.list.Append(value)
}

func (s *queue) Pop() (interface{}, error) {
	return s.list.Remove(0)
}

func (s *queue) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *queue) Size() int {
	return s.list.Size()
}
