package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func getflags() (a, b *string) {
	var file = flag.String("file", "b", "path to file")
	var substring = flag.String("substring", "a", "string to find")
	flag.Parse()
	// fmt.Println(file.)
	return file, substring
}

func scanFile(file *os.File, substr *string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), *substr) {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	filePath, substr := getflags()
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanFile(file, substr)

}
