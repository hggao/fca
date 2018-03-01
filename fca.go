package main

import (
	"flag"
	"fmt"
	"io"
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
	var lastByte byte = 0
	var lastCnt uint64 = 0
	buffer := make([]byte, 4096)

	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Reached EOF")
			} else {
				log.Fatal(err)
			}
			break
		}
		for i := 0; i < bytesRead; i++ {
			counts[buffer[i]].count_total += 1
			if buffer[i] == lastByte {
				lastCnt += 1
				if lastCnt > counts[buffer[i]].block_max {
					counts[buffer[i]].block_max = lastCnt
				}
			} else { //A new value
				if lastCnt < 2 {
				} else if lastCnt == 2 {
					counts[lastByte].count_2 += 1
				} else if lastCnt <= 4 {
					counts[lastByte].count_3_4 += 1
				} else if lastCnt <= 8 {
					counts[lastByte].count_5_8 += 1
				} else if lastCnt <= 512 {
					counts[lastByte].count_9_512 += 1
				} else {
					counts[lastByte].count_512p += 1
				}
				lastByte = buffer[i]
				lastCnt = 1
			}
		}
	}

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
