package huffman

type minHeap []node

func makeMinHeap() minHeap {
	return make(minHeap, 0, 16)
}

// 0->(1,2), 1->(3,4)
func getNodeParent(index int) int {
	return (index - 1) / 2
}

func getChildren(index int) (int, int) {
	left := 2*index + 1
	return left, left + 1
}

func (m *minHeap) Size() int {
	return len(*m)
}

func (m *minHeap) Push(node node) {
	lastIndex := len(*m)
	*m = append(*m, node)
	m.fixUpstream(lastIndex)
}

func (m *minHeap) Pop() (*node, bool) {
	arr := *m
	if len(arr) == 0 {
		return nil, false
	}
	result := arr[0]
	lastIndex := len(arr) - 1
	arr[0] = arr[lastIndex]
	*m = arr[:lastIndex]

	m.fixDownstream(0)
	return &result, true
}

func (m *minHeap) fixDownstream(startIndex int) {
	arr := *m
	N := len(arr)
	if startIndex < 0 || startIndex >= N {
		return
	}
	current := arr[startIndex]
	left, right := getChildren(startIndex)
	isLeftValid := left < N
	isRightValid := right < N
	if isLeftValid && isRightValid && arr[right].val < arr[left].val {
		temp := right
		right = left
		left = temp
	}
	if isLeftValid && arr[left].val < current.val {
		arr[startIndex] = arr[left]
		arr[left] = current
		m.fixDownstream(left)
	} else if isRightValid && arr[right].val < current.val {
		arr[startIndex] = arr[right]
		arr[right] = current
		m.fixDownstream(right)
	}
}

func (m *minHeap) fixUpstream(startIndex int) {
	arr := *m
	if startIndex <= 0 {
		return
	}
	parentIndex := getNodeParent(startIndex)
	parent := arr[parentIndex]
	if parent.val > arr[startIndex].val {
		arr[parentIndex] = arr[startIndex]
		arr[startIndex] = parent
		m.fixUpstream(parentIndex)
	}
}
