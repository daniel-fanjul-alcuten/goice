package main

import (
	"flag"
	"fmt"
	. "github.com/daniel-fanjul-alcuten/goice"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func main() {

	procs := runtime.GOMAXPROCS(-1)
	read_channel_size := flag.Int("read_channel_size", 2*procs, "number of chunks in the reading channel")
	hashing_channel_size := flag.Int("hashing_channel_size", 2*procs, "number of chunks in the hashing channel")
	sorted_channel_size := flag.Int("sorted_channel_size", procs, "number of chunks in the sorted channel")
	unsorted_channel_size := flag.Int("unsorted_channel_size", procs, "number of chunks in the unsorted channel")
	write_channel_size := flag.Int("write_channel_size", 3*procs, "number of buffers in the write hashing channel")
	flag.Parse()
	args := flag.Args()

	chunks := make(chan *Chunk, *read_channel_size)
	rerrs := make(chan error, 1)
	go read(args, chunks, rerrs)

	sorted := make(chan *ShaChunk, *sorted_channel_size)
	unsorted := make(chan *ShaChunk, *unsorted_channel_size)
	go HashChunks(chunks, sorted, unsorted, *hashing_channel_size)

	name := make(chan string, 1)
	werrs := make(chan error, 1)
	go write(unsorted, name, werrs, *write_channel_size)

	for sorted != nil || rerrs != nil || werrs != nil {
		select {
		case chunk, ok := <-sorted:
			if ok {
				fmt.Println(chunk.Sha.String())
			} else {
				sorted = nil
			}
		case err, ok := <-rerrs:
			if ok {
				log.Println(err)
			} else {
				rerrs = nil
			}
		case err, ok := <-werrs:
			if ok {
				log.Println(err)
			} else {
				werrs = nil
			}
		}
	}

	if n, ok := <-name; ok {
		fmt.Println(n)
	}
}

func read(filenames []string, chunks chan<- *Chunk, errs chan<- error) {
	sent, chunk, err := false, &Chunk{}, error(nil)
	for _, filename := range filenames {
		if input, err := os.Open(filename); err == nil {
			if err = chunk.ReadFrom(input); err == nil {
				chunks <- chunk
				sent, chunk = true, &Chunk{}
				err = chunk.ReadFrom(input)
				for err == nil {
					chunks <- chunk
					chunk = &Chunk{}
					err = chunk.ReadFrom(input)
				}
			}
			if err == io.EOF {
				err = nil
			}
			if err != nil {
				errs <- err
			}
			if err2 := input.Close(); err2 != nil {
				err = err2
				errs <- err
			}
		} else {
			errs <- err
		}
		if err != nil {
			break
		}
	}
	if chunk.Len > 0 || !sent {
		chunks <- chunk
	}
	close(chunks)
	close(errs)
}

func write(chunks <-chan *ShaChunk, name chan<- string, errs chan<- error, size int) {
	file, err := ioutil.TempFile(".", ".ice-berg-")
	if err != nil {
		errs <- err
		close(errs)
		return
	}
	writer, err := NewBergWriter(file, size)
	if err != nil {
		errs <- err
	} else {
		for chunk := range chunks {
			if err = writer.Write(chunk); err != nil {
				errs <- err
				break
			}
		}
		for _ = range chunks {
		}
	}
	if err2 := writer.Close(); err2 != nil {
		err = err2
		errs <- err
	}
	var sha *Sha
	if err == nil {
		if sha, err = writer.Sha(); err != nil {
			errs <- err
		}
	}
	if err2 := file.Close(); err2 != nil {
		err = err2
		errs <- err
	}
	if err == nil {
		newName := sha.String() + ".berg"
		if err = os.Rename(file.Name(), newName); err == nil {
			name <- newName
		} else {
			errs <- err
		}
	}
	if err != nil {
		if err = os.Remove(file.Name()); err != nil {
			errs <- err
		}
	}
	close(errs)
}
