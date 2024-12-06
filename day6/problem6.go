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
	var guardStartPosX, guardStartPosY, dirIdx int
	foundGuard := false
	// i is the row in the map j is the column
	for i, row := range guardMap {
		for j, col := range row {
			if col == '^' {
				guardStartPosX, guardStartPosY, dirIdx = i, j, 0 // we start facing up so 0 in the directions index
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
	// simulates the guards movements with a new obstacle placed in it
	simulateGuardWithObstacle := func(mapCopy [][]rune) bool {
		visited := make(map[string]bool)
		x, y, direction := guardStartPosX, guardStartPosY, 0 // start from the original position and direction

		for {
			// track the current state (position and direction)
			state := fmt.Sprintf("%d,%d,%d", x, y, direction)
			if visited[state] {
				fmt.Printf("Loop detected at state: %s\n", state)
				return true // loop is detected
			}
			visited[state] = true

			// calculate the next position
			nx, ny := x+directions[direction].x, y+directions[direction].y

			// check if it's outside the map
			if nx < 0 || ny < 0 || nx >= len(mapCopy) || ny >= len(mapCopy[0]) {
				return false // guard leaves the map
			}

			// check the next position's contents
			if mapCopy[nx][ny] == '#' { // obstacle
				direction = (direction + 1) % 4 // turn right
			} else { // open space or visited
				x, y = nx, ny // move the guard
			}
		}
	}

	// PART 1
	positionsCovered := 1

	for {
		// next position - nx is the row in the map ny is the column
		nx, ny := guardStartPosX+directions[dirIdx].x, guardStartPosY+directions[dirIdx].y

		// check if it's outside the map
		if nx < 0 || ny < 0 || nx >= len(guardMap) || ny >= len(guardMap[0]) {
			fmt.Println("Guard exited the map")
			break
		}

		// check the next positions contents
		if guardMap[nx][ny] == '#' { // obstacle
			dirIdx = (dirIdx + 1) % 4 // this basically makes the guard turn right
		} else if guardMap[nx][ny] == '.' { // open space
			guardMap[nx][ny] = 'X'
			positionsCovered++
			// move the guard to the new spot
			guardStartPosX, guardStartPosY = nx, ny
		} else if guardMap[nx][ny] == 'X' { // already visited just move the guard
			guardStartPosX, guardStartPosY = nx, ny
		}
	}

	fmt.Println("Guard covered this amount of ground: ", positionsCovered)

	// PART 2
	possibleLoops := 0
	for i, row := range guardMap {
		for j, col := range row {
			// skip invalid positions
			if col != '.' {
				continue
			}

			// create a copy of the map with a new obstruction in the location of i & j
			mapCopy := make([][]rune, len(guardMap)) // if i just copy the guardMap it fucks up cause they then share the same underlying data - this creates an empty array of arrays with the length of guard map but it has no data
			for k := range guardMap {
				mapCopy[k] = make([]rune, len(guardMap[k])) // this allocated a new row with the same length as the one in the guardmap
				copy(mapCopy[k], guardMap[k])               // this copies the data of the row from guardMap to the mapCopy
			}
			mapCopy[i][j] = '#'

			// simulate and check for a loop

			if simulateGuardWithObstacle(mapCopy) {
				possibleLoops++
			}
		}
	}

	fmt.Println("Possible loops detected: ", possibleLoops)

}
