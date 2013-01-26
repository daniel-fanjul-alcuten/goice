package ice

import (
	"io"
)

var BERG_MAGIC = []byte("ice1b")

type BergWriter struct {
	writer io.Writer
	hasher *HashWriter
	shas   map[Sha]bool
}

func NewBergWriter(writer io.Writer, size int) (*BergWriter, error) {
	if _, err := writer.Write(BERG_MAGIC); err != nil {
		return nil, err
	}
	hasher := NewHashWriter(writer, size)
	shas := make(map[Sha]bool, 1024)
	return &BergWriter{writer, hasher, shas}, nil
}

func (w *BergWriter) Write(shaChunk *ShaChunk) error {
	if !w.shas[*shaChunk.Sha] {
		if _, err := w.hasher.Write(shaChunk.Sha.Data()); err != nil {
			return err
		}
		if _, err := w.hasher.Write(shaChunk.Chunk.LenData()); err != nil {
			return err
		}
		if _, err := w.hasher.Write(shaChunk.Chunk.Data()); err != nil {
			return err
		}
		w.shas[*shaChunk.Sha] = true
	}
	return nil
}

func (w *BergWriter) Close() error {
	return w.hasher.Close()
}

func (w *BergWriter) Sha() (*Sha, error) {
	sha := w.hasher.Sha()
	_, err := w.writer.Write(sha.Data())
	return sha, err
}
