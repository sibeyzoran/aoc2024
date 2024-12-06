package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sibeyzoran/aoc2024/go-helpers"
)

func main() {
	// read the input.txt
	file, err := os.Open("input.txt")
	helpers.ErrCheck(err)

	defer file.Close()

	// turn the map into array of arrays
	var guardMap [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		guardMap = append(guardMap, []rune(scanner.Text()))
	}

	// map directions e.g. ^ means up (step down one in array index) > means right (increment positively in array) < means left (step down in the array) v means down (step up one in array index)
	directions := []struct{ x, y int }{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	// find the array that contains the ^
	var x, y, dirIdx int
	foundGuard := false

	// i is the row in the map j is the column
	for i, row := range guardMap {
		for j, cell := range row {
			if cell == '^' {
				x, y, dirIdx = i, j, 0 // we start facing up so 0 in the directions index
				foundGuard = true
				guardMap[i][j] = 'X' // this cell becomes visited because the guard starts there
				break
			}

			if foundGuard {
				break
			}
		}
	}
	// count++ for each '.' covered the '.' should be changed to an X so if it's visited again the count doesn't get increased
	// always check what char is next if its a # rotate the guards pointer: if pointer == ^ then pointer == >, if pointer == > then pointer == v, if pointer == v then pointer == <, if pointer == < then pointer == ^
	// if the next char is nil/out of range or w/e then we break
	positionsCovered := 1

	for {
		// next position - nx is the row in the map ny is the column
		nx, ny := x+directions[dirIdx].x, y+directions[dirIdx].y

		// check if it's outside the map
		if nx < 0 || ny < 0 || nx >= len(guardMap) || ny >= len(guardMap[0]) {
			fmt.Println("Guard exited the map")
			break
		}

		// check the next positions contents
		if guardMap[nx][ny] == '#' { // obstacle
			dirIdx = (dirIdx + 1) % 4 // this basically makes the guard turn right
			// fmt.Printf("Guard hit obstacle at: %v , %v \n", nx, ny)
		} else if guardMap[nx][ny] == '.' { // open space
			guardMap[nx][ny] = 'X'
			positionsCovered++
			// move the guard to the new spot
			x, y = nx, ny
		} else if guardMap[nx][ny] == 'X' { // already visited just move the guard
			x, y = nx, ny
		}
	}

	fmt.Println("Guard covered this amount of ground: ", positionsCovered)

}
