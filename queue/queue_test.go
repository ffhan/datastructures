package queue

import "testing"

func TestQueue_Push(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	expected := []int{1, 2, 3}
	for _, val := range expected {
		result, err := q.list.Remove(0)
		if err != nil {
			t.Fatal(err)
		}
		if result != val {
			t.Errorf("expected %d, got %d\n", val, result)
		}
	}
}

func TestQueue_Pop(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	expected := []int{1, 2, 3}
	for _, expect := range expected {
		val, err := q.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if expect != val {
			t.Errorf("expected %d, got %d\n", expect, val)
		}
	}
	if q.Size() != 0 || !q.IsEmpty() {
		t.Error("expected empty queue")
	}
}
