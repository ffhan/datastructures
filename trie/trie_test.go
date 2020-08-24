package trie

import "testing"

func TestTrie_Put(t *testing.T) {
	tr := NewTrie()
	expected := 12
	tr.Put("test", expected)
	tr.Put("tent", expected)
	tr.Put("testing", expected)
	value := tr.root.children['t'].children['e'].children['s'].children['t'].value
	value2 := tr.root.children['t'].children['e'].children['n'].children['t'].value
	value3 := tr.root.children['t'].children['e'].children['s'].children['t'].children['i'].children['n'].children['g'].value
	if value != expected {
		t.Errorf("expected %d, got %d", expected, value)
	}
	if value2 != expected {
		t.Errorf("expected %d, got %d", expected, value2)
	}
	if value3 != expected {
		t.Errorf("expected %d, got %d", expected, value3)
	}
}

func TestSliceMechanisms(t *testing.T) {
	m := make([]int, 0, 4)
	m = append(m, 1)
	t.Log(m, len(m), cap(m))
	m = append(m, 2)
	t.Log(m, len(m), cap(m))
	m = append(m, 3)
	t.Log(m, len(m), cap(m))
	m = append(m, 4)
	t.Log(m, len(m), cap(m))
	m = append(m, 5)
	t.Log(m, len(m), cap(m))
	m = m[:2]
	t.Log(m, len(m), cap(m))
	m = append(m, 100)
	t.Log(m, len(m), cap(m))
	m = m[:cap(m)]
	t.Log(m, len(m), cap(m))
}

func checkElements(expected, results []string) bool {
	if len(expected) != len(results) {
		return false
	}
	db := make(map[string]bool)
	for _, s := range expected {
		db[s] = true
	}
	for _, s := range results {
		if _, ok := db[s]; !ok {
			return false
		}
	}
	return true
}

func TestAutocomplete(t *testing.T) {
	tr := NewTrie()
	s1 := "testing"
	tr.Put(s1, 1)
	s2 := "test"
	tr.Put(s2, 2)
	s3 := "tent"
	tr.Put(s3, 3)
	s4 := "tenant"
	tr.Put(s4, 4)

	results := tr.Autocomplete("tes")
	expected := []string{s1, s2}
	if !checkElements(expected, results) {
		t.Errorf("expected %v, got %v\n", expected, results)
	}
	results = tr.Autocomplete("tena")
	expected = []string{s4}
	if !checkElements(expected, results) {
		t.Errorf("expected %v, got %v\n", expected, results)
	}
}
