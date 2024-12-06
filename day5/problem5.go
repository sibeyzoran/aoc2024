package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Graph struct {
	edges    map[int][]int // adjacency list: key = node, value = list of neighbors
	vertices []int         // list of all nodes
}

func NewGraph() *Graph {
	return &Graph{
		edges:    make(map[int][]int),
		vertices: []int{},
	}
}

func createSubgraph(graph *Graph, update []int) *Graph {
	subgraph := NewGraph()
	pageSet := make(map[int]bool)
	for _, page := range update {
		pageSet[page] = true
	}

	for u, edges := range graph.edges {
		if pageSet[u] {
			for _, v := range edges {
				if pageSet[v] {
					subgraph.addEdge(u, v)
				}
			}
		}
	}

	return subgraph
}

func (g *Graph) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)

	if !contains(g.vertices, u) {
		g.vertices = append(g.vertices, u)
	}

	if !contains(g.vertices, v) {
		g.vertices = append(g.vertices, v)
	}
}

func contains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func (g *Graph) topologicalSortUtil(v int, visited map[int]bool, stack *[]int) {
	visited[v] = true

	for _, u := range g.edges[v] {
		if !visited[u] {
			g.topologicalSortUtil(u, visited, stack)
		}
	}
	*stack = append([]int{v}, *stack...)
}

func (g *Graph) topologicalSort() []int {
	stack := []int{}
	visited := make(map[int]bool)

	for _, v := range g.vertices {
		if !visited[v] {
			g.topologicalSortUtil(v, visited, &stack)
		}
	}

	return stack
}

func ToInt(s string) int {
	var number, err = strconv.Atoi(s)
	if err != nil {
		fmt.Println("Failed to convert string number to int")
	}
	return number
}

// goes through the pages in an update and if the pages appear in the same order as the sorted graph then its true
func isValidUpdate(update []int, position map[int]int) bool {
	for i := 0; i < len(update)-1; i++ {
		if position[update[i]] > position[update[i+1]] {
			return false
		}
	}
	return true
}

func fixUpdate(update []int, subgraph *Graph) []int {
	// Perform topological sort on the subgraph
	sorted := subgraph.topologicalSort()

	// Create a map for quick lookup of positions in the sorted order
	position := make(map[int]int)
	for i, page := range sorted {
		position[page] = i
	}

	// Sort the update based on the topological order
	fixedUpdate := make([]int, len(update))
	copy(fixedUpdate, update)
	sort.Slice(fixedUpdate, func(i, j int) bool {
		return position[fixedUpdate[i]] < position[fixedUpdate[j]]
	})

	return fixedUpdate
}

func main() {
	// read the input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	// split into updates and rules
	var pageOrdering []string
	var updates []string

	scanner := bufio.NewScanner(file)
	isUpdates := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip empty lines
		if line == "" {
			continue
		}

		// check if the line is part of updates or page ordering
		if strings.Contains(line, ",") {
			isUpdates = true
		}

		if isUpdates {
			updates = append(updates, line)
		} else {
			pageOrdering = append(pageOrdering, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// turn page ordering into my graph structure
	graph := NewGraph()
	for _, rule := range pageOrdering {
		parts := strings.Split(rule, "|")
		if len(parts) != 2 {
			fmt.Println("Invalid rule: ", rule)
			continue
		}

		u := ToInt(parts[0])
		v := ToInt(parts[1])
		graph.addEdge(u, v)
	}
	// ensure all pages in vertifces have an entry in the edges map
	for _, v := range graph.vertices {
		if _, exists := graph.edges[v]; !exists {
			graph.edges[v] = []int{}
		}
	}

	var parsedUpdates [][]int
	for _, update := range updates {
		parts := strings.Split(update, ",")
		var pages []int
		for _, part := range parts {
			pages = append(pages, ToInt(part))
		}
		parsedUpdates = append(parsedUpdates, pages)
	}

	// UNSORTED
	// vertices shows all the unique page numbers in the rules
	// fmt.Println("Vertices:", graph.vertices)
	// edges show the directed connections
	// fmt.Println("Edges:")
	// for key, value := range graph.edges {
	// 	fmt.Printf("%d -> %v\n", key, value)
	// }

	// sort and add
	middlePagesSum := 0
	fixedUpdateSum := 0
	for _, update := range parsedUpdates {
		subgraph := createSubgraph(graph, update)
		sorted := subgraph.topologicalSort()

		// Create position map for subgraph
		position := make(map[int]int)
		for i, page := range sorted {
			position[page] = i
		}

		// Validate update
		if isValidUpdate(update, position) {
			middlePage := update[len(update)/2]
			middlePagesSum += middlePage
		} else {
			fixed := fixUpdate(update, subgraph)

			if isValidUpdate(fixed, position) {
				middlePage := fixed[len(fixed)/2]
				fixedUpdateSum += middlePage
			}
		}
	}

	fmt.Println("Sum of middle pages of valid updates:", middlePagesSum)
	fmt.Println("Sum of middle pages of valid updates:", fixedUpdateSum)

}
