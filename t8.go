package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	csvFile, err := os.Open("stri.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(csvFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		var value []string
		value = append(value, line[0])
		value = append(value, line[1])

		fmt.Println(value)

	}

}
