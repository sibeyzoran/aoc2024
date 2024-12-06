package main

import (
	"bufio"
	"os"
)

func main() {
	// read the input.txt
	file, err := os.Open("input.txt")
	errCheck(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
}
