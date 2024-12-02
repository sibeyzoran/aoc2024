package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)

	// get the numbers from input.txt and  then add to lists
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		if len(parts) == 2 {
			var number1, err = strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Failed to convert string number 1 to int")
			}

			var number2, err2 = strconv.Atoi(parts[1])
			if err2 != nil {
				fmt.Println("Failed to convert string number 2 to int")
			}

			list1 = append(list1, number1)
			list2 = append(list2, number2)
		}
	}

	// sort the lists from low to high
	sort.Ints(list1)
	sort.Ints(list2)

	// find which one's bigger and then find the diff and store in it's own list
	var diffList []int
	for i := 0; i < len(list1); i++ {
		if list1[i] >= list2[i] {
			diffList = append(diffList, list1[i]-list2[i])
		} else {
			diffList = append(diffList, list2[i]-list1[i])
		}
	}

	// add up the differences
	var result int
	for i := 0; i < len(diffList); i++ {
		result = result + diffList[i]
	}

	fmt.Println("Difference Result: ", result)

	// dictionary to map number of times number in the list1 appears in the list2
	var dictionary = make(map[int]int)

	for i := 0; i < len(list1); i++ {
		dictionary[list1[i]] = 0
	}

	// go through list1, find the key in dictionary, increase the value by one
	for i := 0; i < len(list2); i++ {
		key := list2[i]
		if _, exists := dictionary[key]; exists {
			dictionary[key]++
		}
	}

	var similarityList []int

	for key, value := range dictionary {
		var similarityScore = key * value
		similarityList = append(similarityList, similarityScore)
	}

	var result2 int
	for _, value := range similarityList {
		result2 = result2 + value
	}

	fmt.Println("Similarity score: ", result2)

}
