package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// read the input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data string
	for scanner.Scan() {
		data += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
	}

	// use regular expression to pull out data
	pattern := `(mul\((\d+),(\d+)\)|do\(\)|don\'t\(\))` // the (\d+) captures digits, do and don't cases

	regex := regexp.MustCompile(pattern)

	matches := regex.FindAllStringSubmatch(data, -1) // submatch captures all capture groups

	// go through matches and multiply
	var result1 = 0

	// boolean true to process, false to skip
	skip := false

	// multiply
	for _, match := range matches {
		switch {
		case match[0] == "do()":
			skip = false
		case match[0] == "don't()":
			skip = true
		case len(match) == 4 && !skip:
			num1, _ := strconv.Atoi(match[2])
			num2, _ := strconv.Atoi(match[3])
			result1 += num1 * num2
		}
	}

	fmt.Println("Result 1 adds up to: ", result1)
}
