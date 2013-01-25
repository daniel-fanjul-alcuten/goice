package ice

func HashChunksSorted(input <-chan *Chunk, output chan<- *ShaChunk, size int) {
	waiting := make(chan chan *ShaChunk, size)
	go func() {
		for chunk := range input {
			wait := make(chan *ShaChunk, 1)
			waiting <- wait
			go func(chunk *Chunk) {
				wait <- chunk.Hash()
			}(chunk)
		}
		close(waiting)
	}()
	for wait := range waiting {
		output <- <-wait
	}
	close(output)
}

func HashChunksUnsorted(input <-chan *Chunk, output chan<- *ShaChunk, size int) {
	waiting := make(chan chan bool, size)
	go func() {
		for chunk := range input {
			wait := make(chan bool, 1)
			waiting <- wait
			go func(chunk *Chunk) {
				output <- chunk.Hash()
				wait <- true
			}(chunk)
		}
		close(waiting)
	}()
	for wait := range waiting {
		<-wait
	}
	close(output)
}

func HashChunks(input <-chan *Chunk, sorted, unsorted chan<- *ShaChunk, size int) {
	waiting := make(chan chan *ShaChunk, size)
	go func() {
		for chunk := range input {
			wait := make(chan *ShaChunk, 1)
			waiting <- wait
			go func(chunk *Chunk) {
				shaChunk := chunk.Hash()
				unsorted <- shaChunk
				wait <- shaChunk
			}(chunk)
		}
		close(waiting)
	}()
	for wait := range waiting {
		sorted <- <-wait
	}
	close(sorted)
	close(unsorted)
}
