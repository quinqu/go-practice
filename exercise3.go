package main

// Exercise 1.3: Experiment to measure the difference in running time between our
// potentially inefficient versions and the one that uses strings.Join.
// (Section 1.6 illustrates part of the
// time package, and Section 11.4 shows how to write benchmark tests for
// systematic performance evaluation.

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// first version
	start := time.Now()
	strings.Join(os.Args[1:], " ")
	fmt.Printf("%.9fs elapsed\n", time.Since(start).Seconds())

	// second version (not as efficient)
	start2 := time.Now()
	var s, space string
	for i := 1; i < len(os.Args); i++ {
		s += space + os.Args[i]
		space = ""
	}

	fmt.Printf("%.9fs elapsed\n", time.Since(start2).Seconds())

}
