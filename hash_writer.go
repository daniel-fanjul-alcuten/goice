package ice

import (
	"crypto/sha256"
	"hash"
	"io"
)

type HashWriter struct {
	Writer io.Writer
	hash   hash.Hash
	queue  chan []byte
	wait   chan bool
}

func NewHashWriter(writer io.Writer, size int) *HashWriter {
	hash, queue, wait := sha256.New(), make(chan []byte, size), make(chan bool)
	go func() {
		for data := range queue {
			hash.Write(data)
		}
		close(wait)
	}()
	return &HashWriter{writer, hash, queue, wait}
}

func (w *HashWriter) Write(data []byte) (int, error) {
	n, err := w.Writer.Write(data)
	if n > 0 {
		w.queue <- data[:n]
	}
	return n, err
}

func (w *HashWriter) Close() error {
	close(w.queue)
	return nil
}

func (w *HashWriter) Sha() *Sha {
	<-w.wait
	sha := &Sha{}
	w.hash.Sum(sha[:0])
	return sha
}
