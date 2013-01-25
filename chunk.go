package ice

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
