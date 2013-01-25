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

func TestChunkHash(t *testing.T) {
	chunk := &Chunk{}
	chunk.Copy([]byte("foo\n"))
	shaChunk := chunk.Hash()
	if shaChunk == nil {
		t.Error("Unexpected shaChunk", shaChunk)
	} else {
		if shaChunk.Sha == nil {
			t.Error("Unexpected shaChunk.Sha", shaChunk.Sha)
		} else {
			str := shaChunk.Sha.String()
			if str != "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c" {
				t.Error("Unexpected String()", str)
			}
		}
		if shaChunk.Chunk != chunk {
			t.Error("Unexpected shaChunk.Chunk", shaChunk.Chunk)
		}
	}
}
