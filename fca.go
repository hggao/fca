package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func getFileInfo(file_path string) error {
	fileInfo, err := os.Stat(file_path)
	if err != nil {
		return err
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
	return nil
}

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

	if err := getFileInfo(file_path); err != nil {
		log.Fatal(err)
	}
}
