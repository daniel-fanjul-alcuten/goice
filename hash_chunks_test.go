package ice

import (
	"testing"
)

func TestHashChunksSorted(t *testing.T) {
	chunk1, chunk2 := &Chunk{}, &Chunk{}
	input, output := make(chan *Chunk, 2), make(chan *ShaChunk, 2)
	input <- chunk1
	input <- chunk2
	close(input)
	HashChunksSorted(input, output, 2)
	assertSorted(t, chunk1, chunk2, output)
}

func TestHashChunksUnsorted(t *testing.T) {
	chunk1, chunk2 := &Chunk{}, &Chunk{}
	input, output := make(chan *Chunk, 2), make(chan *ShaChunk, 2)
	input <- chunk1
	input <- chunk2
	close(input)
	HashChunksUnsorted(input, output, 2)
	assertUnsorted(t, chunk1, chunk2, output)
}

func TestHashChunks(t *testing.T) {
	chunk1, chunk2 := &Chunk{}, &Chunk{}
	input, sorted, unsorted := make(chan *Chunk, 2), make(chan *ShaChunk, 2), make(chan *ShaChunk, 2)
	input <- chunk1
	input <- chunk2
	close(input)
	HashChunks(input, sorted, unsorted, 2)
	assertSorted(t, chunk1, chunk2, sorted)
	assertUnsorted(t, chunk1, chunk2, unsorted)
}

func assertSorted(t *testing.T, chunk1, chunk2 *Chunk, output <-chan *ShaChunk) {
	shaChunk1, ok := <-output
	if !ok {
		t.Error("Unexpected ok", ok)
	}
	shaChunk2, ok := <-output
	if !ok {
		t.Error("Unexpected ok", ok)
	}
	_, ok = <-output
	if ok {
		t.Error("Unexpected ok", ok)
	}
	if shaChunk1 == nil {
		t.Error("Unexpected shaChunk1", shaChunk1)
	} else if shaChunk2 == nil {
		t.Error("Unexpected shaChunk2", shaChunk2)
	} else {
		if shaChunk1.Chunk != chunk1 {
			t.Error("Unexpected shaChunk1.Chunk")
		}
		if shaChunk2.Chunk != chunk2 {
			t.Error("Unexpected shaChunk2.Chunk")
		}
	}
}

func assertUnsorted(t *testing.T, chunk1, chunk2 *Chunk, output <-chan *ShaChunk) {
	shaChunk1, ok := <-output
	if !ok {
		t.Error("Unexpected ok", ok)
	}
	shaChunk2, ok := <-output
	if !ok {
		t.Error("Unexpected ok", ok)
	}
	_, ok = <-output
	if ok {
		t.Error("Unexpected ok", ok)
	}
	if shaChunk1 == nil {
		t.Error("Unexpected shaChunk1", shaChunk1)
	} else if shaChunk2 == nil {
		t.Error("Unexpected shaChunk2", shaChunk2)
	} else if shaChunk1.Chunk == chunk1 {
		if shaChunk2.Chunk != chunk2 {
			t.Error("Unexpected shaChunk2.Chunk")
		}
	} else {
		if shaChunk1.Chunk != chunk2 {
			t.Error("Unexpected shaChunk1.Chunk")
		}
		if shaChunk2.Chunk != chunk1 {
			t.Error("Unexpected shaChunk2.Chunk")
		}
	}
}
