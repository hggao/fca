package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	n_worker := flag.Int("workers", 1, "Number of workers")
	flag.Parse()
	flag.Usage = func() {
		fmt.Println("Usage:\n ", os.Args[0], "file_path [OPTIONS]")
		fmt.Println("\nOPTIONS:")
		flag.PrintDefaults()
		fmt.Println("\nExample:")
		fmt.Println(" ", os.Args[0], "test.raw --workers=8\n")
	}
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Number of workers: ", *n_worker)
}
