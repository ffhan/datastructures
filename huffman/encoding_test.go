package huffman

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestSample(t *testing.T) {
	buffer := []byte("abcdefghijklmnoprstuvwxyz encoding is really fun! especially with a lot of special characters like spaces")

	encoder, err := SampleBytes(buffer)
	if err != nil {
		t.Fatal(err)
	}
	input := []byte("encoding makes things really fun!")
	encoder.SetReader(bytes.NewBuffer(input))
	for key, value := range encoder.mapping {
		t.Logf("%d -> %s", key, value.String())
	}
	result := make([]byte, 1024)
	n, err := encoder.Write(result)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			t.Error(err)
		}
	}
	t.Logf("wrote %d bytes: %v with compression rate of %.2f%%", n, result[:n], 100.0*float64(len(input)-n)/float64(len(input)))
}
