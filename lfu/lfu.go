package lfu

import (
	"fmt"
	"math"
)

type node struct {
	Key   interface{}
	Freq  int
	Index int
	Value interface{}
}

func (n node) String() string {
	return fmt.Sprintf("{key: %v, value: %v, freq: %d, index: %d}", n.Key, n.Value, n.Freq, n.Index)
}

type lfu struct {
	nodes    []*node
	elements map[interface{}]*node
	capacity int
}

func NewLfu(capacity int) *lfu {
	return &lfu{
		nodes:    make([]*node, 0, capacity),
		elements: make(map[interface{}]*node),
		capacity: capacity,
	}
}

func toIndex(depth, index int) int {
	return int(math.Exp2(float64(depth))) - 1 + index
}

/*
0 1 2 3 4 5 6
            0
           / \
          1   2
         / \  /\
        3  4 5  6
       /\ /\/\ / \
      7 89 10

*/
func parent(index int) int {
	if index == 0 {
		return -1
	}
	return (index - 1) / 2
}

func children(index int) (int, int) {
	return index*2 + 1, index*2 + 2
}

func (l *lfu) upHeap(index int) {
	if index >= len(l.nodes) || index == 0 {
		return
	}
	parentIndex := parent(index)
	parent := l.nodes[parentIndex]
	node := l.nodes[index]

	if parent.Freq > node.Freq {
		l.nodes[index] = parent
		l.nodes[parentIndex] = node
		node.Index = parentIndex
		parent.Index = index
		l.upHeap(parentIndex)
	}
}

func (l *lfu) downHeap(index int) {
	if index >= len(l.nodes) {
		return
	}
	leftI, rightI := children(index)
	parent := l.nodes[index]
	var left, right *node
	if leftI < len(l.nodes) {
		left = l.nodes[leftI]
	}
	if rightI < len(l.nodes) {
		right = l.nodes[rightI]
	}
	if left == nil && right == nil {
		return
	}

	if left != nil && parent.Freq > left.Freq {
		l.nodes[index] = left
		left.Index = index

		l.nodes[leftI] = parent
		parent.Index = leftI
		l.downHeap(leftI)
	} else if right != nil && parent.Freq > right.Freq {
		l.nodes[index] = right
		right.Index = index

		l.nodes[rightI] = parent
		parent.Index = rightI
		l.downHeap(rightI)
	}
}

func (l *lfu) evictAndAdd(key, value interface{}) {
	oldKey := l.nodes[0].Key
	l.nodes[0] = &node{
		Key:   key,
		Freq:  1,
		Index: 0,
		Value: value,
	}
	delete(l.elements, oldKey)
	l.elements[key] = l.nodes[0]
}

func (l *lfu) Get(key interface{}) interface{} {
	if val, ok := l.elements[key]; ok {
		val.Freq += 1
		l.downHeap(val.Index)
		return val.Value
	}
	return nil
}

func (l *lfu) Put(key, value interface{}) {
	if val, ok := l.elements[key]; ok {
		val.Value = value
		val.Freq += 1
		l.downHeap(val.Index)
		return
	}
	if len(l.nodes) >= l.capacity {
		l.evictAndAdd(key, value)
		return
	}
	newNode := &node{
		Key:   key,
		Freq:  1,
		Index: len(l.nodes),
		Value: value,
	}
	l.nodes = append(l.nodes, newNode)
	l.elements[key] = newNode
	l.upHeap(newNode.Index)
}

func (l *lfu) Size() int {
	return len(l.nodes)
}
