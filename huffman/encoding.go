package huffman

import (
	"errors"
	"fmt"
	"io"
)

type encoder struct {
	in      io.Reader
	mapping map[byte]code
	eot     code // end of transfer
}

func (e *encoder) Write(p []byte) (n int, err error) {
	if e.in == nil {
		return 0, errors.New("reader is nil")
	}
	buffer := make([]byte, 1)
	currentBitSize := byte(0)
	currentByte := uint64(0)

	var read int
	for n < len(p) {
		if currentBitSize >= 8 {
			p[n] = byte((currentByte >> (currentBitSize - 8)) & 0xFF)
			currentBitSize -= 8
			n += 1
			currentByte &= (1 << currentBitSize) - 1
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
			currentByte = (currentByte << e.eot.bitsize) | uint64(e.eot.value)
			currentBitSize += e.eot.bitsize
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
		currentByte = (currentByte << code.bitsize) | uint64(code.value)
	}
	return n, nil
}
