package ice

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type BergReader struct {
	reader io.Reader
}

func NewBergReader(reader io.Reader) (*BergReader, error) {
	magic := make([]byte, len(BERG_MAGIC))
	if _, err := io.ReadFull(reader, magic); err != nil {
		return nil, err
	}
	if !bytes.Equal(magic, BERG_MAGIC) {
		return nil, fmt.Errorf("invalid magic %v", magic)
	}
	return &BergReader{reader}, nil
}

func (r *BergReader) Read() (*ShaChunk, *Sha, error) {
	sha := &Sha{}
	if _, err := io.ReadFull(r.reader, sha.Data()); err != nil {
		return nil, nil, err
	}
	lenData := make([]byte, 2)
	if _, err := io.ReadFull(r.reader, lenData); err != nil {
		if err == io.EOF {
			return nil, sha, nil
		}
		return nil, nil, err
	}
	chunk := &Chunk{Len: int(binary.BigEndian.Uint16(lenData))}
	if _, err := io.ReadFull(r.reader, chunk.Data()); err != nil {
		return nil, nil, err
	}
	return &ShaChunk{sha, chunk}, nil, nil
}
