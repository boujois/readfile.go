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
		// fmt.Println("gziped file")
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
				for i := 5; i < len(results); i++ {
				// for _, v := range results {
					words := strings.Fields(results[i])
					
					if len(words) >1 {
						for i := 1; i < len(words); i++ {
							var word string = words[i]
							s := strings.Split(word, ":")
							cat,stat := s[0],s[len(s)-1]
							//convert to string to int
    						status, _ := strconv.Atoi(stat)
    						category, _ := strconv.Atoi(cat)
    						counter = updateCounter(category,counter,status)
							if counter != nil{}
    						
						}
					}
					i++
				}
				// fmt.Printf("%v ", counter) 
				fmt.Printf("Total number of segments: %d \n",len(counter))
				fmt.Printf("\n  Segment        Sum     \n") 
				fmt.Printf("+--------------+---------+\n") 
				for _, b := range counter {
					if b.Sum <10{
						fmt.Printf("| %d        | %d       |\n",b.Cat,b.Sum)
					} 
					if b.Sum > 9 && b.Sum < 99{
						fmt.Printf("| %d        | %d      |\n",b.Cat,b.Sum)
					}
					if b.Sum > 99 && b.Sum < 999 {
						fmt.Printf("| %d        | %d     |\n",b.Cat,b.Sum)
					}
					if b.Sum > 999 && b.Sum < 9999 {
						fmt.Printf("| %d        | %d    |\n",b.Cat,b.Sum)
					}
					 
					fmt.Printf("+--------------+---------+\n") 
				}
				os.Exit(1)
			}
		}
		
	} else if len(command) < 1 {
		fmt.Println("No argument passed")
		os.Exit(1)
	}
	os.Exit(1)
	
}//end of main


func updateCounter(category int, counter []AudienceCategory,status int) []AudienceCategory {
	
	for i := 0; i < len(counter); i++ {
		if counter[i].Cat == category {
			if status >= 0 {
				// insert
				counter[i].Sum = counter[i].Sum+1
			} else if status < 0 {
				// remove from counter array
				counter[i].Sum = counter[i].Sum-1
			}
            return counter
        }
	}
    x := AudienceCategory{Cat:category,Sum:1}
    counter = append( counter, x)
    return counter
}
