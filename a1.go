package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("DB2_sample.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	counts := wordCount(file)
	fmt.Println("Number of count is", counts["Db2"])
}

func wordCount(rdr io.Reader) map[string]int {
	counts := map[string]int{}
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()

		counts[word]++
	}
	return counts
}
