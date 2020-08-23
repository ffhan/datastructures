package lfu

import "testing"

func TestLfu_Put(t *testing.T) {
	size := 5
	l := NewLfu(size)
	for i := 0; i < size; i++ {
		l.Put(i, i)
	}
	if len(l.elements) != size {
		t.Errorf("expected len %d, got %d", size, len(l.elements))
	}
	if len(l.nodes) != size {
		t.Errorf("expected len %d, got %d", size, len(l.elements))
	}
	t.Logf("%+v\n", l)
}

func TestLfu_Put_With_Eviction(t *testing.T) {
	size := 5
	capacity := 3
	l := NewLfu(capacity)
	for i := 0; i < size; i++ {
		l.Put(i, i)
		printElems(t, l)
		for j := 0; j < i; j++ {
			l.Get(i)
			printElems(t, l)
		}
	}
	if len(l.elements) != capacity {
		t.Errorf("expected elements len %d, got %d", capacity, len(l.elements))
	}
	if len(l.nodes) != capacity {
		t.Errorf("expected nodes len %d, got %d", capacity, len(l.nodes))
	}
	printElems(t, l)
}

func TestLfu_Reverse_Put_With_Eviction(t *testing.T) {
	size := 5
	capacity := 3
	l := NewLfu(capacity)
	for i := 0; i < size; i++ {
		l.Put(i, i)
		printElems(t, l)
		for j := 0; j < size-1-i; j++ {
			l.Get(i)
			printElems(t, l)
		}
	}
	if len(l.elements) != capacity {
		t.Errorf("expected elements len %d, got %d", capacity, len(l.elements))
	}
	if len(l.nodes) != capacity {
		t.Errorf("expected nodes len %d, got %d", capacity, len(l.nodes))
	}
	printElems(t, l)
}

func TestLfu_Put_With_Eviction_Without_Get(t *testing.T) {
	size := 5
	capacity := 3
	l := NewLfu(capacity)
	for i := 0; i < size; i++ {
		l.Put(i, i)
		printElems(t, l)
		for j := 0; j < i; j++ {
			l.Put(i, j)
			printElems(t, l)
		}
	}
	if len(l.elements) != capacity {
		t.Errorf("expected elements len %d, got %d", capacity, len(l.elements))
	}
	if len(l.nodes) != capacity {
		t.Errorf("expected nodes len %d, got %d", capacity, len(l.nodes))
	}
	printElems(t, l)
}

func printElems(t *testing.T, l *lfu) {
	for i := 0; i < len(l.nodes); i++ {
		t.Logf("%+v\n", l.nodes[i])
	}
	t.Log("-------------------")
}
