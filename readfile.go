package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func main() {

	file, err := os.Open("/tmp/sample.txt")
	// if there was a problem reading the file, exit with error
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()

	// read each line in for loop
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	    fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}
}