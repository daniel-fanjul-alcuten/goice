package ice

import (
	"io"
)

var BERG_MAGIC = []byte("ice1b")

type BergWriter struct {
	writer io.Writer
	hasher *HashWriter
}

func NewBergWriter(writer io.Writer, size int) (*BergWriter, error) {
	if _, err := writer.Write(BERG_MAGIC); err != nil {
		return nil, err
	}
	hasher := NewHashWriter(writer, size)
	return &BergWriter{writer, hasher}, nil
}

func (w *BergWriter) Write(shaChunk *ShaChunk) error {
	if _, err := w.hasher.Write(shaChunk.Sha.Data()); err != nil {
		return err
	}
	if _, err := w.hasher.Write(shaChunk.Chunk.LenData()); err != nil {
		return err
	}
	if _, err := w.hasher.Write(shaChunk.Chunk.Data()); err != nil {
		return err
	}
	return nil
}

func (w *BergWriter) Close() (*Sha, error) {
	w.hasher.Close()
	sha := w.hasher.Sha()
	_, err := w.writer.Write(sha.Data())
	return sha, err
}
