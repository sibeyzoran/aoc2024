package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		valuesString := strings.Fields(line) // splits each line into an array on the white space

		var levels []int
		for _, str := range valuesString {
			val, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Couldn't convert string to integer: ", err)
				continue
			}
			levels = append(levels, val)
		}

		// check if it's a safe report!
		if isSafeWithDampener(levels) {
			safeCount++
		}
	}

	fmt.Println("Safe reports: ", safeCount)
}

func isSafeReport(levels []int) bool {
	if len(levels) < 2 {
		return false // single level or empty report can't be verified
	}

	isIncreasing := levels[1] > levels[0]
	isDecreasing := levels[1] < levels[0]

	for i := 1; i < len(levels); i++ {
		// check the diff between the 2nd node and the first - increment from there
		diff := levels[i] - levels[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false // false because if the increase is bigger than three (either way) it is unsafe
		}

		if (isIncreasing && diff <= 0) || (isDecreasing && diff >= 0) {
			return false // false because it's not strictly increasing or decreasing
		}
	}
	return true // if we get to here it's a safe report
}

func isSafeWithDampener(levels []int) bool {
	if isSafeReport(levels) {
		return true // already safe
	}

	// loop through the levels in the report and try removing and rechecking
	for i := 0; i < len(levels); i++ {
		// create a modified slice without the i-th level - have to use make here because if you alter the array by appending to the same one it freaks out
		modified := make([]int, 0, len(levels)-1)
		modified = append(modified, levels[:i]...)
		modified = append(modified, levels[i+1:]...)
		if isSafeReport(modified) {
			return true // safe after removing one level
		}
	}

	return false
}
