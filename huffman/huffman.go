package huffman

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type code struct {
	value   int64
	bitsize byte
	eot     bool
}

func (c code) String() string {
	return fmt.Sprintf("%0"+strconv.Itoa(int(c.bitsize))+"b", c.value)
}

type node struct {
	left, right *node
	key         byte
	isEOT       bool
	val         int
}

type tree struct {
	root node
	eot  *node
}

type huffman struct {
	tree tree
}

func (h huffman) Encoder(writer io.Writer) *encoder { // todo: implement with a Writer
	e := buildEncoderFromTree(&h.tree)
	e.out = writer
	return e
}

func (h huffman) Decoder(writer io.Writer) *decoder { // todo: implement with a Reader
	return &decoder{out: writer, tree: h.tree}
}

func SampleBytes(bytes []byte) (*huffman, error) {
	freqs := getFreqs(bytes)
	tree, err := buildTree(freqs)
	if err != nil {
		return nil, err
	}
	h := huffman{tree: *tree}
	return &h, nil
}

func Sample(reader io.Reader, n int) (*huffman, error) {
	bytes := make([]byte, n)
	readBytes, err := reader.Read(bytes)
	if err != nil {
		return nil, err
	}

	if readBytes < n { // if less bytes are available than n, sample and encode only available bytes.
		n = readBytes
		bytes = bytes[:readBytes]
	}
	return SampleBytes(bytes)
}

func (e *encoder) buildBranch(tree *node, val code) {
	if tree.left == nil && tree.right == nil {
		if tree.isEOT {
			e.eot = val
		}
		e.mapping[tree.key] = val
		return
	}
	if tree.left != nil {
		leftCode := val
		leftCode.value = (leftCode.value << 1) | 1
		leftCode.bitsize += 1
		e.buildBranch(tree.left, leftCode)
	}
	if tree.right != nil {
		rightCode := val
		rightCode.value = rightCode.value << 1
		rightCode.bitsize += 1
		e.buildBranch(tree.right, rightCode)
	}
}

func buildEncoderFromTree(t *tree) *encoder {
	e := &encoder{mapping: make(map[byte]code)}
	e.eot.eot = true
	e.buildBranch(&t.root, code{0, 0, false})
	return e
}

func buildTree(freqs minHeap) (*tree, error) {
	t := &tree{
		root: node{},
		eot:  nil,
	}
	for {
		left, ok := freqs.Pop()
		if !ok {
			return nil, errors.New("no elements provided for sampling")
		}
		if left.isEOT {
			t.eot = left
		}
		right, ok := freqs.Pop()
		if !ok {
			t.root = node{
				key: left.key,
				val: left.val,
			}
			return t, nil
		}
		if right.isEOT {
			t.eot = right
		}
		if right.val > left.val { // ensure right is bigger than left
			temp := right
			right = left
			left = temp
		}
		newNode := node{
			left:  left,
			right: right,
			val:   left.val + right.val,
		}
		if freqs.Size() == 0 {
			t.root = newNode
			return t, nil
		}
		freqs.Push(newNode)
	}
}

func getFreqs(bytes []byte) minHeap {
	n := len(bytes)
	freqs := make(map[byte]int)
	for i := 0; i < n; i++ {
		freqs[bytes[i]] += 1
	}

	data := make(minHeap, 0, len(freqs))
	for key, val := range freqs {
		data.Push(node{
			key: key,
			val: val,
		})
	}
	data.Push(node{key: 0, isEOT: true, val: 1})
	return data
}
