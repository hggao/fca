package main

import (
	"flag"
	"fmt"
)

func main() {
	n_worker := flag.Int("workers", 1, "Number of workers")
	flag.Parse()

	fmt.Println("Number of workers: ", *n_worker)
}
