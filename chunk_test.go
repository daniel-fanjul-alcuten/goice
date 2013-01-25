package ice

import (
	"testing"
)

func TestChunkSize(t *testing.T) {
	if ChunkSize != 65535 {
		t.Error("ChunkSize", ChunkSize)
	}
}

func TestCopy(t *testing.T) {
	str, chunk := "foo\n", &Chunk{}
	chunk.Copy([]byte(str))
	if chunk.Len != len(str) {
		t.Error("Unexpected chunk.Len", chunk.Len)
	}
	if string(chunk.Dat[:len(str)]) != str {
		t.Error("Unexpected chunk.Dat", chunk.Dat[:len(str)])
	}
}

func TestData(t *testing.T) {
	str, chunk := "foo\n", &Chunk{}
	chunk.Copy([]byte(str))
	if string(chunk.Data()) != str {
		t.Error("Unexpected chunk.Data()", chunk.Data())
	}
}
