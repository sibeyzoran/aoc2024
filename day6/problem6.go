package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sibeyzoran/aoc2024/go-helpers"
)

func main() {
	// Read the input.txt
	file, err := os.Open("input.txt")
	helpers.ErrCheck(err)
	defer file.Close()

	// Turn the map into a 2D slice
	var guardMap [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		guardMap = append(guardMap, []rune(scanner.Text()))
	}

	// Map directions
	directions := []struct{ x, y int }{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	// Find guard's starting position
	var guardStartPosX, guardStartPosY, dirIdx int
	for i, row := range guardMap {
		for j, col := range row {
			if col == '^' {
				guardStartPosX, guardStartPosY, dirIdx = i, j, 0
				guardMap[i][j] = '.' // Replace '^' with open space
				break
			}
		}
	}
	simulateGuardWithObstacle := func(mapCopy [][]rune) bool {
		visited := make(map[string]bool)
		x, y, direction := guardStartPosX, guardStartPosY, 0 // Start from the original position and direction

		for {
			// Track the current state (position and direction)
			state := fmt.Sprintf("%d,%d,%d", x, y, direction)
			if visited[state] {
				fmt.Printf("Loop detected at state: %s\n", state)
				return true // Loop detected
			}
			visited[state] = true

			// Calculate the next position
			nx, ny := x+directions[direction].x, y+directions[direction].y

			// Check bounds
			if nx < 0 || ny < 0 || nx >= len(mapCopy) || ny >= len(mapCopy[0]) {
				return false // Guard leaves the map
			}

			// Check the next position's contents
			if mapCopy[nx][ny] == '#' { // Obstacle
				direction = (direction + 1) % 4 // Turn right
			} else { // Open space
				x, y = nx, ny // Move the guard
			}
		}
	}

	// PART 1: Simulate Guard's Movement
	positionsCovered := 1
	x, y, direction := guardStartPosX, guardStartPosY, dirIdx
	for {
		// Calculate the next position
		nx, ny := x+directions[direction].x, y+directions[direction].y

		// Check bounds
		if nx < 0 || ny < 0 || nx >= len(guardMap) || ny >= len(guardMap[0]) {
			break // Guard leaves the map
		}

		// Check the next position's contents
		if guardMap[nx][ny] == '#' {
			direction = (direction + 1) % 4 // Turn right
		} else if guardMap[nx][ny] == '.' {
			guardMap[nx][ny] = 'X' // Mark as visited
			positionsCovered++
			x, y = nx, ny
		} else if guardMap[nx][ny] == 'X' {
			x, y = nx, ny
		}
	}
	fmt.Println("Guard covered this amount of ground:", positionsCovered)

	// PART 2: Test Obstructions for Loops
	possibleLoops := 0
	for i, row := range guardMap {
		for j, col := range row {
			if col != '.' {
				continue
			}

			// Create a copy of the map with a new obstruction
			mapCopy := make([][]rune, len(guardMap))
			for k := range guardMap {
				mapCopy[k] = make([]rune, len(guardMap[k]))
				copy(mapCopy[k], guardMap[k])
			}
			mapCopy[i][j] = '#' // Place obstruction

			// Simulate and check for a loop
			if simulateGuardWithObstacle(mapCopy) {
				possibleLoops++
			}
		}
	}
	fmt.Println("Possible loops detected:", possibleLoops)
}
