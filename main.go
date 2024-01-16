package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var printNumBytes, printNumLines, printNumWords, printNumChars bool
var filePath string
var inputSource string

func defineFlags() {
	flag.BoolVar(&printNumBytes, "c", false, "Print number of bytes in the file")
	flag.BoolVar(&printNumLines, "l", false, "Print number of lines in the file")
	flag.BoolVar(&printNumWords, "w", false, "Print number of words in the file")
	flag.BoolVar(&printNumChars, "m", false, "Print number of characters in the file")
	flag.Parse()
	filePath = flag.CommandLine.Arg(0)
}

func readData() {
	stdin := os.Stdin
	data, err := stdin.Stat()
	handlePathError(err)
	if data.Size() > 0 {
		inputSource = "stdin"
	} else {
		inputSource = "file"
	}
}

func handlePathError(err error) {
	if err != nil {
		fmt.Printf("Error -> %s\n", err.Error())
		os.Exit(1)
	}
}

func closeFile(file *os.File) {
	if inputSource == "file" {
		file.Close()
	}
}

func setReader() (*bufio.Scanner, *os.File) {
	var fs *bufio.Scanner
	var file *os.File
	var err error
	if inputSource == "file" {
		file, err = os.Open(filePath)
		handlePathError(err)
		fs = bufio.NewScanner(file)
	} else {
		fs = bufio.NewScanner(os.Stdin)
	}
	return fs, file
}

func scanNumBytes() {
	fs, file := setReader()
	byteCount := 0
	fs.Split(bufio.ScanBytes)
	for fs.Scan() {
		byteCount++
	}
	fmt.Println("\tc:", byteCount)
	closeFile(file)
}

func scanNumLines() {
	fs, file := setReader()
	lineCount := 0
	fs.Split(bufio.ScanLines)
	for fs.Scan() {
		lineCount++
	}
	fmt.Println("\tl:", lineCount)
	closeFile(file)
}

func scanNumWords() {
	fs, file := setReader()
	wordCount := 0
	fs.Split(bufio.ScanWords)
	for fs.Scan() {
		wordCount++
	}
	fmt.Println("\tw:", wordCount)
	closeFile(file)
}

func scanNumChars() {
	fs, file := setReader()
	charCount := 0
	fs.Split(bufio.ScanRunes)
	for fs.Scan() {
		charCount++
	}
	fmt.Println("\tm:", charCount)
	closeFile(file)
}

func main() {
	defineFlags()
	readData()
	if flag.NFlag() == 0 {
		printNumBytes = true
		printNumLines = true
		printNumWords = true
	}
	if printNumBytes {
		scanNumBytes()
	}
	if printNumLines {
		scanNumLines()
	}
	if printNumWords {
		scanNumWords()
	}
	if printNumChars {
		scanNumChars()
	}
	fmt.Println("\tin:", filePath)
}
