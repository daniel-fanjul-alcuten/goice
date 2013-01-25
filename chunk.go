package ice

import (
	"crypto/sha256"
)

const ChunkSize = 2<<15 - 1

type Chunk struct {
	Len int
	Dat [ChunkSize]byte
}

func (c *Chunk) Copy(data []byte) {
	c.Len = copy(c.Dat[:], data)
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
