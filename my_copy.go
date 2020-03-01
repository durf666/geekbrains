package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
func main() {
	var from = flag.String("from", "b", "path to file destination")
	var to = flag.String("to", "a", "path to new file destination")
	flag.Parse()
	a, b := copy(*from, *to)
	fmt.Println(a, b)
}
