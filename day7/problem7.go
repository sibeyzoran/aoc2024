package main

import (
	"os"
	"strings"

	"github.com/sibeyzoran/aoc2024/go-helpers"
)

func main() {
	// read the input.txt
	file, err := os.ReadFile("input.txt")
	helpers.ErrCheck(err)

	defer file.Close()

	equals := make([][]int, 0)
	for _, line := range strings.Split(file, "\n") {
		parts := strings.Split(line, ": ")
		row := make([]int, 0)
		row := append(row, helpers.ToInt(parts[0]))

	}

}
