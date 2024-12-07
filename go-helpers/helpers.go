package helpers

import (
	"fmt"
	"strconv"
)

// ErrCheck checks for errors and prints them
func ErrCheck(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func ToInt(s string) int {
	var number, err = strconv.Atoi(s)
	if err != nil {
		fmt.Println("Failed to convert string number to int")
	}
	return number
}
