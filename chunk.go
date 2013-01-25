package ice

import (
	"crypto/sha256"
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

func (c *Chunk) ReadFrom(reader io.Reader) (err error) {
	c.Len, err = io.ReadFull(reader, c.Dat[:])
	if err == io.ErrUnexpectedEOF {
		err = io.EOF
	}
	return
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
