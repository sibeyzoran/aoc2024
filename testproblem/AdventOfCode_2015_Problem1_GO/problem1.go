package main

import (
	"fmt"
	"os"
)

func main() {
	var up rune = '('
	var down rune = ')'

	var floor int = 0
	var position int = 1

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	content := string(data)

	for i := 0; i < len(content); i++ {
		char := rune(content[i])

		if char == up {
			floor++
		} else if char == down {
			floor--
		}
		if floor == -1 {
			fmt.Printf("Floor: %d, Position: %d\n", floor, position)
			return
		}
		position++
	}
}
