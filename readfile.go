package main

import (
	"fmt"
	"bufio"
	"os"
	"compress/gzip"
	"strings"
	"strconv"
)

type Q int32
const (
	Q_unknown Q  = iota
	Q_phred33
	Q_phred64
	Q_solexa
)
var results = [] string{}
type AudienceCategory struct {
	Cat int
	Sum int
}
var counter = []AudienceCategory{}


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
	
	// counter := map[int] string{}
	if len(command) > 0 {
		fq := NewFastqFile(command)
		var i int = 0
		
		for {
			line, _, err := fq.File.ReadLine()
			i++
			results = append(results, string(line))
			
			if err != nil {
				// at the end of the loop...
				i = 0
				for _, v := range results {
					words := strings.Fields(v)
					
					if len(words) >1 {
						for i := 1; i < len(words); i++ {
							var word string = words[i]
							s := strings.Split(word, ":")
							cat,stat := s[0],s[len(s)-1]
    						status, _ := strconv.ParseFloat(stat, 64)
    						category, _ := strconv.Atoi(cat)
    						
    						if status >= 0 {
	    						// insert
	    						counter = floatInSlice(category,counter)
								if counter != nil{}
    						} else if status < 0 {
    						// 	// remove from counter array
    						}
						}
					}
					i++
				}
				// fmt.Printf("%v ", counter) 
				for _, b := range counter {
					fmt.Printf("Category: %d - Sum: %d\n\n",b.Cat,b.Sum) 
				}
				fmt.Printf("\n -------- LOOP ------\n") 
				os.Exit(1)
			}
		}
		
	} else if len(command) < 1 {
		fmt.Println("No argument passed")
		os.Exit(1)
	}
	os.Exit(1)
	
}//end of main


func floatInSlice(category int, counter []AudienceCategory) []AudienceCategory {
	
	for i := 0; i < len(counter); i++ {
		if counter[i].Cat == category {
            counter[i].Sum = counter[i].Sum+1

            return counter
        }
	}
    x := AudienceCategory{Cat:category,Sum:1}
    counter = append( counter, x)
    return counter
}
