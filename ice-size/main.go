package main

import (
	"flag"
	"fmt"
	. "github.com/daniel-fanjul-alcuten/goice"
	"log"
	"os"
)

func main() {

	flag.Parse()
	args := flag.Args()

	if err := list(args); err != nil {
		log.Fatalln(err)
	}

}

func list(filenames []string) error {
	count := 0
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		reader, err := NewBergReader(file)
		if err != nil {
			file.Close()
			return err
		}
		for {
			_, sha, err := reader.Read()
			if err != nil {
				file.Close()
				return err
			}
			if sha != nil {
				if err := file.Close(); err != nil {
					return err
				}
				break
			}
			count++
		}
	}
	fmt.Println(count)
	return nil
}
