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

func fileContentAnalysis(file_path string, fileSize int64, workers int) []ByteStats {
	counts := make([]ByteStats, 256)
	var lastByte byte = 0
	var lastCnt uint64 = 0
	buffer := make([]byte, 4096)

	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalRead int64 = 0
	fmt.Printf("Processing %5.2f completed.", float64(totalRead)/float64(fileSize)*100.0)
	for {
		bytesRead, err := file.Read(buffer)
		totalRead += int64(bytesRead)
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
		fmt.Printf("\rProcessing %5.2f completed.", float64(totalRead)/float64(fileSize)*100.0)
	}

	return counts
}

func printCountResult(counts []ByteStats, fileSize int64) {
	fmt.Println("================= file content statistic =================")
	for i := 0; i < 256; i++ {
		if counts[i].count_total > 0 {
			fmt.Printf("Byte: %3d, Total: %-10d, 2-2: %-10d, 3-4: %-10d, 5-8: %-10d, 9-512: %-10d, 512+: %-10d, block_max: %-10d, percent: %5.2f\n",
				i, counts[i].count_total, counts[i].count_2, counts[i].count_3_4, counts[i].count_5_8,
				counts[i].count_9_512, counts[i].count_512p, counts[i].block_max, float64(counts[i].count_total)/float64(fileSize)*100)
		}
	}
}

func getFileInfo(file_path string) (int64, error) {
	fileInfo, err := os.Stat(file_path)
	if err != nil {
		return 0, err
	}
	fmt.Println("================= file info =================")
	fmt.Println("File name:            ", fileInfo.Name())
	fmt.Println("Size in bytes:        ", fileInfo.Size())
	fmt.Println("Permissions:          ", fileInfo.Mode())
	fmt.Println("Last modified:        ", fileInfo.ModTime())
	fmt.Println("Is Directory:         ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info:           %+v\n\n", fileInfo.Sys())
	return fileInfo.Size(), nil
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

	fileSize, err := getFileInfo(file_path)
	if err != nil {
		log.Fatal(err)
	}

	counts := fileContentAnalysis(file_path, fileSize, *n_worker)
	printCountResult(counts, fileSize)

	fmt.Println("Done, Bye!")
}
