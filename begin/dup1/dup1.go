package main
// create dup1 like uniq

import (
	"os"
	"fmt"
	"bufio"
	"log"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lastline string
	for scanner.Scan() {
		curline := scanner.Text()
		if curline != lastline {
			fmt.Println(curline)
			lastline = curline
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}