package helpers

import "fmt"

// ErrCheck checks for errors and prints them
func ErrCheck(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
