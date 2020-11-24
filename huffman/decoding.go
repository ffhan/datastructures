package huffman

import "io"

type decoder struct {
	in   io.Reader
	tree tree
}

func (d *decoder) read(buffer []byte) (byte, error) {
	_, err := d.in.Read(buffer)
	if err != nil {
		return 0, err
	}
	return buffer[0], nil
}

func (d *decoder) Read(p []byte) (n int, err error) {
	inBuffer := make([]byte, 1)

	bitOffset := 0
	currentByte, err := d.read(inBuffer)
	if err != nil {
		return n, err
	}
	node := &d.tree.root

	eot := false

readLoop:
	for n <= len(p) && !eot {
		node = &d.tree.root
		for node.left != nil || node.right != nil {
			if bitOffset >= 8 {
				if n >= len(p) {
					break readLoop
				}
				currentByte, err = d.read(inBuffer)
				if err != nil {
					return n, err
				}
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

		if !node.isEOT {
			p[n] = node.key
			n += 1
		} else {
			eot = true
		}
	}
	return n, nil
}
