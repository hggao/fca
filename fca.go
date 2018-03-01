package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	n_worker := flag.Int("workers", 1, "Number of workers")
	flag.Parse()
	flag.Usage = func() {
		fmt.Println("Usage:\n ", os.Args[0], "[OPTIONS] file")
		fmt.Println("\nOPTIONS:")
		flag.PrintDefaults()
		fmt.Println("\nExample:")
		fmt.Println(" ", os.Args[0], "-workers=8 test.raw\n")
	}
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	file_path := flag.Args()[0]

	fmt.Println("File to analysis:", file_path)
	fmt.Println("Number of workers:", *n_worker)

	fileInfo, err := os.Stat(file_path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo)
}
