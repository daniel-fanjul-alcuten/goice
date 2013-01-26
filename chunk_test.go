package ice

import (
	"bytes"
	"io"
	"testing"
)

func TestChunkSize(t *testing.T) {
	if ChunkSize != 65535 {
		t.Error("ChunkSize", ChunkSize)
	}
}

func TestChunkCopy(t *testing.T) {
	str, chunk := "foo\n", &Chunk{}
	chunk.Copy([]byte(str))
	if chunk.Len != len(str) {
		t.Error("Unexpected chunk.Len", chunk.Len)
	}
	if string(chunk.Dat[:len(str)]) != str {
		t.Error("Unexpected chunk.Dat", chunk.Dat[:len(str)])
	}
}

func TestChunkReadFrom(t *testing.T) {
	str, buffer, chunk := "foo\n", &bytes.Buffer{}, &Chunk{}
	buffer.Write([]byte(str))
	err := chunk.ReadFrom(buffer)
	if err != io.EOF {
		t.Error("Unexpected err", err)
	}
	if string(chunk.Data()) != str {
		t.Error("Unexpected chunk.Data()", chunk.Data())
	}
}

func TestChunkLenData(t *testing.T) {
	str, chunk := "foo\n", &Chunk{}
	chunk.Copy([]byte(str))
	if !bytes.Equal(chunk.LenData(), []byte{0, byte(len(str))}) {
		t.Error("Unexpected chunk.LenData()", chunk.LenData())
	}
}

func TestChunkData(t *testing.T) {
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
