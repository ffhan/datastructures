package linkedlist

type doubleNode struct {
	prev  *doubleNode
	next  *doubleNode
	value interface{}
}

type doublyLinkedList struct {
	head *doubleNode
	tail *doubleNode
	size int
}

func NewDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func (l *doublyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, IndexOutOfRangeErr
	}
	isReversed := index > l.size/2
	for iter := l.Iterator(isReversed); iter.Has(); iter.Next() {
		if index == iter.Index() {
			return iter.Get(), nil
		}
	}
	return nil, CorruptListErr
}

func (l *doublyLinkedList) Insert(index int, value interface{}) error {
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
	isReversed := index > l.size/2
	for iter := l.Iterator(isReversed); iter.Has(); iter.Next() {
		if iter.Index() == index {
			n := iter.getCurrent()
			newNode := &doubleNode{prev: n.prev, next: n, value: value}
			n.prev.next = newNode
			n.prev = newNode
			l.size += 1
			return nil
		}
	}
	return CorruptListErr
}

func (l *doublyLinkedList) Prepend(value interface{}) {
	newNode := &doubleNode{next: l.head, value: value}
	if l.head != nil {
		l.head.prev = newNode
	}
	l.head = newNode
	if l.tail == nil {
		l.tail = l.head
	}
	l.size += 1
}

func (l *doublyLinkedList) Append(value interface{}) {
	newNode := &doubleNode{prev: l.tail, value: value}
	if l.tail != nil {
		l.tail.next = newNode
	}
	l.tail = newNode
	if l.head == nil {
		l.head = l.tail
	}
	l.size += 1
}

func (l *doublyLinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, IndexOutOfRangeErr
	}
	isReversed := index > l.size/2
	for iter := l.Iterator(isReversed); iter.Has(); iter.Next() {
		if iter.Index() == index {
			n := iter.getCurrent()
			if n == l.head {
				l.head = n.next
			}
			if n == l.tail {
				l.tail = n.prev
			}
			if n.next != nil {
				n.next.prev = n.prev
			}
			if n.prev != nil {
				n.prev.next = n.next
			}
			l.size -= 1
			return n.value, nil
		}
	}
	return nil, CorruptListErr
}

func (l *doublyLinkedList) ToArray() []interface{} {
	result := make([]interface{}, 0, l.size)
	for node := l.head; node != nil; node = node.next {
		result = append(result, node.value)
	}
	return result
}

type doubleIterator struct {
	current  *doubleNode
	index    int
	reversed bool
}

func (l *doublyLinkedList) Iterator(reversed bool) *doubleIterator {
	if reversed {
		return &doubleIterator{
			current:  l.tail,
			index:    l.size - 1,
			reversed: reversed,
		}
	}
	return &doubleIterator{
		current:  l.head,
		index:    0,
		reversed: reversed,
	}
}

func (i *doubleIterator) Get() interface{} {
	if i.current == nil {
		return nil
	}
	return i.current.value
}

func (i *doubleIterator) getCurrent() *doubleNode {
	return i.current
}

func (i *doubleIterator) Has() bool {
	return i.current != nil
}

func (i *doubleIterator) Next() {
	if i.reversed {
		i.current = i.current.prev
		i.index -= 1
	} else {
		i.current = i.current.next
		i.index += 1
	}
}

func (i *doubleIterator) Index() int {
	return i.index
}
