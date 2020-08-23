package lru

import "testing"

func TestLinkedQueue_Push(t *testing.T) {
	l := &queue{}
	k1 := 1
	v1 := 123
	l.Push(k1, v1)
	size := 2
	k2 := size
	v2 := 50
	l.Push(k2, v2)
	if l.head.Key != k1 || l.head.Value != v1 {
		t.Errorf("expected (%v, %v), got (%v, %v)", k1, v1, l.head.Key, l.head.Value)
	}
	if l.head.Next.Key != k2 || l.head.Next.Value != v2 {
		t.Errorf("expected (%v, %v), got (%v, %v)", k2, v2, l.head.Next.Key, l.head.Next.Value)
	}
	if l.tail.Prev != l.head {
		t.Error("tail prev has to be equal to head")
	}
	if l.head.Next != l.tail {
		t.Error("head next has to be equal to tail")
	}
	if l.head.Prev != nil {
		t.Error("head prev has to be nil")
	}
	if l.tail.Next != nil {
		t.Error("tail next has to be nil")
	}
	if l.size != size {
		t.Errorf("expected size %d, got %d", size, l.size)
	}
}

func TestLinkedQueue_Pop(t *testing.T) {
	l := &queue{}
	k := 1
	v := 2
	l.Push(k, v)
	k2 := 2
	v2 := 3
	l.Push(k2, v2)

	key, val := l.Pop()
	if key != k || val != v {
		t.Errorf("expected (%v, %v), got (%v, %v)", k, v, key, val)
	}
	if l.head.Key != k2 || l.head.Value != v2 {
		t.Errorf("expected (%v, %v), got (%v, %v)", k2, v2, l.head.Key, l.head.Value)
	}
	if l.head.Next != nil || l.head.Prev != nil {
		t.Error("expected nil next & prev for head")
	}
	if l.tail != l.head {
		t.Error("expected tail to be the same as head")
	}
	size := 1
	if l.size != size {
		t.Errorf("expected %d, got %d", size, l.size)
	}
}

func TestLinkedQueue_cut(t *testing.T) {
	l := &queue{}
	k1 := 1
	v1 := 2
	l.Push(k1, v1)
	l.Push(2, 3)
	k2 := 3
	v2 := 4
	l.Push(k2, v2)

	n := l.head.Next
	l.cut(n)

	if l.head.Key != k1 || l.head.Value != v1 {
		t.Errorf("expected (%v, %v), got (%v, %v)", k1, v1, l.head.Key, l.head.Value)
	}
	if l.head.Next.Key != k2 || l.head.Next.Value != v2 {
		t.Errorf("expected (%v, %v), got (%v, %v)", k2, v2, l.head.Next.Key, l.head.Next.Value)
	}
	if l.tail.Prev != l.head {
		t.Error("tail prev has to be equal to head")
	}
	if l.head.Next != l.tail {
		t.Error("head next has to be equal to tail")
	}
	if l.head.Prev != nil {
		t.Error("head prev has to be nil")
	}
	if l.tail.Next != nil {
		t.Error("tail next has to be nil")
	}
	size := 2
	if l.size != size {
		t.Errorf("expected %d, got %d", size, l.size)
	}
}

func TestLru_Put(t *testing.T) {
	l := NewLru(2)
	k1 := 1
	v1 := 2
	k2 := 3
	v2 := 4
	l.Put(k1, v1)
	l.Put(k2, v2)

	size := 2
	if l.queue.size != size {
		t.Errorf("expected %d, got %d", size, l.queue.size)
	}
	if len(l.elements) != size {
		t.Errorf("expected %d, got %d", size, len(l.elements))
	}
	if l.queue.head.Key != k1 || l.queue.head.Value != v1 {
		t.Errorf("expected (%d, %d), got (%d, %d)", k1, v1, l.queue.head.Key, l.queue.head.Value)
	}
	if l.queue.tail.Key != k2 || l.queue.tail.Value != v2 {
		t.Errorf("expected (%d, %d), got (%d, %d)", k2, v2, l.queue.tail.Key, l.queue.tail.Value)
	}
}
func TestLru_Put_With_Evict(t *testing.T) {
	l := NewLru(2)
	k1 := 1
	v1 := 2
	k2 := 3
	v2 := 4
	k3 := 5
	v3 := 6
	l.Put(k1, v1)
	l.Put(k2, v2)
	l.Put(k3, v3)

	size := 2
	if l.queue.size != size {
		t.Errorf("expected %d, got %d", size, l.queue.size)
	}
	if len(l.elements) != size {
		t.Errorf("expected %d, got %d", size, len(l.elements))
	}
	if l.queue.head.Key != k2 || l.queue.head.Value != v2 {
		t.Errorf("expected (%d, %d), got (%d, %d)", k2, v2, l.queue.head.Key, l.queue.head.Value)
	}
	if l.queue.tail.Key != k3 || l.queue.tail.Value != v3 {
		t.Errorf("expected (%d, %d), got (%d, %d)", k3, v3, l.queue.tail.Key, l.queue.tail.Value)
	}
}
func TestLru_Put_With_Evict_And_Reorder(t *testing.T) {
	l := NewLru(2)
	k1 := 1
	v1 := 2
	k2 := 3
	v2 := 4
	k3 := 5
	v3 := 6

	v12 := 100
	l.Put(k1, v1)
	t.Log(l.String())
	l.Put(k2, v2)
	t.Log(l.String())
	l.Put(k1, v12)
	t.Log(l.String())
	l.Put(k3, v3)
	t.Log(l.String())

	size := 2
	if l.queue.size != size {
		t.Errorf("expected %d, got %d", size, l.queue.size)
	}
	if len(l.elements) != size {
		t.Errorf("expected %d, got %d", size, len(l.elements))
	}
	if l.queue.head.Key != k1 || l.queue.head.Value != v12 {
		t.Errorf("expected (%d, %d), got (%d, %d)", k1, v12, l.queue.head.Key, l.queue.head.Value)
	}
	if l.queue.tail.Key != k3 || l.queue.tail.Value != v3 {
		t.Errorf("expected (%d, %d), got (%d, %d)", k3, v3, l.queue.tail.Key, l.queue.tail.Value)
	}
}
