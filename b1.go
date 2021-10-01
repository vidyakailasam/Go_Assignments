package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("DB2_sample.txt")
	if err != nil {
		log.Fatalln(err)
	}

	output := bytes.Replace(input, []byte("Db2"), []byte("HI"), -1)

	err = ioutil.WriteFile("DB2_sample.txt", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
