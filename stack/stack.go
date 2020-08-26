package stack

import "errors"

var (
	EmptyStackErr = errors.New("stack is empty")
)

type stack struct {
	values []interface{}
}

func NewStack() *stack {
	return &stack{values: make([]interface{}, 0)}
}

func (q *stack) Push(value interface{}) {
	q.values = append(q.values, value)
}

func (q *stack) Pop() (interface{}, error) {
	if len(q.values) == 0 {
		return nil, EmptyStackErr
	}
	val := q.values[len(q.values)-1]
	q.values = q.values[:len(q.values)-1]
	return val, nil
}

func (q *stack) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *stack) Size() int {
	return len(q.values)
}
