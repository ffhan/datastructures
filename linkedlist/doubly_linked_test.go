package linkedlist

import "testing"

func TestDoublyLinkedList_Append(t *testing.T) {
	l := NewDoublyLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	expected := []int{1, 2, 3}
	for iter := l.Iterator(false); iter.Has(); iter.Next() {
		args := expected[iter.Index()]
		if iter.Get() != args {
			t.Errorf("expected %d, got %d\n", args, iter.Get())
		}
	}
	for iter := l.Iterator(true); iter.Has(); iter.Next() {
		args := expected[iter.Index()]
		if iter.Get() != args {
			t.Errorf("expected %d, got %d\n", args, iter.Get())
		}
	}
	t.Log(l.ToArray())
}

func TestDoublyLinkedList_Prepend(t *testing.T) {
	l := NewDoublyLinkedList()
	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

	expected := []int{3, 2, 1}
	for iter := l.Iterator(false); iter.Has(); iter.Next() {
		args := expected[iter.Index()]
		if iter.Get() != args {
			t.Errorf("expected %d, got %d\n", args, iter.Get())
		}
	}
	for iter := l.Iterator(true); iter.Has(); iter.Next() {
		args := expected[iter.Index()]
		if iter.Get() != args {
			t.Errorf("expected %d, got %d\n", args, iter.Get())
		}
	}
	t.Log(l.ToArray())
}

func TestDoubleIterator_Get(t *testing.T) {
	l := NewDoublyLinkedList()
	l.Append(1)
	expected := 2
	l.Append(expected)
	l.Append(3)

	val, err := l.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if val != expected {
		t.Errorf("expected %d, got %d\n", expected, val)
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	l := NewDoublyLinkedList()
	l.Append(1)
	l.Append(3)
	expected := 2
	if err := l.Insert(1, expected); err != nil {
		t.Fatal(err)
	}
	val, err := l.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if val != expected {
		t.Errorf("expected %d, got %d\n", expected, val)
	}
	t.Log(l.ToArray())
}

func TestDoublyLinkedList_Remove_Head(t *testing.T) {
	l := NewDoublyLinkedList()
	expectedNum := 1
	l.Append(1)
	l.Append(2)
	l.Append(3)

	val, err := l.Remove(0)
	if err != nil {
		t.Fatal(err)
	}
	if val != expectedNum {
		t.Errorf("expected %d, got %d\n", expectedNum, val)
	}
	expected := []int{2, 3}
	for iter := l.Iterator(false); iter.Has(); iter.Next() {
		args := expected[iter.Index()]
		if iter.Get() != args {
			t.Errorf("expected %d, got %d\n", args, iter.Get())
		}
	}
	for iter := l.Iterator(true); iter.Has(); iter.Next() {
		args := expected[iter.Index()]
		if iter.Get() != args {
			t.Errorf("expected %d, got %d\n", args, iter.Get())
		}
	}
}

func TestDoublyLinkedList_Remove_Tail(t *testing.T) {
	l := NewDoublyLinkedList()
	expectedNum := 3
	l.Append(1)
	l.Append(2)
	l.Append(3)

	val, err := l.Remove(2)
	if err != nil {
		t.Fatal(err)
	}
	if val != expectedNum {
		t.Errorf("expected %d, got %d\n", expectedNum, val)
	}
	expected := []int{1, 2}
	for iter := l.Iterator(false); iter.Has(); iter.Next() {
		args := expected[iter.Index()]
		if iter.Get() != args {
			t.Errorf("expected %d, got %d\n", args, iter.Get())
		}
	}
	for iter := l.Iterator(true); iter.Has(); iter.Next() {
		args := expected[iter.Index()]
		if iter.Get() != args {
			t.Errorf("expected %d, got %d\n", args, iter.Get())
		}
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	l := NewDoublyLinkedList()
	expectedNum := 2
	l.Append(1)
	l.Append(2)
	l.Append(3)

	val, err := l.Remove(1)
	if err != nil {
		t.Fatal(err)
	}
	if val != expectedNum {
		t.Errorf("expected %d, got %d\n", expectedNum, val)
	}
	expected := []int{1, 3}
	for iter := l.Iterator(false); iter.Has(); iter.Next() {
		args := expected[iter.Index()]
		if iter.Get() != args {
			t.Errorf("expected %d, got %d\n", args, iter.Get())
		}
	}
	for iter := l.Iterator(true); iter.Has(); iter.Next() {
		args := expected[iter.Index()]
		if iter.Get() != args {
			t.Errorf("expected %d, got %d\n", args, iter.Get())
		}
	}
}
