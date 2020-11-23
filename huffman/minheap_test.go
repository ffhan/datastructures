package huffman

import "testing"

func TestMinHeap_Push(t *testing.T) {
	heap := makeMinHeap()

	heap.Push(node{key: 5, val: 5})
	heap.Push(node{key: 3, val: 3})
	heap.Push(node{key: 1, val: 1})
	heap.Push(node{key: 2, val: 2})
	heap.Push(node{key: 4, val: 4})

	expected := []int{1, 2, 3, 5, 4}

	for i, expect := range expected {
		if heap[i].val != expect {
			t.Errorf("expected %d, got %d\n", expect, heap[i].key)
		}
	}
}

func TestMinHeap_Pop(t *testing.T) {
	heap := makeMinHeap()

	heap.Push(node{key: 5, val: 5})
	heap.Push(node{key: 3, val: 3})
	heap.Push(node{key: 1, val: 1})
	heap.Push(node{key: 2, val: 2})
	heap.Push(node{key: 4, val: 4})

	expected := [][]int{
		{1, 2, 3, 5, 4},
		{2, 4, 3, 5},
		{3, 4, 5},
		{4, 5},
		{5},
		{},
	}

	var pop *node
	var ok bool
	for i := 1; len(heap) > 0; i++ {
		pop, ok = heap.Pop()

		for j, elem := range expected[i] {
			if heap[j].val != elem {
				t.Errorf("expected %d, got %d\n", elem, heap[j].val)
			}
		}

		if len(expected[i-1]) > 0 {
			if !ok {
				t.Errorf("expected ok, got not ok")
			}
			if expected[i-1][0] != pop.val {
				t.Errorf("expected popped value to be %d, got %d\n", expected[i][0], pop.val)
			}
		}
	}
	for i := 0; i < 3; i++ {
		pop, ok = heap.Pop()
		if ok {
			t.Errorf("expected not ok, got ok")
		}
		if len(heap) != 0 {
			t.Errorf("expected len 0, got %d\n", len(heap))
		}
	}
}
