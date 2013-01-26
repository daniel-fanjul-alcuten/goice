package ice

import (
	"crypto/sha256"
	"encoding/binary"
	"io"
)

const ChunkSize = 2<<15 - 1

type Chunk struct {
	Len int
	Dat [ChunkSize]byte
}

func (c *Chunk) Copy(data []byte) {
	c.Len = copy(c.Dat[:], data)
}

func (c *Chunk) ReadFrom(reader io.Reader) error {
	n, err := io.ReadFull(reader, c.Dat[c.Len:])
	c.Len += n
	if err == io.ErrUnexpectedEOF {
		err = io.EOF
	}
	return err
}

func (c *Chunk) LenData() []byte {
	length := make([]byte, 2)
	binary.BigEndian.PutUint16(length, uint16(c.Len))
	return length
}

func (c *Chunk) Data() []byte {
	return c.Dat[:c.Len]
}

func (c *Chunk) Hash() *ShaChunk {
	sha, h := &Sha{}, sha256.New()
	h.Write(c.Data())
	h.Sum(sha[:0])
	return &ShaChunk{sha, c}
}
