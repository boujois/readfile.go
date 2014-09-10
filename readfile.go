package main

import (
	"fmt"
	"bufio"
	"os"
	"compress/gzip"
)

type Q int32
const (
	Q_unknown Q  = iota
	Q_phred33
	Q_phred64
	Q_solexa
)

type FastqFile struct {
	File *bufio.Reader
	Q_encoding Q
	N int64
}

func NewFastqFile(filename string) *FastqFile {
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open file %s: error: %s\n",filename, err)
		os.Exit(1)
	}
	// try to open as a gzip file
	fz, err := gzip.NewReader(fh)
	if err == nil {
		fmt.Println("gziped file")
		return &FastqFile{bufio.NewReader(fz), Q_unknown, 0}
	}
	// fall back on uncompressed bufio reader
	fmt.Println("Not a gziped file")
	return &FastqFile{bufio.NewReader(fh), Q_unknown, 0}
}


func main() {
	command := string(os.Args[1])
	if len(command) > 0 {
		fq := NewFastqFile(command)
		i := 0
		results := []string{}
		for {
			line, _, err := fq.File.ReadLine()
			i++
			// if err == nil {
			//     fmt.Println("Successfully read file")
			//     os.Exit(0)
			// }
			if err != nil {
				fmt.Println("Error reading file")
				os.Exit(1)
			}
			results = append(results, string(line))
			fmt.Printf("%q \n",string(line))
		}
	} else if len(command) < 1{
		//this catch isn't working
		fmt.Println("No argument passed")
		os.Exit(1)
	}

	
}

