package main

// Modify dup2 to print the names of all files
// in which each duplicated line occurs

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsWithFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countsWithFiles, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countsWithFiles, arg) //passing each file into countLines

			//if counts increments on a word with a different arg name (file name) append to
			//output string

			f.Close()
		}
	}

	for k, v := range countsWithFiles {
		if counts[k] > 1 {
			fmt.Println(k)
			fmt.Println("\t duplicated in ", v)

		}
	}
}

func countLines(f *os.File, counts map[string]int, countsWithFiles map[string][]string, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		countsWithFiles[input.Text()] = append(countsWithFiles[input.Text()], fileName)
	}

}
