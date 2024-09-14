package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func explainUsage(){
	fmt.Println("Flags: -b (number of bytes), -l (number of lines), -w (number of words), -c (number of characters)\n")
	fmt.Println("explainUsage:\nccwc <Flag> <Filename>")
	fmt.Fprintf(os.Stdout, "\033[0;34m %s  \033[0m %s", "Note:", "Omitting a <Flag> displays -b, -l, and -w together.")
	os.Exit(1)
}

func scanWithDelimiter(f *os.File, b bufio.SplitFunc) {
	scanner := bufio.NewScanner(f)
	scanner.Split(b)	

	var count int

	for scanner.Scan() {
		count++
	}

	if _, rewindErr := f.Seek(0, 0); rewindErr != nil {
		fmt.Println("Error occurred while scanning the file:", rewindErr)
		os.Exit(1)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error occurred while scanning the file:", err)
		os.Exit(1)
	}

	fmt.Println(count)
}

func displaySpecificDetail(f *os.File, flag string) {

	switch flag {
	case "-b":
	scanWithDelimiter(f, (bufio.SplitFunc)(bufio.ScanBytes))
	case "-l": 
	scanWithDelimiter(f, (bufio.SplitFunc)(bufio.ScanLines))
	case "-w": 
	scanWithDelimiter(f, (bufio.SplitFunc)(bufio.ScanWords))
	case "-c": 
	scanWithDelimiter(f, (bufio.SplitFunc)(bufio.ScanRunes))
	default: 
	fmt.Println("Invalid flag provided. Please refer to the flags down below:")

	explainUsage()
	}
}

func main(){
	
	numArgs := len(os.Args[1:]) 
	switch {
	case numArgs < 1:
		explainUsage()
	case numArgs == 1 || numArgs == 2:

		var fileIndex int = 2

		if numArgs == 1 {
			fileIndex = 1	
		} 

		filename := os.Args[fileIndex]
		f, err := os.Open(filename)
		
		if numArgs == 2 {
			displaySpecificDetail(f, os.Args[2])	
		} else {

		scanWithDelimiter(f, (bufio.SplitFunc)(bufio.ScanBytes))
		scanWithDelimiter(f, (bufio.SplitFunc)(bufio.ScanLines))
		scanWithDelimiter(f, (bufio.SplitFunc)(bufio.ScanWords))

		if err != nil {
			log.Fatalf("Unable to read file: %v", err)
		}

		defer f.Close()
		}
	}

}
	