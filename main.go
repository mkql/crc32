package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Printf("exactly two arg needed, got %v", os.Args)
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
	h, err := hashFileCrc32(os.Args[2], polynomial)
	if err != nil {
		panic(err)
	}
	fmt.Print(h)
}
func hashFileCrc32(filePath string, polynomial uint32) (string, error) {
	//Initialize an empty return string now in case an error has to be returned
	var returnCRC32String string
	//Open the fhe file located at the given path and check for errors
	file, err := os.Open(filePath)
	if err != nil {
		return returnCRC32String, err
	}
	//Tell the program to close the file when the function returns
	defer file.Close()
	//Create the table with the given polynomial
	tablePolynomial := crc32.MakeTable(polynomial)
	//Open a new hash interface to write the file to
	hash := crc32.New(tablePolynomial)
	//Copy the file in the interface
	if _, err := io.Copy(hash, file); err != nil {
		return returnCRC32String, err
	}
	return fmt.Sprintln("%08X", hash.Sum32()), nil
}

