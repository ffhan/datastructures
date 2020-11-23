package huffman

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Encoder interface {
	Encode(reader io.Reader) io.Writer
}

type code struct {
	value   int64
	bitsize byte
}

func (c code) String() string {
	return fmt.Sprintf("%0"+strconv.Itoa(int(c.bitsize))+"b", c.value)
}

type node struct {
	left, right *node
	key         byte
	val         int
}

type encoder struct {
	in      io.Reader
	mapping map[byte]code
}

func (e *encoder) SetReader(reader io.Reader) {
	e.in = reader
}

func (e *encoder) Write(p []byte) (n int, err error) {
	if e.in == nil {
		return 0, errors.New("reader is nil")
	}
	buffer := make([]byte, 1)
	currentBitSize := byte(0)
	currentByte := uint16(0)

	var read int
	for n < len(p) {
		if currentBitSize >= 8 {
			currentBitSize -= 8
			p[n] = byte(currentByte & 0xFF)
			n += 1
			currentByte >>= 8
			continue
		}
		if err != nil && currentBitSize > 0 {
			currentByte <<= 8 - currentBitSize // e.g. 3 bits unused - push them to the front of the byte and push the byte to a buffer
			p[n] = byte(currentByte)
			n += 1
			currentBitSize = 0
			continue
		} else if err != nil {
			return n, err
		}
		read, err = e.in.Read(buffer)
		if err != nil {
			continue
		}
		if read <= 0 {
			err = errors.New("read 0 bytes, expected >1 bytes")
			continue
		}
		symbol := buffer[0]
		code, ok := e.mapping[symbol]
		if !ok {
			err = fmt.Errorf("symbol %d not found in encoder", symbol)
			continue
		}
		currentBitSize += code.bitsize
		currentByte = (currentByte << code.bitsize) | uint16(code.value)
	}
	return n, nil
}

func SampleBytes(bytes []byte) (*encoder, error) {
	freqs := getFreqs(bytes)
	tree, err := buildTree(freqs)
	if err != nil {
		return nil, err
	}
	return buildEncoderFromTree(tree), nil
}

func Sample(reader io.Reader, n int) (*encoder, error) {
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

func buildEncoderFromTree(tree *node) *encoder {
	e := &encoder{mapping: make(map[byte]code)}
	e.buildBranch(tree, code{0, 0})
	return e
}

func buildTree(freqs minHeap) (*node, error) {
	for {
		left, ok := freqs.Pop()
		if !ok {
			return nil, errors.New("no elements provided for sampling")
		}
		right, ok := freqs.Pop()
		if !ok {
			return &node{
				key: right.key,
				val: right.val,
			}, nil
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
			return &newNode, nil
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
	return data
}
