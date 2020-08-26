package stack

import "testing"

func TestStack_Push(t *testing.T) {
	q := NewStack()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	expected := []int{1, 2, 3}

	for i := range expected {
		expVal := expected[i]
		qVal := q.values[i]
		if qVal != expVal {
			t.Errorf("expected %d, got %d\n", expVal, qVal)
		}
	}
	if q.Size() != 3 {
		t.Errorf("expected stack len 3, got %d\n", q.Size())
	}
}

func TestStack_Pop(t *testing.T) {
	q := NewStack()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	expected := []int{3, 2, 1}
	for _, exp := range expected {
		val, err := q.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if val != exp {
			t.Errorf("expected %d, got %d\n", exp, val)
		}
	}
	if q.Size() != 0 || !q.IsEmpty() {
		t.Error("expected empty stack")
	}
}
