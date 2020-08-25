package linkedlist

import (
	"errors"
	"testing"
)

func TestLinkedList_IsCircular_CircularList(t *testing.T) {
	n4 := &node{value: 4}
	n3 := &node{next: n4, value: 3}
	n2 := &node{next: n3, value: 2}
	n4.next = n2
	n1 := &node{next: n2, value: 1}
	l := &linkedList{head: n1, tail: n4, size: 4}

	if !l.IsCircularAndCheckSizeMatches() {
		t.Errorf("didn't detect a circular list")
	}
}

func TestLinkedList_IsCircular_InvalidSize(t *testing.T) {
	n4 := &node{value: 4}
	n3 := &node{next: n4, value: 3}
	n2 := &node{next: n3, value: 2}
	n1 := &node{next: n2, value: 1}
	l := &linkedList{head: n1, tail: n4, size: 5}

	if !l.IsCircularAndCheckSizeMatches() {
		t.Errorf("didn't detect invalid size")
	}
}

func TestLinkedList_Append(t *testing.T) {
	l := NewLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	expected := []int{1, 2, 3}

	for iter := l.Iterator(); iter.HasNext(); iter.Next() {
		get := iter.Get().(int)
		args := expected[iter.Index()]
		if get != args {
			t.Errorf("expected %d, got %d\n", args, get)
		}
	}
	t.Log(l.ToArray())
}

func TestLinkedList_Prepend(t *testing.T) {
	l := NewLinkedList()
	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

	expected := []int{3, 2, 1}

	for iter := l.Iterator(); iter.HasNext(); iter.Next() {
		get := iter.Get().(int)
		args := expected[iter.Index()]
		if get != args {
			t.Errorf("expected %d, got %d\n", args, get)
		}
	}
	t.Log(l.ToArray())
}

func TestLinkedList_Insert(t *testing.T) {
	l := NewLinkedList()
	l.Append(1)
	l.Append(3)
	l.Append(4)

	err := l.Insert(1, 2)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int{1, 2, 3, 4}

	for iter := l.Iterator(); iter.HasNext(); iter.Next() {
		get := iter.Get().(int)
		args := expected[iter.Index()]
		if get != args {
			t.Errorf("expected %d, got %d\n", args, get)
		}
	}
	t.Log(l.ToArray())
}

func TestLinkedList_Insert_InvalidIndex(t *testing.T) {
	l := NewLinkedList()
	err := l.Insert(-2, 1)
	if !errors.Is(err, IndexOutOfRangeErr) {
		t.Errorf("expected error %v to be of %v\n", err, IndexOutOfRangeErr)
	}
	err = l.Insert(2, 1)
	if !errors.Is(err, IndexOutOfRangeErr) {
		t.Errorf("expected error %v to be of %v\n", err, IndexOutOfRangeErr)
	}
}

func TestLinkedList_Remove(t *testing.T) {
	l := NewLinkedList()
	_, err := l.Remove(-2)
	if !errors.Is(err, IndexOutOfRangeErr) {
		t.Errorf("expected error %v to be of %v\n", err, IndexOutOfRangeErr)
	}
	_, err = l.Remove(2)
	if !errors.Is(err, IndexOutOfRangeErr) {
		t.Errorf("expected error %v to be of %v\n", err, IndexOutOfRangeErr)
	}

	l.Append(1)
	_, err = l.Remove(0)
	if err != nil {
		t.Errorf("expected no errors, got %v\n", err)
	}
	if l.IsCorrupted() {
		t.Error("corrupted list")
	}
	l.Append(1)
	expectedVal := 2
	l.Append(expectedVal)
	l.Append(3)

	val, err := l.Remove(1)
	if err != nil {
		t.Errorf("expected no errors, got %v\n", err)
	} else if val != expectedVal {
		t.Errorf("expected removed val %d, got %d\n", expectedVal, val)
	}
	expected := []int{1, 3}
	for iter := l.Iterator(); iter.HasNext(); iter.Next() {
		get := iter.Get().(int)
		args := expected[iter.Index()]
		if get != args {
			t.Errorf("expected %d, got %d\n", args, get)
		}
	}
	t.Log(l.ToArray())
}
