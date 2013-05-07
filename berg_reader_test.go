package ice

import (
	"bytes"
	"testing"
)

func TestBergReader(t *testing.T) {

	buffer := &bytes.Buffer{}
	writer, err := NewBergWriter(buffer, 1)
	if err != nil {
		t.Fatal("Unexpected err", err)
	}

	str, chunk := "foo\n", &Chunk{}
	chunk.Copy([]byte(str))
	err = writer.Write(chunk.Hash())
	if err != nil {
		t.Fatal("Unexpected err", err)
	}

	err = writer.Close()
	if err != nil {
		t.Fatal("Unexpected err", err)
	}

	_, err = writer.Sha()
	if err != nil {
		t.Fatal("Unexpected err", err)
	}

	reader, err := NewBergReader(buffer)
	if err != nil {
		t.Fatal("Unexpected err", err)
	}

	shaChunk, sha, err := reader.Read()
	if err != nil {
		t.Fatal("Unexpected err", err)
	}
	if sha != nil {
		t.Error("Unexpected sha", sha)
	}
	if shaChunk.Sha == nil {
		t.Error("Unexpected shaChunk.Sha", shaChunk.Sha)
	} else if shaChunk.Sha.String() != "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c" {
		t.Error("Unexpected shaChunk.Sha", shaChunk.Sha.String())
	}
	if shaChunk.Chunk == nil {
		t.Error("Unexpected shaChunk.Chunk", shaChunk.Chunk)
	} else {
		if shaChunk.Chunk.Len != 4 {
			t.Error("Unexpected shaChunk.Chunk.Len", shaChunk.Chunk.Len)
		}
		if string(shaChunk.Chunk.Data()) != str {
			t.Error("Unexpected shaChunk.Chunk.Data", shaChunk.Chunk.Data())
		}
	}

	shaChunk, sha, err = reader.Read()
	if err != nil {
		t.Fatal("Unexpected err", err)
	}
	if sha == nil {
		t.Error("Unexpected sha", sha)
	} else if sha.String() != "2d9416ef7150ace93dbf69cb813ef0b071337c2cd97befca68bfc22d5434bb9b" {
		t.Error("Unexpected sha", sha.String())
	}
	if shaChunk != nil {
		t.Error("Unexpected shaChunk", shaChunk)
	}
}
