package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Reader reads an input from an http site
type Reader struct {
	address string
	file    string
}

func NewReader(day int) *Reader {
	return &Reader{
		address: makeAddressString(day),
		file:    makeFilePath(day),
	}
}

func makeAddressString(day int) string {
	return fmt.Sprintf("http://adventofcode.com/2021/day/%d/input", day)
}

func makeFilePath(day int) string {
	return fmt.Sprintf("./input/%d.txt", day)
}

func (r *Reader) GetInputHttp() string {
	response, err := http.Get(r.address)
	if err != nil {
		panic(fmt.Sprintf("Http error retrieving Input ", err))
	}
	// if response.StatusCode != http.StatusOK {
	// 	panic(fmt.Sprintf("Error retrieving Input ", err))
	// }
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(fmt.Sprintf("Error reading data ", err))
	}
	return string(data)
}

func (r *Reader) GetInputLines(day int) ([]string, error) {
	file, err := os.Open(r.file)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
