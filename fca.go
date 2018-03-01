package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type ByteStats struct {
	count_total uint64
	count_2     uint64
	count_3_4   uint64
	count_5_8   uint64
	count_9_512 uint64
	count_512p  uint64
	block_max   uint64
}

func fileContentAnalysis(file_path string, workers int) {
	counts := make([]ByteStats, 256)
	//var last byte = 0
	//var ccc uint64 = 0

	fmt.Println("================= file content statistic =================")
	for i := 0; i < 256; i++ {
		fmt.Printf("Byte: %3d, Total: %d, 2-2: %d, 3-4: %d, 5-8: %d, 9-512: %d, 512+: %d, block_max: %d\n",
			i, counts[i].count_total, counts[i].count_2, counts[i].count_3_4, counts[i].count_5_8,
			counts[i].count_9_512, counts[i].count_512p, counts[i].block_max)
	}
}

func getFileInfo(file_path string) error {
	fileInfo, err := os.Stat(file_path)
	if err != nil {
		return err
	}
	fmt.Println("================= file info =================")
	fmt.Println("File name:            ", fileInfo.Name())
	fmt.Println("Size in bytes:        ", fileInfo.Size())
	fmt.Println("Permissions:          ", fileInfo.Mode())
	fmt.Println("Last modified:        ", fileInfo.ModTime())
	fmt.Println("Is Directory:         ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info:           %+v\n\n", fileInfo.Sys())
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

	fileContentAnalysis(file_path, *n_worker)

	fmt.Println("Done, Bye!")
}
