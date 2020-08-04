package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	for i := 1; i < len(os.Args); i++ {
		var iteration string
		iteration = os.Args[i] + " " + strconv.Itoa(i)
		fmt.Println(iteration)
	}
}
