package huffman

import (
	"errors"
	"fmt"
	"io"
)

type encoder struct {
	out     io.Writer
	mapping map[byte]code
	eot     code // end of transfer
}

func (e *encoder) write(b byte) error {
	write, err := e.out.Write([]byte{b})
	if err != nil {
		return err
	}
	if write != 1 {
		return errors.New("invalid number of bytes written")
	}
	return nil
}

func (e *encoder) Write(p []byte) (n int, err error) {
	if e.out == nil {
		return 0, errors.New("writer is nil")
	}
	currentBitSize := byte(0)
	currentByte := uint64(0)

	index := 0

	for index < len(p) {
		if currentBitSize < 8 && index < len(p) {
			symbol := p[index]
			index += 1
			code, ok := e.mapping[symbol]
			if !ok {
				return n, fmt.Errorf("symbol %d not found in encoder", symbol)
			}
			currentBitSize += code.bitsize
			currentByte = (currentByte << code.bitsize) | uint64(code.value)
		}
		if currentBitSize >= 8 {
			if err := e.write(byte((currentByte >> (currentBitSize - 8)) & 0xFF)); err != nil {
				return n, err
			}
			currentBitSize -= 8
			n += 1
			currentByte &= (1 << currentBitSize) - 1
			continue
		}
	}
	currentByte = (currentByte << e.eot.bitsize) | uint64(e.eot.value)
	currentBitSize += e.eot.bitsize

	for currentBitSize > 0 {
		if currentBitSize >= 8 {
			if err := e.write(byte((currentByte >> (currentBitSize - 8)) & 0xFF)); err != nil {
				return n, err
			}
			currentBitSize -= 8
			n += 1
			currentByte &= (1 << currentBitSize) - 1
		} else {
			currentByte <<= 8 - currentBitSize // e.g. 3 bits unused - push them to the front of the byte and push the byte to a buffer
			if err := e.write(byte(currentByte)); err != nil {
				return n, err
			}
			n += 1
			currentBitSize = 0
		}
	}
	return n, err
}
