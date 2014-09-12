package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "time"
)

type Q int32

const (
	Q_unknown Q = iota
	Q_phred33
	Q_phred64
	Q_solexa
)

var results = []string{}

type AudienceCategory struct {
	Cat int
	Sum int
	Additions int
	Subtractions int
}

var counter = []AudienceCategory{}

type FastqFile struct {
	File       *bufio.Reader
	Q_encoding Q
	N          int64
}

func NewFastqFile(filename string) *FastqFile {
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open file %s: error: %s\n", filename, err)
		os.Exit(1)
	}
	// try to open as a gzip file
	fz, err := gzip.NewReader(fh)
	if err == nil {
		// fmt.Println("gziped file")
		return &FastqFile{bufio.NewReader(fz), Q_unknown, 0}
	}
	// fall back on uncompressed bufio reader
	fmt.Println("Not a gziped file")
	return &FastqFile{bufio.NewReader(fh), Q_unknown, 0}
}

func updateCounter(category int, counter []AudienceCategory, status int) []AudienceCategory {
	for i := 0; i < len(counter); i++ {
		if counter[i].Cat == category {
			if status >= 0 {
				// insert
				counter[i].Sum = counter[i].Sum + 1
				counter[i].Additions = counter[i].Additions + 1
			} else if status < 0 {
				// remove from counter array
				counter[i].Sum = counter[i].Sum - 1
				counter[i].Subtractions = counter[i].Subtractions + 1
			}
			return counter
		}
	}
	x := AudienceCategory{Cat: category, Sum: 1,Additions:1,Subtractions:0}
	counter = append(counter, x)
	return counter
}


func main() {
	// start := time.Now()
	// the the bash terminal argument
	command := string(os.Args[1])

	if len(command) > 0 {
		fq := NewFastqFile(command)

		i := 0
		for {
			line, _, err := fq.File.ReadLine()
			//skip first few lines
			if i > 4 {
				// results = append(results, string(line))
				// skip the first 5 lines of file with metadata
				words := strings.Fields(string(line))
				if len(words) > 1 {
					for i := 1; i < len(words); i++ {
						var word string = words[i]
						s := strings.Split(word, ":")
						cat, stat := s[0], s[len(s)-1]
						//convert from string to int
						status, _ := strconv.Atoi(stat)
						category, _ := strconv.Atoi(cat)
						counter = updateCounter(category, counter, status)
						if counter != nil {
						}
					}
				}
			}
			i++
			
			// at the end of the loop...
			if err != nil {
				fmt.Printf("Segment,Segment Users,Total Additions,Total Subtractions\n")
				for _, b := range counter {
					fmt.Printf("%d,%d,%d,%d\n", b.Cat, b.Sum,b.Additions,b.Subtractions)
				}
				os.Exit(1)
			}
		}

	} else if len(command) < 1 {
		fmt.Println("No argument passed")
		os.Exit(1)
	}
	os.Exit(1)
} //end of main
