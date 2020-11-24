package huffman

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestSampleEncodeDecode(t *testing.T) {
	buffer := []byte("abcdefghijklmnoprstuvwxyz encoding is really fun! especially with a lot of special characters like spaces")

	var encodedBuffer bytes.Buffer

	huffman, err := SampleBytes(buffer)
	if err != nil {
		t.Fatal(err)
	}
	origin := "encoding makes things really fun!"
	input := []byte(origin)
	t.Logf("input bytes: %v", input)
	encoder := huffman.Encoder(&encodedBuffer)
	//for key, value := range encoder.mapping {
	//	t.Logf("%d -> %s", key, value.String())
	//}
	//t.Logf("eot -> %s", encoder.eot.String())
	n, err := encoder.Write(input)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			t.Error(err)
		}
	}
	t.Logf("wrote %d bytes: %v with compression rate of %.2f%%", n, encodedBuffer.Bytes(), 100.0*float64(len(input)-n)/float64(len(input)))

	var buf bytes.Buffer
	decodedBytes, err := huffman.Decoder(&buf).Read(encodedBuffer.Bytes())
	if err != nil {
		t.Error(err)
	}
	if decodedBytes != n {
		t.Errorf("expected %d read bytes, got %d\n", n, decodedBytes)
	}
	t.Logf("decoded bytes: %v", buf.Bytes())
	decodedResult := buf.String()
	if decodedResult != origin {
		t.Errorf("expected \"%s\", got \"%s\"", origin, decodedResult)
	}
	t.Logf("decoded result \"%s\"", decodedResult)
}
