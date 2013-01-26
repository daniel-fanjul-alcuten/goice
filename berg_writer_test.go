package ice

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestBergWriter(t *testing.T) {

	buffer := &bytes.Buffer{}
	writer, err := NewBergWriter(buffer, 1)
	if writer == nil {
		t.Fatal("Unexpected writer", writer)
	}
	if err != nil {
		t.Error("Unexpected err", err)
	}

	str, chunk := "foo\n", &Chunk{}
	chunk.Copy([]byte(str))
	err = writer.Write(chunk.Hash())
	if err != nil {
		t.Error("Unexpected err", err)
	}

	err = writer.Close()
	if err != nil {
		t.Error("Unexpected err", err)
	}

	sha, err := writer.Sha()
	if sha == nil {
		t.Error("Unexpected sha", sha)
	} else if sha.String() != "2d9416ef7150ace93dbf69cb813ef0b071337c2cd97befca68bfc22d5434bb9b" {
		t.Error("Unexpected sha", sha.String())
	}
	if err != nil {
		t.Error("Unexpected err", err)
	}

	magic := buffer.Next(len(BERG_MAGIC))
	if !bytes.Equal(magic, BERG_MAGIC) {
		t.Error("Unexpected magic", magic)
	}

	sha1 := hex.EncodeToString(buffer.Next(ShaSize))
	if sha1 != "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c" {
		t.Error("Unexpected sha1", sha1)
	}

	length := buffer.Next(2)
	if !bytes.Equal(length, []byte{0, byte(len(str))}) {
		t.Error("Unexpected length", length)
	}

	data := buffer.Next(len(str))
	if string(data) != str {
		t.Error("Unexpected data", data)
	}

	sha2 := hex.EncodeToString(buffer.Next(ShaSize))
	if sha2 != "2d9416ef7150ace93dbf69cb813ef0b071337c2cd97befca68bfc22d5434bb9b" {
		t.Error("Unexpected sha2", sha2)
	}

	if buffer.Len() > 0 {
		t.Error("Unexpected buffer", buffer)
	}
}

func TestBergWriterDuplicatedChunks(t *testing.T) {

	buffer := &bytes.Buffer{}
	writer, err := NewBergWriter(buffer, 1)
	if writer == nil {
		t.Fatal("Unexpected writer", writer)
	}
	if err != nil {
		t.Error("Unexpected err", err)
	}

	chunk := &Chunk{ChunkSize, [ChunkSize]byte{'f', 'o', 'o', '\n'}}
	err = writer.Write(chunk.Hash())
	if err != nil {
		t.Error("Unexpected err", err)
	}
	err = writer.Write(chunk.Hash())
	if err != nil {
		t.Error("Unexpected err", err)
	}

	err = writer.Close()
	if err != nil {
		t.Error("Unexpected err", err)
	}

	sha, err := writer.Sha()
	if sha == nil {
		t.Error("Unexpected sha", sha)
	} else if sha.String() != "d5f449ffe19a2a6f2573c9ed1fa03ab32c1ba44b1f4e4c1e8c241e22c2a6ef46" {
		t.Error("Unexpected sha", sha.String())
	}
	if err != nil {
		t.Error("Unexpected err", err)
	}

	magic := buffer.Next(len(BERG_MAGIC))
	if !bytes.Equal(magic, BERG_MAGIC) {
		t.Error("Unexpected magic", magic)
	}

	sha1 := hex.EncodeToString(buffer.Next(ShaSize))
	if sha1 != "090679444268a26337721a6f3819395feeedde88238a05b6d6d7eab26ae755a9" {
		t.Error("Unexpected sha1", sha1)
	}

	length := buffer.Next(2)
	if !bytes.Equal(length, []byte{255, 255}) {
		t.Error("Unexpected length", length)
	}

	data := buffer.Next(ChunkSize)
	if !bytes.Equal(data, chunk.Dat[:]) {
		t.Error("Unexpected data", data)
	}

	sha2 := hex.EncodeToString(buffer.Next(ShaSize))
	if sha2 != "d5f449ffe19a2a6f2573c9ed1fa03ab32c1ba44b1f4e4c1e8c241e22c2a6ef46" {
		t.Error("Unexpected sha2", sha2)
	}

	if buffer.Len() > 0 {
		t.Error("Unexpected buffer data", buffer.Len())
	}
}
