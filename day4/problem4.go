package main

import (
	"bufio"
	"fmt"
	"os"
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

	// need to turn input into a grid
	scanner := bufio.NewScanner(file)
	var grid [][]string

	for scanner.Scan() {
		line := scanner.Text() // read each line

		row := strings.Split(line, "") // splits into characters

		grid = append(grid, row) // append row to grid
	}

	// word to find
	word := "XMAS"

	xMAScount := wordSearch(grid, word)
	fmt.Println("Number of XMAS found: ", xMAScount)

	// A is always in the center and has . and . next to it M always top left S always top right M always bottom left and S alright bottom right
	MAScount := countXMASPatterns(grid)
	fmt.Println("M.A.S found: ", MAScount)
}

func wordSearch(grid [][]string, word string) int {
	// directions for moving in the grid
	directions := [][]int{
		{0, 1},   // right
		{1, 0},   // down
		{0, -1},  // left
		{-1, 0},  // up
		{1, 1},   // down-right
		{1, -1},  // down-left
		{-1, 1},  // up-right
		{-1, -1}, // up-left
	}

	count := 0

	// iterate over each cell in the grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// check all directions from this cell
			for _, dir := range directions {
				if searchFrom(grid, word, i, j, dir) {
					count++
				}
			}
		}
	}
	return count
}

func searchFrom(grid [][]string, word string, row, col int, direction []int) bool {
	r, c := row, col
	for k := 0; k < len(word); k++ {
		// check bounds and character match
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != string(word[k]) {
			return false
		}
		// move in the current direction
		r += direction[0]
		c += direction[1]
	}
	return true
}

func countXMASPatterns(grid [][]string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// iterate through the grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// check all conditions for a valid M.A.S pattern
			if isXMASPattern(grid, i, j) {
				count++
			}
		}
	}

	return count
}

func isXMASPattern(grid [][]string, row, col int) bool {
	// check bounds and the center 'A'
	if row-1 < 0 || row+1 >= len(grid) || col-1 < 0 || col+1 >= len(grid) {
		return false
	}
	if grid[row][col] != "A" {
		return false
	}

	// check the top-left and bottom-right diagonals
	validDiagonal1 := (grid[row-1][col-1] == "M" && grid[row+1][col+1] == "S") ||
		(grid[row-1][col-1] == "S" && grid[row+1][col+1] == "M")

	// check the top-right and bottom-left diagonals
	validDiagonal2 := (grid[row-1][col+1] == "S" && grid[row+1][col-1] == "M") ||
		(grid[row-1][col+1] == "M" && grid[row+1][col-1] == "S")

	return validDiagonal1 && validDiagonal2
}
