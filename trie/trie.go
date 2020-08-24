package trie

type node struct {
	children map[rune]*node
	value    interface{}
}

func newNode() *node {
	return &node{children: make(map[rune]*node)}
}

type trie struct {
	root *node
}

func NewTrie() *trie {
	return &trie{}
}

func (t *trie) find(key string) *node {
	node := t.root
	var ok bool
	for _, r := range key {
		node, ok = node.children[r]
		if !ok {
			return nil
		}
	}
	return node
}

func (t *trie) Put(key string, value interface{}) {
	if t.root == nil {
		t.root = newNode()
	}
	n := t.root
	for _, r := range key {
		if child, ok := n.children[r]; ok {
			n = child
		} else {
			n2 := newNode()
			n.children[r] = n2
			n = n2
		}
	}
	n.value = value
}

func autocomplete(n *node, prefix []rune, results []string) []string {
	if n == nil {
		return results
	}
	if n.value != nil {
		results = append(results, string(prefix))
	}
	for key, child := range n.children {
		prefix = append(prefix, key)
		results = autocomplete(child, prefix, results)
		prefix = prefix[:len(prefix)-1]
	}
	return results
}

func (t *trie) Autocomplete(prefix string) []string {
	node := t.find(prefix)
	if node == nil {
		return nil
	}
	results := make([]string, 0)
	return autocomplete(node, []rune(prefix), results)
}
