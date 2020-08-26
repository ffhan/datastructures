package linkedlist

import "errors"

var (
	IndexOutOfRangeErr = errors.New("index out of range")
	CorruptListErr     = errors.New("corrupt list")
)

type node struct {
	next  *node
	value interface{}
}

type singlyLinkedList struct {
	head *node
	tail *node
	size int
}

func NewSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{}
}

func (l *singlyLinkedList) Insert(index int, value interface{}) error {
	if index < 0 || index > l.size {
		return IndexOutOfRangeErr
	}
	if index == 0 {
		l.Prepend(value)
		return nil
	}
	if index == l.size {
		l.Append(value)
		return nil
	}
	i := 0
	for n := l.head; n != nil; n = n.next {
		if i == index-1 {
			newNode := &node{
				next:  n.next,
				value: value,
			}
			n.next = newNode
			l.size += 1
			return nil
		}
		i += 1
	}
	return CorruptListErr
}

func (l *singlyLinkedList) Prepend(value interface{}) {
	newNode := &node{next: l.head, value: value}
	l.head = newNode
	if l.tail == nil {
		l.tail = l.head
	}
	l.size += 1
}

func (l *singlyLinkedList) Append(value interface{}) {
	newNode := &node{value: value}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
		l.size = 1
		return
	}
	l.tail.next = newNode
	l.tail = newNode
	l.size += 1
}

func (l *singlyLinkedList) Size() int {
	return l.size
}

func (l *singlyLinkedList) Remove(index int) (interface{}, error) {
	if l.size == 0 || index < 0 || index >= l.size {
		return nil, IndexOutOfRangeErr
	}
	if index == 0 {
		oldHead := l.head
		if l.size == 1 {
			l.tail = nil
		}
		l.head = l.head.next
		l.size -= 1
		return oldHead.value, nil
	}
	i := 0
	for n := l.head; n != nil; n = n.next {
		if i == index-1 {
			forDeletion := n.next
			if forDeletion == l.tail {
				l.tail = n
			}
			n.next = forDeletion.next
			return forDeletion.value, nil
		}
		i += 1
	}
	return nil, CorruptListErr
}

func (l *singlyLinkedList) IsCorrupted() bool {
	return l.IsCircularAndCheckSizeMatches() || l.IsHeadCorrupted() || l.IsTailCorrupted()
}

func (l *singlyLinkedList) IsCircularAndCheckSizeMatches() bool {
	fast := l.head
	slow := l.head

	size := 0
	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next
		if fast != nil {
			fast = fast.next
		}
		if slow == fast {
			return true
		}
		size += 2
	}
	if fast != nil {
		size += 1
	}
	return size != l.size
}

func (l *singlyLinkedList) IsHeadCorrupted() bool {
	if l.size == 0 {
		return l.head != nil
	}
	return l.head != nil && l.head.next != nil
}

func (l *singlyLinkedList) IsTailCorrupted() bool {
	if l.size == 0 {
		return l.tail != nil
	}
	if l.size == 1 {
		return !(l.head == l.tail && l.tail.next == nil)
	}
	return !(l.tail != nil && l.head != l.tail && l.tail.next == nil)
}

func (l *singlyLinkedList) ToArray() []interface{} {
	result := make([]interface{}, 0, l.size)
	for node := l.head; node != nil; node = node.next {
		result = append(result, node.value)
	}
	return result
}

func (l *singlyLinkedList) Iterator() *iterator {
	return &iterator{current: l.head, index: 0}
}

type iterator struct {
	current *node
	index   int
}

func (i *iterator) Get() interface{} {
	if i.current == nil {
		return nil
	}
	return i.current.value
}

func (i *iterator) Index() int {
	return i.index
}

func (i *iterator) Next() {
	i.current = i.current.next
	i.index += 1
}

func (i *iterator) Has() bool {
	return i.current != nil
}
