package lru

import (
	"fmt"
	"strings"
)

type node struct {
	Prev  *node
	Next  *node
	Key   interface{}
	Value interface{}
}

func (n node) String() string {
	return fmt.Sprintf("(%v, %v)", n.Key, n.Value)
}

type queue struct {
	head *node
	tail *node
	size int
}

func (q queue) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	delimiter := ", "
	for node := q.head; node != nil; node = node.Next {
		if node.Next == nil {
			delimiter = "]"
		}
		line := fmt.Sprintf("%v%s", node, delimiter)
		sb.WriteString(line)
	}
	return sb.String()
}

func NewLinkedQueue() *queue {
	return &queue{}
}

func (q *queue) cut(n *node) {
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
	q.size -= 1
	if q.head == n {
		q.head = n.Next
	}
	if q.tail == n {
		q.tail = n.Prev
	}
	n.Prev = nil
	n.Next = nil
}

func (q *queue) push(node *node) {
	if q.head == nil {
		q.head = node
	}
	if q.tail != nil {
		q.tail.Next = node
	}
	q.tail = node
	q.size += 1
}

func (q *queue) Push(key, value interface{}) {
	newNode := &node{Prev: q.tail, Key: key, Value: value}
	q.push(newNode)
}

func (q *queue) pop() *node {
	node := q.head
	if node == nil {
		return nil
	}
	q.head = node.Next
	if q.head != nil {
		q.head.Prev = nil
	} else {
		q.tail = nil
	}
	q.size -= 1
	if q.size == 1 {
		q.tail = q.head
	}
	node.Next = nil
	node.Prev = nil
	return node
}

func (q *queue) Pop() (interface{}, interface{}) {
	if node := q.pop(); node != nil {
		return node.Key, node.Value
	}
	return nil, nil
}

type lru struct {
	elements map[interface{}]*node
	queue    queue
	capacity int
}

func (l lru) String() string {
	return fmt.Sprintf("{elements: %v, queue: %s, capacity: %d}", l.elements, l.queue, l.capacity)
}

func NewLru(capacity int) *lru {
	return &lru{
		elements: make(map[interface{}]*node),
		queue:    queue{},
		capacity: capacity,
	}
}

func (l *lru) Get(key interface{}) interface{} {
	val, ok := l.elements[key]
	if !ok {
		return nil
	}
	l.queue.cut(val)
	l.queue.push(val)
	return val
}

func (l *lru) Put(key, value interface{}) {
	if node, ok := l.elements[key]; ok {
		node.Value = value
		l.queue.cut(node)
		l.queue.push(node)
		return
	}
	if l.queue.size >= l.capacity {
		l.evict()
	}
	l.queue.Push(key, value)
	node := l.queue.tail
	l.elements[node.Key] = node
}

func (l *lru) evict() {
	node := l.queue.pop()
	delete(l.elements, node.Key)
}

func (l *lru) Size() int {
	return len(l.elements)
}
