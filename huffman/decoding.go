package huffman

import "io"

type decoder struct {
	out  io.Writer
	tree tree
}

func (d *decoder) Read(p []byte) (n int, err error) {
	const bufferMaxSize = 16

	bitOffset := 0
	currentByte := p[n]
	n += 1
	node := &d.tree.root

	buffer := make([]byte, 0, bufferMaxSize)

	eot := false

readLoop:
	for n <= len(p) && !eot {
		node = &d.tree.root
		for node.left != nil || node.right != nil {
			if bitOffset >= 8 {
				if n >= len(p) {
					break readLoop
				}
				currentByte = p[n]
				n += 1
				bitOffset = 0
			}
			bit := (currentByte >> 7) & 1
			currentByte <<= 1
			bitOffset += 1

			if bit == 1 {
				node = node.left
			} else {
				node = node.right
			}
		}

		if len(buffer) == bufferMaxSize {
			if _, err := d.out.Write(buffer); err != nil {
				return n, err
			}
			buffer = buffer[:0] // clear the buffer slice
		}
		if !node.isEOT {
			buffer = append(buffer, node.key)
		} else {
			eot = true
		}
	}
	if _, err = d.out.Write(buffer); err != nil {
		return n, err
	}
	return n, nil
}
