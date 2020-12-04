package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Printf("two or more args needed, got %v\n", os.Args)
		log.Println("Usage: crc32 -c/-k/-i file1 file2 ...")
		os.Exit(1)
	}
	var polynomial uint32
	switch os.Args[1] {
	case "-i":
		polynomial = crc32.IEEE
	case "-c":
		polynomial = crc32.Castagnoli
	case "-k":
		polynomial = crc32.Koopman
	default:
		log.Println("wrong arg")
		os.Exit(1)
	}
	for _, file := range os.Args[2:] {
		h, err := hashFileCrc32(file, polynomial)
		if err != nil {
			panic(err)
		}
		fmt.Print(h)
	}
}
func hashFileCrc32(filePath string, polynomial uint32) (string, error) {
	var result string

	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	tablePolynomial := crc32.MakeTable(polynomial)
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}

	return fmt.Sprintf("%08X", hash.Sum32()), nil
}
